import React from "react"
import ChartComponent from "../components/chart"
import ChessTeste from "../components/chess"
import styles from '../styles/partida.module.css'

const url_back = process.env.REACT_APP_BACKEND_URL

const Partida = () => {

  // fetch(url_back + `/cadastro`, {
  //   method: 'POST',
  //   headers: {
  //     'Authorization': `Bearer ${token}`,
  //     alg: 'HS256',
  //     typ: 'JWT',
  //   },
  //   body: formData,
  //   "nome": nome,
  //   "senha": senha
  // })
  return (
    <div>
      <h1>Partida</h1>
      <ChessTeste />
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
      <ChartComponent />
    </div>
  )
}

export default Partida
