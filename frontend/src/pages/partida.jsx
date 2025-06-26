import React, { useEffect, useState, useRef } from "react"
import { useNavigate } from "react-router-dom"
import Loading from "../components/loading"
import styles from '../styles/partida.module.css'

const url_back = process.env.REACT_APP_BACKEND_URL

const Partida = () => {
  const [searching, setSearching] = useState(false)
  const ws = useRef(null);

  const navigate = useNavigate()

  useEffect(() => {
    // Connect to WebSocket server
    const token = localStorage.getItem("token")
    console.log(token)
    ws.current = new WebSocket("ws://" + url_back + "/esperaJogo", token);


    ws.current.onmessage = (event) => {
      let data = event.data;
      data = JSON.parse(data)
      console.log(data)
      let msg = data.mensagem;

      if (data.hasOwnProperty('encontrou')) {
        let encontrou = data.encontrou;
        if (encontrou === "S") {
          // Caso tudo tenha dado certo
          localStorage.setItem("partida", data.partida)
          localStorage.setItem("color", data.color)
          alert("Sua partida foi encontrada")
          navigate('/jogo')
        }
      }
      return
    };

    ws.current.onclose = () => {
      console.log("WebSocket closed");
    };

    return () => {
      ws.current.close();
    };

  }, []);

  const sendMessage = () => {
    setSearching(true)
    const token = localStorage.getItem("token")
    console.log("iniciou")
    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      const data = {
        method: 'GET',
        headers: {
          'Authorization': token,
          alg: 'HS256',
          typ: 'JWT',
        }
      }
      console.log(data)
      ws.current.send(data);
    }
  };

  return (
    <div className={styles.container}>
      <h1>Pesquisando por sua partida </h1>
      <Loading />
      <p>Aguardando oponente</p>
    </div>

  )
}

export default Partida
