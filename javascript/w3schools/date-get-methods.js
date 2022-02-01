const date = new Date();

console.log(`Full Year: ${date.getFullYear()}`);
console.log(`Month: ${date.getMonth()}`); // 0 - 11
console.log(`Date: ${date.getDate()}`); // 1 - 31
console.log(`Hours: ${date.getHours()}`); // 0 - 23
console.log(`Minutes: ${date.getMinutes()}`); // 0 - 59
console.log(`Seconds: ${date.getSeconds()}`); // 0 - 59
console.log(`Milliseconds: ${date.getMilliseconds()}`); // 0 - 999
console.log(`Time: ${date.getTime()}`); // menghitung jumlah milidetik sejak 1 Januari 1970
console.log(`Day: ${date.getDay()}`); // 0 - 6
console.log(`Now: ${Date.now()}`);
