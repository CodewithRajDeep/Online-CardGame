import React from 'react';

function Card({ card }) {
  return <div className={`card ${card.type}`}>{card.icon} {card.type}</div>;
}

export default Card;
