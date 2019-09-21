const fs = require("fs");
const path = require("path");
const filePath = path.join(__dirname, "input.txt");

const file = fs.readFileSync(filePath, { encoding: "utf-8" });
fileLines = file.split("\n");
numbers = fileLines[0].split(" ").map(value => parseInt(value));
targetSum = parseInt(fileLines[1]);

const result = (numbers, targetSum) => {
  const difference = new Set();
  for (var i = 0; i < numbers.length; i++) {
    if (difference.has(numbers[i])) {
      return true;
    }
    difference.add(targetSum - numbers[i]);
  }
  return false;
};

console.log(
  `Any 2 numbers from list add up to k: ${result(numbers, targetSum)}`
);
