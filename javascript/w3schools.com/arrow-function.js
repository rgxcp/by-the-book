const helloBefore = function () {
  return "Hello, World!";
};

const helloAfter = () => {
  return "Hello, World!";
};

// hanya untuk function yang memiliki satu statement
const helloAfterShorter = () => "Hello, World!";

// dengan param
const sayHello = (name) => `Hello, ${name}`;
