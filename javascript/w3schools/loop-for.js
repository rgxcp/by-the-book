// For
for (i = 0; i < 5; i++) {
  console.log(`Hello, World ke-${i}!`);
}

// For/In
const cars = ["BMW", "Tesla", "Audi"];
for (car in cars) {
  console.log(cars[car]);
}

// For/Of
const fruits = ["Apple", "Mango", "Durian"];
for (fruit of fruits) {
  console.log(fruit);
}

const name = "Rommy";
for (letter of name) {
  console.log(letter);
}
