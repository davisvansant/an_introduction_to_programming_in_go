#### Structs and Interfaces

- What's the difference between a method and a function?
```
A function is an independent section of code that maps zero or more input parameters to zero or more output parameters
A method is a function with a special receiver argument
The receiver is like a parameter – it has a name and a type – but by creating the function in this way it allows us to call the function using the . operator
```

- Why would you use an embedded anonymous field instead of a normal named field?
```
Its a good way say an the specific struct is a specific type, rather than the struct has that type
```

- Add a new method to the Shape interface called perimeter which calculates the perimeter of a shape. Implement the method for Circle and Rectangle.
```
shapes.go
```
