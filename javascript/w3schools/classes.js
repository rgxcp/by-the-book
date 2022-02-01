class Person {
  constructor(name, sex, dob) {
    this.name = name;
    this.sex = sex;
    this.dob = dob;
  }

  greets(lastName) {
    return `Hello, ${this.name} ${lastName}!`;
  }
}

const person = new Person("Rommy", "Male", "01-01-1999");
console.log(person.greets("Gustiawan"));
