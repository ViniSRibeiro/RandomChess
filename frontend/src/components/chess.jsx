import { useState, useEffect, useRef } from "react";
import Chess from "chess.js";
import { Chessboard } from "react-chessboard";

export default function ChessTeste() {
  const [game, setGame] = useState(new Chess());
  const [turn, setTurn] = useState("");
  const ws = useRef(null);

  useEffect(() => {
    // Connect to WebSocket server
    ws.current = new WebSocket("ws://localhost:8080/partida/" + localStorage.getItem("idPartida"));

    ws.current.onmessage = (event) => {
      let msg = event.data;
      let from = msg.from
      let to = msg.to
      let promotion = msg.promotion

      const move = makeAMove({
        from: from,
        to: to,
        promotion: promotion,
      });

      setTurn(msg.turn)
    };

    ws.current.onclose = () => {
      console.log("WebSocket closed");
    };

    return () => {
      ws.current.close();
    };
  }, []);

  const sendMessage = (move) => {
    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      ws.current.send(move);
    }
  };


  // function makeAMove(move) {
  //   const gameCopy = { ...game };
  //   const result = gameCopy.move(move);
  //   setGame(gameCopy);
  //   return result; // null if the move was illegal, the move object if the move was legal
  // }
  function makeAMove(move) {
    const gameCopy = new Chess(game.fen()); // clone game safely

    const result = gameCopy.move(move, { sloppy: true });

    if (result === null) return null;

    // Manually set back the turn to the previous player
    const newFen = gameCopy.fen().split(' ');
    newFen[1] = turn; // force the turn back
    gameCopy.load(newFen.join(' '));
    setGame(gameCopy);

    if (game.game_over() || game.in_draw()) {
      alert("Jogo acabou")
      return
    }

    return result;
  }

  // function makeRandomMove() {
  //   const possibleMoves = game.moves();
  //   if (game.game_over() || game.in_draw() || possibleMoves.length === 0)
  //     return; // exit if the game is over
  //   const randomIndex = Math.floor(Math.random() * possibleMoves.length);
  //   makeAMove(possibleMoves[randomIndex]);
  // }

  function onDrop(sourceSquare, targetSquare) {
    if (!turn) return false
    const move = makeAMove({
      from: sourceSquare,
      to: targetSquare,
      promotion: "q", // always promote to a queen for example simplicity
    });

    // if illegal move
    if (move === null) return false;

    const data = {
      "from": sourceSquare,
      "to": targetSquare,
      "promotion": ["q", "r", "b", "n"].sort(() => Math.random() - 0.5)[0],
    }
    sendMessage(data)
    setTurn("")
    return true;
  }

  return <Chessboard
    position={game.fen()}
    boardOrientation={localStorage.getItem("color")}
    onPieceDrop={onDrop}
    autoPromoteToQueen={true} // always promote to a queen for example simplicity
  />;
}
