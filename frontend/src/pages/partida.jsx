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
    ws.current = new WebSocket("ws://localhost:8080/esperaJogo", [localStorage.getItem("token")]);

    ws.current.onmessage = (event) => {
      console.log(event)
      let msg = event.data;
      console.log(msg)

      // Caso tudo tenha dado certo
      alert("Sua partida foi encontrada")
      setTimeout(2000)
      navigate('/jogo')
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
      {
        searching ? (
          <Loading />
        ) : (
          <button onClick={sendMessage}>Buscar partida</button>
        )
      }
    </div>

  )
}

export default Partida
