import { useState, useEffect, useRef } from "react";
import Chess from "chess.js";
import { Chessboard } from "react-chessboard";

const url_back = process.env.REACT_APP_BACKEND_URL

export default function ChessOficial() {
  const [game, setGame] = useState(new Chess());
  const [turn, setTurn] = useState("");
  const ws = useRef(null);

  useEffect(() => {
    // Connect to WebSocket server
    const token = localStorage.getItem("token")
    const partida = localStorage.getItem("partida")
    ws.current = new WebSocket("ws://" + url_back + "/partida/" + partida, token);

    const color = localStorage.getItem("color")
    if (color === "white") {
      setTurn("w")
    }
    else if (color === "black") {
      setTurn("b")
    }
    else {
      console.log("CHESS: Um erro bizarro ocorreu. Recebi uma cor que não deveria")
    }

    ws.current.onmessage = (event) => {
      console.log("RECEBEU lance --------------------")
      let msg = event.data;
      msg = JSON.parse(msg)
      console.log(msg, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
      console.log("dlavhiavvavrj 123")

      let from = msg.from
      let to = msg.to
      let promotion = msg.promotion

      console.log("RECEBEU lance")
      console.log(msg)
      const move = makeAMove({
        from: from,
        to: to,
        promotion: promotion,
      }, msg.turn);
      console.log(" - Adversario fez o lance", from, to, promotion)
      setTurn(msg.nextTurn)
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
      console.log(JSON.stringify(move))
      ws.current.send(JSON.stringify(move));
      console.log("Lance enviado ao backend", JSON.stringify(move))
    }
    else {
      console.log("Erro enviando mensagem para o backend")
    }
  };

  function makeAMove(move, player) {
    console.log(move, player)
    const gameCopy = new Chess(game.fen()); // clone game safely

    const result = gameCopy.move(move, { sloppy: true });

    if (result === null) {
      console.log("Lance inválido")
      return null;
    }

    // Manually set back the turn to the previous player
    const newFen = gameCopy.fen().split(' ');
    newFen[1] = player; // force the turn back
    gameCopy.load(newFen.join(' '));
    setGame(gameCopy);

    if (game.game_over() || game.in_draw()) {
      alert("Jogo acabou")
      return null
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
    console.log("Entrou no ondrop")
    const color = localStorage.getItem("color")
    console.log(color, turn)
    if (!((turn === 'w' && color === "white") || (turn === 'b' && color === "black"))) {
      console.log("Não é sua vez de jogar")
      return
    }
    if (!ws.current || ws.current.readyState !== WebSocket.OPEN) {
      console.log("backend não conectado")
      return
    }

    let player = ""
    if (color === "white") {
      player = "w"
    } else if (color === "black") {
      player = "b"
    }

    const move = makeAMove({
      from: sourceSquare,
      to: targetSquare,
      promotion: "q", // always promote to a queen for example simplicity
    }, player);

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
