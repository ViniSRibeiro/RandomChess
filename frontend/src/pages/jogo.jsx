import { useEffect, useRef, useState } from 'react';
import BitcoinMonitor from '../components/bitcoinMonitor';

function Jogo() {
  const [messages, setMessages] = useState([]);
  const ws = useRef(null);

  return (
    <div style={{ padding: "2rem" }}>
      <h1>WebSocket Demo</h1>
      <div>
        <BitcoinMonitor />
      </div>
    </div>
  );
}

export default Jogo;
