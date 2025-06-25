import React from "react";
import ChartComponent from "./chart";
import styles from '../styles/bitcoinMonitor.module.css'
import { useState, useEffect, useRef } from "react";

const numbers = [
  6.127, 0.913, -6.861, 0.384, -2.263, 7.995, -0.842,
  8.000, -4.337, -7.652, 2.789, 1.073, -3.982, 2.473, -1.157, 0.000, -3.509,
  5.448, 4.782, -5.916, 1.936, -6.004, 3.621, -8.000, -0.175, 6.845, -1.490, 2, 3.311, -7.218
];


const MAX = 20

const BitcoinMonitor = () => {
  const [bitcoin, setBitcoin] = useState([]);
  const [bitcoinI, setBitcoinI] = useState([]);
  const ws = useRef(null);

  const updateMessages = (num) => {
    setBitcoin((prev) => {
      if (prev.length >= MAX) {
        return [num, ...prev.slice(0, -1)];
      }
      return [num, ...prev];
    });
    setBitcoinI((prev) => {
      if (prev.length >= MAX) {
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
      updateMessages(msg.valor)
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
    <div>
      <div className={styles.chart_box}>
        <ChartComponent numbers={bitcoinI} />
      </div>
      <div className={styles.values}>
        {bitcoin.map((num, idx) => (
          <p
            key={idx}
            className={num >= 0 ? styles.pos : styles.neg}
          >
            {num >= 0 ? `+${num.toFixed(3)}` : num.toFixed(3)}
          </p>
        ))}
      </div>
    </div>
  )
}

export default BitcoinMonitor
