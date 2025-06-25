import logo from './logo.svg';
import './App.css';
import {chessboard} from 'react-chessboard';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> aaaaaaaa and save to reload.
        </p>
        <div>
        <Chessboard id="BasicBoard" />
        </div>
      </header>
    </div>
  );
}

export default App;
