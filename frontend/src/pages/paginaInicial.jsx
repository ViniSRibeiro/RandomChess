import React from "react"
import styles from '../styles/paginaInicial.module.css'

const PaginaInicial = () => {
  return (
    <div className={styles.page}>
      <h1 className={styles.title}>RandomChess</h1>
      <img src="/page.png" alt="logo" />
    </div>
  )
}

export default PaginaInicial
