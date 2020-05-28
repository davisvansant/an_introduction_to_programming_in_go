package main

import (
  "fmt"
  "math"
)

type Circle struct {
  x, y, r float64
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

type Shape interface {
  area() float64
  perimeter() float64
}

type MultiShape struct {
  shapes []Shape
}

func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

func (c *Circle) perimeter() float64 {
  return 2 * math.Pi * c.r
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

func (r *Rectangle) perimeter() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return 2 * l + 2 * w
}

func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
    area += s.area()
  }
  return area
}

func (m *MultiShape) area() float64 {
  var area float64
  for _, s := range m.shapes {
    area += s.area()
  }
  return area
}

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

func main()  {
  c := Circle{x: 0, y: 0, r: 5}
  fmt.Println("Circle area is", c.area())
  r := Rectangle{0, 0, 10, 10}
  fmt.Println("Rectangle area is", r.area())
  fmt.Println("Total area is", totalArea(&c, &r))
  fmt.Println("Circle perimeter is", c.perimeter())
  fmt.Println("Rectangle perimeter is", r.perimeter())
}
