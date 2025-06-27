import React from "react";
import ChartComponent from "./chart";
import styles from '../styles/bitcoinMonitor.module.css'
import { useState, useEffect, useRef } from "react";

const url_back = process.env.REACT_APP_BACKEND_URL

const MAX = 20

const BitcoinMonitor = () => {
  const [bitcoin, setBitcoin] = useState(Array(MAX).fill(0));
  const [bitcoinI, setBitcoinI] = useState(Array(MAX).fill(0));
  const ws = useRef(null);

  const updateMessages = (num) => {
    setBitcoin((prev) => {
      return [num, ...prev.slice(0, -1)];
    });
    setBitcoinI((prev) => {
      return [...prev.slice(1), num];
    });
  }
  useEffect(() => {
    // Connect to WebSocket server
    ws.current = new WebSocket("ws://" + url_back + "/random");

    ws.current.onmessage = (event) => {
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
