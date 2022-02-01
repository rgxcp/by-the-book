// For Each
const cars = ["Tesla", "BMW", "Mazda"];
function logEachCar(value, index, array) {
  // urutan param harus urut seperti ini
  console.log(`Mobil ke-${index + 1}: ${value}`);
}
cars.forEach(logEachCar);

// Map
function createNewArray(value, index, array) {
  return `Mobil ke-${index + 1}: ${value}`;
}
console.log(cars.map(createNewArray));

// Filter
const usersAge = [10, 20, 30, 40, 50];
function over18(value, index, array) {
  return value > 18;
}
console.log(usersAge.filter(over18));

// Reduce
function reduceAge(total, value, index, array) {
  // urutan param harus urut seperti ini
  return total + value;
}
console.log(usersAge.reduce(reduceAge)); // param kedua bisa di isi dengan nilai awal

// Every
console.log(usersAge.every(over18));

// Some
console.log(usersAge.some(over18));

// Index Of
console.log(cars.indexOf("Mazda"));

// Last Index Of
console.log(cars.lastIndexOf("BMW"));

// Find
console.log(usersAge.find(over18));

// Find Index
console.log(usersAge.findIndex(over18));
