# imonkey
Sources for monkey's interpreter.

Monkey is a programming language defined in Thorsten Ball's book - Writing an Interpreter in Go.\
This repository follows closely with ideas expressed in the book.\
It aims to also extend the implementation with other experimental features.

## Language Features
* C-like syntax
* Variable bindings
* Integers and booleans
* Arithmetic expressions
* Built-in functions
* First-class and higher-order functions
* Closures
* String data structure
* Array data structure
* Hash map data structure

## Examples

Recursive Fibonacci
```
let fib = fn (n) {
    if (n == 0) {
        return 1;
    } else if (n == 1) {
        return 1;
    } else {
        return fib(n - 1) + fib(n - 2);
    }
};
```

## References
* Writing an Interpreter in Go - https://interpreterbook.com/

