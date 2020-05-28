#### Functions

- sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?
```
sum.go
```

- Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).
```
true_or_false.go
```

- Write a function with one variadic parameter that finds the greatest number in a list of numbers.
```
greatest.go
```

- Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.
```
make_odd.go
```

- The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).
```
fib.go
```

- What are defer, panic and recover? How do you recover from a run-time panic?
```
Defer schedules a function call to be run after the function completes, often used when resources need to be freed in some way
Panic function causes a run time error
Recover stops the panic and returns the value that was passed to the call to panic
panic_recover.go
```
