import React, { useState } from "react"

const url_back = process.env.REACT_APP_BACKEND_URL

const Perfil = () => {

  const [nome, setNome] = useState("")
  const [totPartidasJogadas, setTotPartidasJogadas] = useState("")
  const [totVitorias, setTotVitorias] = useState("")
  const [totEmpates, setTotEmpates] = useState("")
  const [totDerrotas, setTotDerrotas] = useState("")

  const getData = () => {
    fetch("http://" + url_back + "/perfil", {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authentication': localStorage.getItem('token')
      },
    })
      .then((response) => {
        if (!response.ok) {
          // Primeiramente é tentado recuperar os campos da resposta que indicam que deu errado. Se não conseguir, trata como um erro genérico
          const obj = response.json()
            .then((data) => {
              if (data.hasOwnProperty('mensagem')) {
                alert(data.mensagem)
                return
              }
              else {
                throw response
              }
            })
            .catch((error) => {
              console.log("cadastro.jsx >>> ", error)
              return
            })
          return
        }
        setTimeout(2000)
        alert("RECEBEU DATA ")
      })
      .catch((error) => {
        alert("um erro inesperado ocorreu ao cadastrar o professor");
        console.log("cadastro.jsx >>> ", error)
        throw new error
      })
  }

  return (
    <div>
      <p>ALSKJLFF</p>
    </div>
  )
}

export default Perfil
