import React from "react"
import styles from "../styles/header.module.css"
import { Link } from "react-router-dom"

const Header = () => {
  return (
    <div className={styles.container}>
      <Link className={styles.title} to="/">RandomChess</Link>
      <div className={styles.container}>
        <Link className={styles.link} to="/partida">Novo jogo</Link>
        <Link className={styles.link} to="/cadastro">Cadastro</Link>
        <Link className={styles.link} to="/login">Login</Link>
      </div>
    </div>
  )
}

export default Header

