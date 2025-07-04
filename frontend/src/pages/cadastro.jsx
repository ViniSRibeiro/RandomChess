import React from "react"
import { useState } from "react"
import { useNavigate } from "react-router-dom"
import { useAuth } from "../components/auth"
import styles from '../styles/cadastro.module.css'

const Cadastro = () => {
  const { cadastro } = useAuth()
  const [nome, setNome] = useState("")
  const [senha, setSenha] = useState("")
  const [senhaRepetida, setSenhaRepetida] = useState("")

  const submit = (e) => {
    e.preventDefault()

    if (senha !== senhaRepetida) {
      alert("ATENÇÃO: As senhas devem ser iguais")
      return
    }

    const data = {
      "nome": nome,
      "senha": senha
    }
    cadastro(data)
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
        <h2>Repita sua senha</h2>
        <input type="password" value={senhaRepetida} onChange={(e) => { setSenhaRepetida(e.target.value) }} required ></input>
        <input className={styles.submit} type="submit" value="Cadastrar" />
      </form>
      <div className={styles.spacer}></div>
    </div>
  )
}

export default Cadastro
