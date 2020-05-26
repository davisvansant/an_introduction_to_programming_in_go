#### 02 - Your First Program

##### Problems

- What is whitespace?
```
Newlines, spaces and tabs, used to make programs easier to read
```

- What is a comment? What are the two ways of writing a comment?
```
Comments are ignored by the Go compiler and are there for your own sake
//
/* */
```

- Our program began with package main. What would the files in the fmt package begin with?
```
package fmt
```

- We used the Println function defined in the fmt package. If we wanted to use the Exit function from the os package what would we need to do?
```
import "os"
os.Exit(exit_code)
```

- Modify the program we wrote so that instead of printing Hello World it prints Hello, my name is followed by your name
