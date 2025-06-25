import React, { useState, useEffect } from "react";
import styles from "../styles/header.module.css";
import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "./auth";

const Header = () => {
  const { logout, isLogged } = useAuth()

  return (
    <div className={styles.container}>
      <Link className={styles.title} to="/">RandomChess</Link>
      <div className={styles.pages}>
        {!isLogged ? (
          <>
            <Link className={styles.link} to="/cadastro">Cadastro</Link>
            <Link className={styles.link} to="/login">Login</Link>
          </>
        ) : (
          <>
            <Link className={styles.link} to="/partida">Novo jogo</Link>
            <Link className={styles.link} to="/perfil">Perfil</Link>
            <button className={styles.logout} onClick={logout}>Logout</button>
          </>
        )}
      </div>
    </div>
  );
};

export default Header;
