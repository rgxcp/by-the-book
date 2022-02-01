const os = ["Android", "iOS", "Windows Phone"];

console.log(Array.isArray(os)); // mengecek apakah sebuah variable array
console.log(os instanceof Array); // `instanceof` untuk mengecek instance sebuah variable

// Converting Arrays to Strings
console.log(os.toString()); // secara default akan di ubah menjadi string
console.log(os.join(" | ")); // sama seperti `toString()`, tapi bisa ditambahkan separasi

// Popping
os.pop(); // menghapus elemen terakhir dalam array, me-return value index yang di hapus
console.log(os);

// Pushing
os.push("Symbian"); // menambahkan elemen di posisi terakhir array, me-return panjang array
console.log(os);

// Shifting Elements
os.shift(); // sama seperti `pop()`, tapi yang di hapus elemen pertama, me-return value index yang di hapus
console.log(os);
os.unshift("KaiOS"); // sama seperti `push()`, tapi yang di tambah elemen di posisi pertama, me-return panjang array
console.log(os);

// Changing Elements
os[0] = "Android";

// Deleting Elements
delete os[0]; // lebih disarankan menggunakan `pop()` atau `shift()`

// Splicing an Array
// untuk menambahkan ke dalam array juga
// param pertama posisi
// param kedua banyak elemen yang akan di hapus
// me-return elemen yang di hapus
// jika param ketiga tidak ada, maka fungsinya untuk menghapus elemen
os.splice(2, 0, "LineageOS", "RemixOS");

// Merging (Concatenating) Arrays
const array1 = ["a", "b"];
const array2 = ["c", "d"];
const array3 = ["e", "f"];
console.log(array1.concat(array2, array3)); // menggabungkan dua atau lebih array
console.log(array1.concat("c")); // juga bisa digunakan untuk menambahkan elemen

// Slicing an Array
const array4 = array1.concat(array2, array3);
// memotong array dan mengambil elemen selebihnya
// param kedua bisa di isi elemen yang di ambil secara spesifik, tapi elemen yang terakhir tidak ikut
console.log(array4.slice(3));
