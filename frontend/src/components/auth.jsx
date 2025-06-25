import React, { createContext, useState, useEffect } from 'react'
import { useContext } from 'react'
import { useNavigate } from 'react-router-dom'


const AuthContext = createContext({})

export const AuthProvider = ({ children }) => {
  const navigate = useNavigate()
  const [isLogged, setIsLogged] = useState(null)

  const logout = () => {
    localStorage.clear()
    setIsLogged(false);
    navigate('/')
  };

  useEffect(() => {
    if (localStorage.getItem("token") !== null) {
      setIsLogged(true)
    }
  }, [])


  const cadastro = (data) => {
    console.log(data)
    fetch("http://localhost:8080/cadastro", {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
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
        // Caso tudo tenha dado certo
        alert("Cadastro realizado com sucesso!!")
        setTimeout(2000)
        navigate('/login')
      })
      .catch((error) => {
        alert("um erro inesperado ocorreu ao cadastrar o professor");
        console.log("cadastro.jsx >>> ", error)
        throw new error
      })
  }

  const login = (data) => {
    console.log(data)
    fetch("http://localhost:8080/login", {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
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
              console.log("login.jsx >>> ", error)
              return
            })
          return
        }
        const obj = response.json()
          .then((data) => {
            localStorage.setItem("token", data.token)
          })
        // Caso tudo tenha dado certo
        alert("Login realizado com sucesso!!")
        setTimeout(2000)
        setIsLogged(true)
        navigate('/partida')
      })
      .catch((error) => {
        alert("um erro inesperado ocorreu ao cadastrar o professor");
        console.log("login.jsx >>> ", error)
        throw new error
      })
  }
  return (
    <AuthContext.Provider
      value={{
        login,
        cadastro,
        logout,
        isLogged
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}

export function useAuth() {
  const context = useContext(AuthContext)
  return context
}
