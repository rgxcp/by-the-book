const name = "Rommy Gustiawan Gustiawan";
const message = "   Hello, World!   ";
const str = "a,b,c,d,e";

// String Length
console.log(name.length);

// Finding a String in a String
console.log(name.indexOf("Gustiawan")); // param kedua bisa di isi starting point dengan tipe number
console.log(name.lastIndexOf("Gustiawan")); // param kedua bisa di isi starting point dengan tipe number

// Searching for a String in a String
console.log(name.search("Gustiawan")); // sama seperti `indexOf()`, tapi tidak ada param ke dua

// Extracting String Parts
console.log(name.slice(6, 11)); // jika hanya satu param, maka kata selanjutnya akan ikut
console.log(name.substring(6, 11)); // sama seperti `slice()`, tapi tidak bisa negatif
console.log(name.substr(6, 5)); // sama seperti `slice()`, tapi param kedua panjang yang ingin di slice

// Replacing String Content
console.log(name.replace("Gustiawan", "Chandra")); // sama seperti search(), hanya kata pertama yang di proses

// Converting to Upper and Lower Case
console.log(name.toUpperCase());
console.log(name.toLowerCase());

// Concat
console.log(name.concat(" ", "Chandra")); // alternatif penggunaan `+`

// Trim
console.log(message.trim()); // menghapus white space di awal dan akhir

// Extracting String Characters
console.log(name.charAt(6));
console.log(name[6]);
console.log(name.charCodeAt(6)); // tipe code dari char tersebut
console.log(name[6]); // hanya digunakan jika tipe-nya array

// Converting a String to an Array
const arrayOfStr = str.split(","); // bisa memisahkan koma, spasi, dan garis lurus
console.log(arrayOfStr[3]);
