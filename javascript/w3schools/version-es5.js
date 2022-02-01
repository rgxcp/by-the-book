// Parse
const text = '{"name":"John", "age":30, "city":"New York"}';
const textJSON = JSON.parse(text);

// Stringify
const textJSON2 = { name: "John", age: 30, city: "New York" };
const text2 = JSON.stringify(textJSON2);

// Getter
const person = {
  firstName: "John",
  lastName: "Doe",
  year: 1999,
  get fullName() {
    return this.firstName + " " + this.lastName;
  },
  get age() {
    return this.year;
  },
  set age(value) {
    this.year = value;
  },
};
console.log(person.fullName);

// Setter
person.age = 21;
console.log(person.age);
