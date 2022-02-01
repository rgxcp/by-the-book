const a = 50;
const b = 5.1234;

console.log(a.toString());
console.log(a.valueOf()); // mengubah object number menjadi primitive
console.log(b.toFixed(0)); // membulatkan angka, param bisa 0
console.log(b.toPrecision(2)); // membulatkan angka, param tidak bisa 0
console.log(Number.isInteger(b));
console.log(Number.isSafeInteger(b));
console.log(isFinite(5 / 0));
console.log(isNaN("Hello"));

// Converting Variables to Numbers
console.log(Number("5.5")); // jika tidak bisa diubah, me-return `NaN`
console.log(parseInt("5.5")); // angka bulat, hanya yang pertama, jika tidak bisa diubah, me-return `NaN`
console.log(parseFloat("5.5")); // sama seperti `parseInt()`, tapi untuk angka desimal
