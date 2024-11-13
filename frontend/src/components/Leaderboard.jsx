import React, { useEffect } from 'react';
import { useSelector } from 'react-redux';

function Leaderboard() {
  const leaderboard = useSelector((state) => state.leaderboard.scores);
  console.log("Rendering leaderboard with data:", leaderboard);
  useEffect(() => {
    console.log('Leaderboard updated:', leaderboard);
  }, [leaderboard]); 

  return (
    <div className="leaderboard">
      <h2>Leaderboard</h2>
      <ul>
        {leaderboard.map((entry, index) => (
          <li key={index}>
            {entry.username}: {entry.wins} wins
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Leaderboard;
