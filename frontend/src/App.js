import React from 'react';
import UsernameForm from './components/UsernameForm';
import Deck from './components/Deck';
import Leaderboard from './components/Leaderboard';
import { useSelector } from 'react-redux';

function App() {
  const username = useSelector((state) => state.game.username);

  return (
    <div className="app">
      <h1>Online Single-Player Card Game</h1>
      {!username ? <UsernameForm /> : <Deck />}
      <Leaderboard />
    </div>
  );
}

export default App;
