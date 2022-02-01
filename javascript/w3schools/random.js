// random number (0 - 10)
console.log(Math.floor(Math.random() * 11));

// random number (1 - 10)
console.log(Math.floor(Math.random() * 10) + 1);

// min (inklusif), max (eksklusif)
function getRandomInteger(min, max) {
  return Math.floor(Math.random() * (max - min)) + min;
}

// min, max (inklusif)
function getRandomInteger(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}
