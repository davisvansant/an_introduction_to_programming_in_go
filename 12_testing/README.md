#### Testing

- Writing a good suite of tests is not always easy, but the process of writings tests often reveals more about a problem then you may at first realize. For example, with our Average function what happens if you pass in an empty list ([]float64{})? How could we modify the function to return 0 in this case?
  ```go
  if len(xs) == 0 {
    return 0
  }
  ```

- Write a series of tests for the Min and Max functions you wrote in the previous chapter.
```
../11_packages/math/math_test.go
```
