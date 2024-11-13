import { createSlice } from '@reduxjs/toolkit';

const leaderboardSlice = createSlice({
  name: 'leaderboard',
  initialState: {
    scores: [],
  },
  reducers: {
    updateLeaderboard: (state, action) => {
      const username = action.payload;
      console.log("Current leaderboard state:", state.scores);
      console.log("Updating leaderboard for:", username);
    
      const existingPlayer = state.scores.find((entry) => entry.username === username);
      if (existingPlayer) {
        existingPlayer.wins += 1;
      } else {
        state.scores.push({ username, wins: 1 });
      }
    
      console.log("Leaderboard state after update:", state.scores);
    },
    
  },
});

export const { updateLeaderboard } = leaderboardSlice.actions;
export default leaderboardSlice.reducer;
