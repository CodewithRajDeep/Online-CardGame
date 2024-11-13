import React from 'react';
import { useDispatch } from 'react-redux';
import { updateLeaderboard } from '../features/leaderboardSlice';

function TestButton() {
  const dispatch = useDispatch();

  const handleClick = () => {
    dispatch(updateLeaderboard("TestUser")); // Use a test username
  };

  return <button onClick={handleClick}>Test Leaderboard Update</button>;
}

export default TestButton;
