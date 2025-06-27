import React from 'react';
import styles from '../styles/loading.module.css'; // CSS for animation

const Loading = ({ size = 40, color = '#52b97d' }) => {
  return (
    <div
      className={styles.spinner}
      style={{
        width: size,
        height: size,
        borderColor: `${color} transparent ${color} transparent`,
      }}
    />
  );
};

export default Loading;
