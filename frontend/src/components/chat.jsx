import { useEffect, useRef, useState } from 'react';
import styles from '../styles/chat.module.css'

const url_back = process.env.REACT_APP_BACKEND_URL

function Chat() {
  const [messages, setMessages] = useState([]);
  const [text, setText] = useState("")
  const ws = useRef(null);

  const updateMessages = (num) => {
    setMessages((prev) => {
      return [...prev, num];
    });
  }
  useEffect(() => {
    // Connect to WebSocket server
    const token = localStorage.getItem("token")
    ws.current = new WebSocket("ws://" + url_back + "/chat", token);

    ws.current.onmessage = (event) => {
      console.log(event)
      let msg = event.data;
      updateMessages(msg)
    };

    ws.current.onclose = () => {
      console.log("WebSocket closed");
    };

    return () => {
      ws.current.close();
    };
  }, []);

  const sendMessage = () => {
    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      const data = {
        "msg": text
      }
      ws.current.send(data);
      setText("")
    }
  };

  return (

    <div className={styles.container}>
      <h1>Chat</h1>

      <ul className={styles.messages}>
        {messages.map((m, i) => (
          <li key={i}>{JSON.stringify(m)}</li>
        ))}
      </ul>

      <input
        type="text"
        value={text}
        placeholder="Digite aqui sua mensagem"
        onChange={(e) => setText(e.target.value)}
      />
      <button onClick={sendMessage}>Send</button>
    </div>
  );
}

export default Chat;
