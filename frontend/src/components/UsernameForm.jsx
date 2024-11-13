import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import { setUsername } from '../features/gameSlice';

function UsernameForm() {
  const [username, setUsernameInput] = useState('');
  const dispatch = useDispatch();

  const handleSubmit = (e) => {
    e.preventDefault();
    if (username) {
      dispatch(setUsername(username));
    }
  };

  return (
    <form onSubmit={handleSubmit} className="username-form">
      <input
        type="text"
        value={username}
        onChange={(e) => setUsernameInput(e.target.value)}
        placeholder="Enter your username"
      />
      <button type="submit">Start Game</button>
    </form>
  );
}

export default UsernameForm;
