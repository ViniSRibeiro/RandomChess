import React from 'react';
import { Line } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  LineElement,
  PointElement,
  LinearScale,
  CategoryScale,
  Tooltip,
  Legend,
} from 'chart.js';

import styles from '../styles/cadastro.module.css'

ChartJS.register(LineElement, PointElement, LinearScale, CategoryScale, Tooltip, Legend);

const numbers = [
  6.127, 0.913, -6.861, 0.384, -2.263, 7.995, -0.842,
  8.000, -4.337, -7.652, 2.789, 1.073, -3.982, 2.473, -1.157, 0.000, -3.509,
  5.448, 4.782, -5.916, 1.936, -6.004, 3.621, -8.000, -0.175, 6.845, -1.490, 2, 3.311, -7.218,
];

const ChartComponent = () => {
  const chartData = {
    labels: numbers.map((_, index) => `#${index + 1}`), // simple x-axis labels
    datasets: [
      {
        label: 'Random Values',
        data: numbers,
        borderColor: 'rgba(75,192,192,1)',
        backgroundColor: 'rgba(75,192,192,0.2)',
        tension: 0.3,
        pointRadius: 2,
      },
    ],
  };

  const chartOptions = {
    responsive: true,
    animation: false,
    plugins: {
      legend: {
        display: false
      },
    },
    scales: {
      y: {
        beginAtZero: false,
      },
    },
  };

  return (
    <div>
      {/* Fading list */}
      <div className="fade-container">
        {numbers.map((num, idx) => (
          <p key={idx} className={num >= 0 ? styles.pos : styles.neg}>
            {num >= 0 ? `+${num.toFixed(3)}` : num.toFixed(3)}
          </p>
        ))}
      </div>

      {/* Chart container */}
      <div style={{ marginTop: '2rem' }}>
        <Line data={chartData} options={chartOptions} />
      </div>
    </div>
  );
};

export default ChartComponent;
