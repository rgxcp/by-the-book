const cars = ["Tesla", "BMW", "Lexus"];
console.log(cars.sort()); // mengurutkan array dari A-z
console.log(cars.reverse()); // mengurutkan array dari Z-a

const points = [40, 100, 1, 5, 25, 10];
console.log(
  points.sort(function (a, b) {
    return a - b;
  })
); // mengurutkan array dari angka terkecil
console.log(
  points.sort(function (a, b) {
    return b - a;
  })
); // mengurutkan array dari angka terbesar
console.log(
  points.sort(function (a, b) {
    return 0.5 - Math.random();
  })
); // mengurutkan array secara acak
for (i = points.length - 1; i > 0; i--) {
  j = Math.floor(Math.random() * i);
  k = points[i];
  points[i] = points[j];
  points[j] = k;
} // mengurutkan array secara acak menggunakan algoritma Fisher Yates

function myArrayMin(array) {
  var len = array.length;
  var min = Infinity;
  while (len--) {
    if (array[len] < min) {
      min = array[len];
    }
  }
  return min;
}
function myArrayMax(array) {
  var len = array.length;
  var max = -Infinity;
  while (len--) {
    if (array[len] > max) {
      max = array[len];
    }
  }
  return max;
}
// mencari nilai terkecil/terbesar dengan mengakali `Math`
console.log(Math.min.apply(null, points));
console.log(Math.max.apply(null, points));

// mencari nilai terkecil/terbesar dengan menggunakan method sendiri (lebih efisien)
console.log(myArrayMin(points));
console.log(myArrayMax(points));
