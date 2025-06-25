import React from "react"
import { useState } from "react"
import { useNavigate } from "react-router-dom"
import styles from '../styles/login.module.css'

const url_back = process.env.REACT_APP_BACKEND_URL

const Login = () => {
  const [nome, setNome] = useState()
  const [senha, setSenha] = useState()

  const router = useNavigate()

  const submit = (e) => {
    e.preventDefault()
    const data = {
      "nome": nome,
      "senha": senha
    }

    fetch(url_back + `/login`, {
      method: 'POST',
      body: data,
    })
      .then(async (response) => {
        if (!response.ok) {
          const errorData = await response.json().catch(() => ({}));
          throw new Error(errorData.mensagem || "Erro ao cadastrar usuário.");
        }
        return response.json();
      })
      .then(() => {
        console.log("login feito com sucesso")
      })
      .catch((error) => console.log(error))
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
