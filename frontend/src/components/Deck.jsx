import React from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { drawCard, restartGame } from '../features/gameSlice';
import Card from './Card';

function Deck() {
  const dispatch = useDispatch();
  const { deck, currentCard, gameOver, message } = useSelector((state) => state.game);

  const handleDraw = () => {
    if (!gameOver) dispatch(drawCard());
  };

  const handleRestart = () => {
    dispatch(restartGame());
  };

  return (
    <div className="deck">
      {message && <p>{message}</p>}
      {currentCard && <Card card={currentCard} />}
      <button onClick={handleDraw} disabled={gameOver || deck.length === 0}>
        Draw Card
      </button>
      {gameOver && <button onClick={handleRestart}>Restart Game</button>}
    </div>
  );
}

export default Deck;
