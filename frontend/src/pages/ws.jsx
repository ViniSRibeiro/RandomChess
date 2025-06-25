import { useEffect, useRef, useState } from 'react';

function Ws() {
  const [messages, setMessages] = useState([]);
  const ws = useRef(null);

  const updateMessages = (num) => {
    setMessages((prev) => {
      if (prev.length >= 30) {
        return [...prev.slice(1), num];
      }
      return [...prev, num];
    });
  }
  useEffect(() => {
    // Connect to WebSocket server
    ws.current = new WebSocket("ws://localhost:8080/random");

    ws.current.onmessage = (event) => {
      console.log(event)
      let msg = event.data;
      msg = JSON.parse(msg)
      updateMessages(msg.token)
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
      ws.current.send("👋 Button clicked!");
    }
  };

  return (
    <div style={{ padding: "2rem" }}>
      <h1>WebSocket Demo</h1>
      <button onClick={sendMessage}>Click me to broadcast!</button>
      <div>
        <h2>Messages:</h2>
        <ul>
          {messages.map((m, i) => (
            <li key={i}>{m}</li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default Ws;
