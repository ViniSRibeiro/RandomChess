import { useEffect, useRef, useState } from 'react';
import BitcoinMonitor from '../components/bitcoinMonitor';
import styles from '../styles/jogo.module.css'
import ChessOficial from '../components/chess';
import Chat from '../components/chat';

function Jogo() {
  const [messages, setMessages] = useState([]);
  const ws = useRef(null);

  return (
    <div className={styles.container}>
      <div className={styles.board_box}>
        <ChessOficial />
      </div>
      <div className={styles.graph_chat_container}>
        <div className={styles.bitcoin}>
          <BitcoinMonitor />
        </div>
        <div className={styles.chat}>
          <Chat />
        </div>    

      </div>
    </div>
  );
}

export default Jogo;
