const date = new Date();

console.log(`Full Year: ${date.setFullYear(2021)}`); // param: tahun, bulan (opsional), hari (opsional)
console.log(`Month: ${date.setMonth(0)}`); // param: 0 - 11
console.log(`Date: ${date.setDate(1)}`); // param: 1 - 31
console.log(`Hours: ${date.setHours(12)}`); // param: 0 - 23
console.log(`Minutes: ${date.setMinutes(12)}`); // param: 0 - 59
console.log(`Seconds: ${date.setSeconds(12)}`); // param: 0 - 59
console.log(`Milliseconds: ${date.setMilliseconds(12)}`); // param: 0 - 999
console.log(`Time: ${date.setTime(12)}`); // param: milidetik sejak 1 Januari 1970
