import React, { useEffect } from "react"
import { useState } from "react"
import { useNavigate } from "react-router-dom"
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
      fetch(url_back + `/iniciarPartida`, {
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
          navigate('/jogo')
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
          <ChartComponent />
        </div>
      </div>
      <div className={styles.bitcoin}>
        <p className={styles.neg}>-4.776</p>
        <p className={styles.pos}>+3.311</p>
        <p className={styles.neg}>-7.218</p>
        <p className={styles.pos}>+6.127</p>
        <p className={styles.pos}>+0.913</p>
        <p className={styles.neg}>-6.861</p>
        <p className={styles.pos}>+0.384</p>
        <p className={styles.neg}>-2.263</p>
        <p className={styles.pos}>+7.995</p>
        <p className={styles.neg}>-0.842</p>
        <p className={styles.pos}>+8.000</p>
        <p className={styles.neg}>-4.337</p>
        <p className={styles.neg}>-7.652</p>
        <p className={styles.pos}>+2.789</p>
        <p className={styles.pos}>+1.073</p>
        <p className={styles.neg}>-3.982</p>
        <p className={styles.pos}>+2.473</p>
        <p className={styles.neg}>-1.157</p>
        <p className={styles.pos}>+0.000</p>
        <p className={styles.neg}>-3.509</p>
        <p className={styles.pos}>+5.448</p>
        <p className={styles.pos}>+4.782</p>
        <p className={styles.neg}>-5.916</p>
        <p className={styles.pos}>+1.936</p>
        <p className={styles.neg}>-6.004</p>
        <p className={styles.pos}>+3.621</p>
        <p className={styles.neg}>-8.000</p>
        <p className={styles.neg}>-0.175</p>
        <p className={styles.pos}>+6.845</p>
        <p className={styles.neg}>-1.490</p>
      </div>

    </div>

  )
}

export default Partida
