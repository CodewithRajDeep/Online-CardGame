import { configureStore } from '@reduxjs/toolkit';
import gameReducer from './features/gameSlice';
import leaderboardReducer from './features/leaderboardSlice';

const store = configureStore({
  reducer: {
    game: gameReducer,
    leaderboard: leaderboardReducer,
  },
});

export default store;
