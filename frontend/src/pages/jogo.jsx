import { useEffect, useRef, useState } from 'react';
import BitcoinMonitor from '../components/bitcoinMonitor';
import styles from '../styles/jogo.module.css'
import ChessTeste from '../components/chess';

function Jogo() {
  const [messages, setMessages] = useState([]);
  const ws = useRef(null);

  return (
    <div className={styles.container}>
      <div className={styles.board_box}>
        <ChessTeste />
      </div>
      <div className={styles.bitcoin}>
        <BitcoinMonitor />
      </div>
    </div>
  );
}

export default Jogo;
