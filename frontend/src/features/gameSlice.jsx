/* eslint-disable */
import { createSlice } from '@reduxjs/toolkit';


const initialState = {
  username: '',
  deck: [],
  currentCard: null,
  defuseAvailable: false,
  gameOver: false,
  message: '',
  points: 0, 
};

const cardTypes = [
  { type: 'Cat', icon: '😼', points: 10 },
  { type: 'Defuse', icon: '🙅‍♂️', points: 20 },
  { type: 'Shuffle', icon: '🔀', points: 5 },
  { type: 'Exploding Kitten', icon: '💣', points: -10 },
];

const gameSlice = createSlice({
  name: 'game',
  initialState,
  reducers: {
    setUsername: (state, action) => {
      state.username = action.payload;
      state.deck = initializeDeck();
      state.gameOver = false;
      state.message = '';
      state.points = 0; // Reset points at the start of a new game
    },
    drawCard: (state) => {
      if (state.deck.length === 0) return;

      const drawnCard = state.deck.pop();
      state.currentCard = drawnCard;
      state.points += drawnCard.points; // Award points based on card type

      switch (drawnCard.type) {
        case 'Cat':
          state.message = 'You drew a Cat card!';
          break;
        case 'Defuse':
          state.defuseAvailable = true;
          state.message = 'You drew a Defuse card!';
          break;
        case 'Shuffle':
          state.deck = initializeDeck();
          state.message = 'Deck has been shuffled!';
          break;
        case 'Exploding Kitten':
          if (state.defuseAvailable) {
            state.defuseAvailable = false;
            state.message = 'Exploding Kitten defused!';
          } else {
            state.gameOver = true;
            state.message = 'Boom! You lost the game.';
          }
          break;
        default:
          state.message = 'An unknown card type was drawn.';
          break;
      }

      if (!state.gameOver && state.deck.length === 0) {
        state.gameOver = true;
        state.message = 'Congratulations! You won!';
      }
    },
    restartGame: (state) => {
      state.username = '';
      state.deck = initializeDeck();
      state.currentCard = null;
      state.defuseAvailable = false;
      state.gameOver = false;
      state.message = '';
      state.points = 0;
    },
    resetGame: (state) => {
      return initialState;
    },
  },
});

const initializeDeck = () => {
  return Array.from({ length: 5 }, () => cardTypes[Math.floor(Math.random() * cardTypes.length)]);
};

export const { setUsername, drawCard, restartGame, resetGame } = gameSlice.actions;
export default gameSlice.reducer;
