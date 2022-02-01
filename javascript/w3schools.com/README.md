# JavaScript

## Resource

[w3schools.com](https://w3schools.com/js)

## Introduction

1. JavaScript bisa digunakan untuk mengubah elemen konten HTML
   `document.getElementById("demo").innerHTML = "Hello JavaScript"`
   Mengubah value atribut elemen HTML
   `document.getElementById("myImage").src='pic_bulbon.gif`
   Mengubah CSS
   `document.getElementById("demo").style.fontSize = "35px"`
   Menampilkan atau menyembunyikan elemen HTML
   `document.getElementById("demo").style.display = "block/none"`

## Where To

1. JavaScript diapit dengan menggunakan tag `<script>`
2. JavaScript bisa diletakkan dalam tag `<head>` atau `<body>` dan file eksternal
3. Pemanggilan JavaScript dengan file eksternal dalam folder website aktif
   `<script src="/js/myScript1.js"></script>`
   Dalam folder halaman aktif
   `<script src="myScript1.js"></script>`

## Output

1. JavaScript bisa menampilkan data dengan cara berbeda-beda
   Menggunakan `innerHTML`
   Menggunakan `document.write()` (hanya untuk testing)
   Menggunakan `window.alert()`
   Menggunakan `console.log()` (untuk debugging)

## Statements

1. JavaScript memperbolehkan tidak menggunakan `;` pada akhir statement, tapi dianjurkan untuk digunakan

## Variables

1. Variable yang sudah di-declare dan di-assign, ketika di-declare lagi, value-nya tidak hilang
2. Variable yang menggunakan keyword `_` digunakan untuk menandakan private variable

## Operators

1. Karena JavaScript dynamic, untuk membandingkan value sekaligus tipe dari variable menggunakan `===`

## Data Types

1. Dalam JavaScript, `null` termasuk object

## Events

1. Beberapa events dalam HTML
   `onchange`
   `onclick`
   `onmouseover`
   `onmouseout`
   `onkeydown`
   `onload`
2. Contoh penggunaan
   `<button onclick="document.getElementById('demo').innerHTML = Date()">The time is?</button>`

## Strings

1. Untuk memasukkan petik dalam string yang terdapat petik yang sama menggunakan escape character
   `const x = "We are the so-called \"Vikings\" from the north.";`
   `const x = 'It\'s alright.';`

## Numbers

1. Angka tanpa koma atau eksponen, hanya akurat sampai 15 digit
2. Maksimum angka desimal adalah 17 digit
3. Angka di dalam string bisa melakukan operasi kurang, kali, atau bagi, tapi tidak bisa operasi tambah
4. Dalam JavaScript, `NaN` termasuk number
5. Dalam JavaScript, `Infinity` termasuk number

## Arrays

1. Dalam JavaScript, array termasuk object (spesial)
2. Dalam array, index menggunakan angka
3. Dalam object, index menggunakan kata
4. Karena array termasuk object, untuk mengetahui apakah sebuah variable itu array menggunakan method `Array.isArray()`
5. Bisa juga menggunakan `namaVariable instanceof Array`

## Sorting Arrays

1. Secara default, array akan dianggap sebagai string
   Maka dari itu, jika array berupa number, hasil `sort()` akan tidak sesuai
   Mengatasinya, di dalam param `sort()` diberikan method untuk meng-compare dua index sebelum di urutkan
2. Tidak ada method untuk mencari nilai terkecil/terbesar dalam array
   Mengatasinya, array di sort terlebih dahulu, index ke-0 menjadi terkecil dan sebaliknya

## Array Iteration Methods

1. `map()` digunakan untuk menciptakan array baru
   `map()` tidak mengeksekusi method dalam index tanpa value
2. `filter()` digunakan untuk menciptakan array baru yang memenuhi kondisi
3. `reduce()` digunakan untuk memadatkan index dalam array menjadi satu
   Secara default memproses dari kiri ke kanan
   Untuk sebaliknya, menggunakan `reduceRight()`
4. `every()` digunakan untuk me-return boolean apakah semua index dalam array memenuhi kondisi
5. `some()` kebalikan dari `every()`
6. `indexOf()` digunakan untuk mencari sebuah index dalam array ada di posisi ke berapa
   Jika tidak ada, akan me-return `-1`
   Jika ada lebih dari satu, akan me-return yang paling pertama
7. `lastIndexOf()` kebalikan dari `indexOf()`
8. `find()` digunakan untuk me-return value pertama dalam array yang memenuhi kondisi
9. `findIndex()` digunakan untuk me-return index pertama dalam array yang memenuhi kondisi

## Dates

1. Penghitungan bulan, jam, dan hari dalam JavaScript di mulai dari `0`
2. Secara otomatis, `Date()` akan di ubah menjadi string

## Date Get Methods

1. Penghitungan hari dalam JavaScript di mulai dari hari `Sunday` (Minggu)
2. Method get yang tersedia dalam objek `Date()` juga dapat menggunakan method format `UTC`

## Date Set Methods

1. Method set yang tersedia dalam objek `Date()` akan me-return dalam format milidetik (lengkap)
2. Jika input param melebihi batas, maka akan disesuaikan secara otomatis

## Random

1. Nilai yang di-return selalu lebih kecil dari `1` (rentang 0 - 1)

## Booleans

1. Variable yang ber-"value" bernilai `true`
2. Variable yang tidak ber-"value" bernilai `false`

## Comparisons

1. JavaScript mendukung `ternary` operator

## Switch

1. Case dalam switch yang memiliki logic sama dapat di-`omit` langsung
2. Switch menggunakan operator `===` untuk perbandingan

## Loop For

1. Jenis loop yang ada dalam JavaScript
   `for`
   `for/in`
   `for/of`
   `while`
   `do/while`
2. Tiga statement dalam loop `for` bisa di-`omit`

## Loop While

1. Loop `while` hampir sama dengan loop `for` (statement pertama dan kedua di-`omit`)

## Break Continue

1. `break` digunakan untuk keluar dari
2. `continue` digunakan untuk loncat ke awal

## Type Conversion

1. Data Types
   `string`
   `number`
   `boolean`
   `object`
   `function`
   `null`
   `undefined`
2. Object Types
   `Object`
   `Date`
   `Array`
   `String`
   `Number`
   `Boolean`
3. `NaN` -> number
   `Array` -> objek
   `Date` -> objek
   `Null` -> objek
   Variable yang belum di-definisikan -> objek
   Variable yang belum di-assign -> objek

## Errors

1. Pattern try/catch
   `try` -> `catch` -> `throw` -> `finally`
2. Error dalam blok `catch` memiliki 2 property, `name` dan `message`
3. Tipe error
   `EvalError` (versi lama)
   `RangeError`
   `ReferenceError`
   `SyntaxError` (versi baru)
   `TypeError`
   `URIError`

## Strict Mode

1. `"use strict"` digunakan untuk mengeksekusi kode dalam mode `strict`
   Artinya program tidak bisa menggunakan variable yang belum di-declare dan beberapa hal lainnya
   Tujuannya untuk mengatasi kelemahan fitur `hoisting` dan membuat program lebih aman

## this Keyword

1. `this` menandakan objek yang dimilikinya

## Let

1. Bisa di-declare terlebih dahulu & di-assign kemudian dan bisa di-assign kembali

## Const

1. Harus di-declare & di-assign terlebih dahulu dan tidak bisa di-assign kembali
2. Khusus untuk object dan array, property/element-nya bisa di ubah dan tambah

## Arrow Function

1. Digunakan untuk mempersingkat pendefinisian function

## Debugging

1. `debugger` digunakan untuk memberhentikan proses sebelum lanjut ke baris selanjutnya

## Style Guide

1. Maksimal 80 character, jika lebih break setelah operator atau koma
   `document.getElementById("demo").innerHTML =`
   ` "Hello, World!"`

## Best Practices

1. Disarankan untuk menghindari variable global, keyword `new`, dan method `eval()`
2. Disarankan untuk langsung menginisialisasi variable
3. Gunakan operator `===`
4. Tambahkan default parameter agar tidak `undefined`
5. Gunakan keyword `default` dalam `switch` walaupun yakin tidak digunakan

## Performance

1. Pindahkan deklarasi variable dalam loop `for` ke luar
2. Jangan simpan value yang hanya digunakan sekali dalam variable
3. Hindari penggunaan keyword `with`

## JSON

1. Key dalam JSON harus menggunakan double quote

## Object Definitions

1. Hampir semua hal dalam JavaScript adalah objek (kecuali tipe data primitive)
2. JSON dalam JavaScript sama seperti:
   `Associative arrays` di PHP
   `Dictionaries` di Python
   `Hash` tables di C
   `Hash maps` di Java
   `Hashes` di Ruby dan Perl

## Last Chapter

https://www.w3schools.com/js/js_object_display.asp
