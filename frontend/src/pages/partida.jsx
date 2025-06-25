import React, { useEffect } from "react"
import { useState } from "react"
import { useNavigate } from "react-router-dom"
import BitcoinMonitor from "../components/bitcoinMonitor"
import ChartComponent from "../components/chart"
import ChessTeste from "../components/chess"
import Loading from "../components/loading"
import styles from '../styles/partida.module.css'

const url_back = process.env.REACT_APP_BACKEND_URL

const Partida = () => {
  const [started, setStarted] = useState(false)

  const navigate = useNavigate()

  useEffect(() => {
    const iniciarPartida = () => {
      const token = localStorage.getItem("token")
      fetch("http://localhost:8080/ok", {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`,
          alg: 'HS256',
          typ: 'JWT',
        }
      })
        .then((response) => {
          if (!response.ok) {
            const obj = response.json()
              .then((data) => {
                if (data.hasOwnProperty('mensagem')) {
                  alert(data.mensagem)
                  return
                }
                else {
                  throw response
                }
              })
              .catch((error) => {
                console.log("cadastro.jsx >>> ", error)
                return
              })
            return
          }
          // Caso tudo tenha dado certo
          alert("Sua partida foi encontrada")
          setTimeout(2000)
          // navigate('/jogo')
        })
        .catch((error) => {
          alert("um erro inesperado ocorreu ao iniciar a partida");
          console.log("cadastro.jsx >>> ", error)
        })

    }
    iniciarPartida()
  }, [])

  return (
    <div className={styles.container}>
      <h1> partida </h1>
      <Loading />
      <div className={styles.partida}>
        <div className={styles.board_box}>
          <ChessTeste />
        </div>
        <div className={styles.chart_box}>
          {/* <ChartComponent /> */}
        </div>
        <BitcoinMonitor />
      </div>
    </div>

  )
}

export default Partida
