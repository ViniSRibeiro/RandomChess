import React from "react"
// import ChessTeste from "../components/chess"

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
      {/* <ChessTeste /> */}
    </div>
  )
}

export default Partida
