import React from "react"
import { useState } from "react"
import { useNavigate } from "react-router-dom"
import { useAuth } from "../components/auth"
import styles from '../styles/login.module.css'

const Login = () => {
  const { login } = useAuth()
  const [nome, setNome] = useState("")
  const [senha, setSenha] = useState("")

  const submit = (e) => {
    e.preventDefault()
    const data = {
      "nome": nome,
      "senha": senha
    }
    login(data)
  }

  return (
    <div className={styles.page}>
      <div className={styles.spacer}></div>
      <form className={styles.form} onSubmit={(e) => submit(e)}>
        <h1>Login</h1>
        <h2>Digite o seu nome</h2>
        <input type="text" value={nome} onChange={(e) => { setNome(e.target.value) }} required ></input>
        <h2>Digite sua senha</h2>
        <input type="password" value={senha} onChange={(e) => { setSenha(e.target.value) }} required ></input>
        <input className={styles.submit} type="submit" value="Login" />
      </form>
      <div className={styles.spacer}></div>
    </div>
  )
}

export default Login 
