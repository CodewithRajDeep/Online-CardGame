import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';


export const updateLeaderboardBackend = createAsyncThunk(
  'leaderboard/updateLeaderboardBackend',
  async (username, { getState }) => {
    
    const response = await fetch('https://your-backend-url/leaderboard', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username }),
    });
    const data = await response.json();
    return data;
  }
);

const leaderboardSlice = createSlice({
  name: 'leaderboard',
  initialState: {
    scores: [],
    status: 'idle', 
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
  extraReducers: (builder) => {
    builder
      .addCase(updateLeaderboardBackend.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(updateLeaderboardBackend.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.scores = action.payload.scores; 
      })
      .addCase(updateLeaderboardBackend.rejected, (state) => {
        state.status = 'failed';
      });
  },
});

export const { updateLeaderboard } = leaderboardSlice.actions;
export default leaderboardSlice.reducer;
