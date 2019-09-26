// This problem is probably easier to have solution written in dynamically typed language to avoid explicit typing of a statically typed language

function cons(a, b) {
  function pair(f) {
    return f(a, b);
  }
  return pair;
}

function car(c) {
  function first(x, y) {
    return x;
  }
  return c(first);
}

function cdr(c) {
  function second(x, y) {
    return y;
  }
  return c(second);
}

console.log(car(cons(3, 4))); // should be 3
console.log(cdr(cons(3, 4))); // should be 4
