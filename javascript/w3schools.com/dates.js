const date = new Date();

console.log(`Original: ${date}`);
console.log(`UTC Format: ${date.toUTCString()}`);
console.log(`ISO Format: ${date.toISOString()}`);
console.log(`Readable Format: ${date.toDateString()}`);
