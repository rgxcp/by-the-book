const a = "10";
const b = "40";

console.log(`a + b = ${a + b}`); // `+` tidak valid, tapi operator lainnya bisa
console.log(`a - b = ${a - b}`);
console.log(`a * b = ${a * b}`);
console.log(`a / b = ${a / b}`);
console.log(`a / "Rommy" = ${a / "Rommy"}`); // tidak valid, menjadi `NaN`
console.log(`a / "Rommy" isNaN = ${isNaN(a / "Rommy")}`); // mengecek apakah `NaN`
console.log(`5 / 0 = ${5 / 0}`); // menjadi `Infinity`
