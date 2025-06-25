import React from "react"
import { Chessboard } from "react-chessboard"

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
      <Chessboard />
    </div>
  )
}

export default Partida
