import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import { setUsername } from '../features/gameSlice';
import { useNavigate } from 'react-router-dom';
import UsernameForm from './UsernameForm';  // Ensure this import is correct
import Leaderboard from './Leaderboard';  // Ensure this import is correct

function HomeScreen() {
    // eslint-disable-next-line
  const [name, setName] = useState('');
  const dispatch = useDispatch();
  const navigate = useNavigate();
// eslint-disable-next-line
  const handleNameSubmit = (e) => {
    e.preventDefault();
    if (name) {
      dispatch(setUsername(name));  
      navigate('/game');
    }
  };

  return (
    <div>
      <h2>Welcome to the Online Single-Player Card Game!</h2>
      <Leaderboard />

     
      {!name ? (
        <UsernameForm /> 
      ) : (
        <div>
          <p>Welcome, {name}!</p>
          <button onClick={() => navigate('/game')}>Start Game</button>
        </div>
      )}
    </div>
  );
}

export default HomeScreen;
