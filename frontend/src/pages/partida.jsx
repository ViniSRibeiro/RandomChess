import React, { useEffect, useState, useRef } from "react"
import { useNavigate } from "react-router-dom"
import Loading from "../components/loading"
import styles from '../styles/partida.module.css'

const url_back = process.env.REACT_APP_BACKEND_URL

const Partida = () => {
  const [searching, setSearching] = useState(false)
  const ws = useRef(null);

  const navigate = useNavigate()

  const click = () => {
    localStorage.setItem("color", Math.random() < 0.5 ? "white" : "black")
    // Connect to WebSocket server
    alert("Sua partida foi encontrada")
    navigate('/jogo')
  }


  return (
    <div className={styles.container}>
      <h1 onClick={click}>Pesquisando por sua partida </h1>
      <Loading />
      <p>Aguardando oponente</p>
    </div>

  )
}

export default Partida
