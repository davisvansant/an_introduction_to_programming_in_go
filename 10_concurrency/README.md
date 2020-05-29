#### Concurrency

- How do you specify the direction of a channel type?
```
sending - c chan<- string
receiving - c <-chan string
```

- Write your own Sleep function using time.After.
```
sleeper.go
```

- What is a buffered channel? How would you create one with a capacity of 20?
```
A buffered channel is asynchronous; sending or receiving a message will not wait unless the channel is already full
c := make(chan int, 20)
```
