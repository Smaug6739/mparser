const { log } = require("console");

console.log = (...a) => {
  log(...a);
  return console;
};

console.log("Hello World %d", "1").log("Hello World 2");
