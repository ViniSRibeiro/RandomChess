import React from "react"
import { useState } from "react"
import { useNavigate } from "react-router-dom"
import styles from '../styles/cadastro.module.css'

const url_back = process.env.REACT_APP_BACKEND_URL

const Cadastro = () => {
  const [nome, setNome] = useState("")
  const [senha, setSenha] = useState("")

  const navigate = useNavigate();

  const submit = (e) => {
    e.preventDefault()
    const data = {
      "nome": nome,
      "senha": senha
    }
    fetch(url_back + `/cadastro`, {
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
        navigate("/login");
      })
      .catch((error) => console.log(error))
  }

  return (
    <div className={styles.page}>
      <div className={styles.spacer}></div>
      <form className={styles.form} onSubmit={(e) => submit(e)}>
        <h1>Cadastrar</h1>
        <h2>Digite o seu nome</h2>
        <input type="text" value={nome} onChange={(e) => { setNome(e.target.value) }} required ></input>
        <h2>Digite sua senha</h2>
        <input type="password" value={senha} onChange={(e) => { setSenha(e.target.value) }} required ></input>
        <input className={styles.submit} type="submit" value="Cadastrar" />
      </form>
      <div className={styles.spacer}></div>
    </div>
  )
}

export default Cadastro
