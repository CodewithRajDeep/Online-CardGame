import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { drawCard, restartGame, resetGame } from '../features/gameSlice';
import { updateLeaderboard } from '../features/leaderboardSlice';
import { useNavigate } from 'react-router-dom';

function GameScreen() {
  const dispatch = useDispatch();
  const navigate = useNavigate();  
  const { gameOver, username, message, points } = useSelector((state) => state.game); // Get points from state

  useEffect(() => {
    if (gameOver && message.includes('won')) {
      console.log("Updating leaderboard for:", username);
      dispatch(updateLeaderboard(username));
    }
  }, [gameOver, message, username, dispatch]);

  const handleDrawCard = () => {
    dispatch(drawCard());
  };

  const handleRestartGame = () => {
    dispatch(restartGame());
  };

  const handleBackToMain = () => {
    dispatch(resetGame());  
    navigate('/'); 
  };

  return (
    <div>
      <h2>Welcome, {username}!</h2>
      <button onClick={handleDrawCard} disabled={gameOver}>Draw Card</button>
      <button onClick={handleRestartGame}>Restart Game</button>
      <button onClick={handleBackToMain}>Back to Main</button>
      {message && <p>{message}</p>}
      <p>Points: {points}</p> 
      {gameOver && <p>Game Over!</p>}
    </div>
  );
}

export default GameScreen;
