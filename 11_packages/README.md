#### Packages

- Why do we use packages?
```
It reduces the chance of having overlapping names. This keeps our function names short and succinct
It organizes code so that its easier to find code you want to reuse.
It speeds up the compiler by only requiring recompilation of smaller chunks of a program. Although we use the package fmt, we don't have to recompile it every time we change our program.
```

- What is the difference between an identifier that starts with a capital letter and one which doesn’t? (Average vs average)
```
In Go if something starts with a capital letter that means other packages (and programs) are able to see it
```

- What is a package alias? How do you make one?
```go
Shorter name for a package, preventing potential import conflicts
import m "path_to_package"
```

- We copied the average function from chapter 7 to our new package. Create Min and Max functions which find the minimum and maximum values in a slice of float64s.
```
math.go
```

- How would you document the functions you created in #3?
```
adding a comment before the function
```
