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

const ChartComponent = ({ numbers }) => {
  const chartData = {
    labels: numbers.map((_, index) => `#${index + 1}`), // simple x-axis labels
    datasets: [
      {
        label: 'Random Values',
        data: numbers,
        borderColor: '#52b97d',
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
      {/* Chart container */}
      <div style={{ marginTop: '2rem' }}>
        <Line data={chartData} options={chartOptions} />
      </div>
    </div>
  );
};

export default ChartComponent;
