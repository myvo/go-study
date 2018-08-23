package main
import (
  "fmt"
  "math"
)

type Circle struct {
  x, y, r float64
}

func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}
func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}
func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

type Shape interface {
  area() float64
}

func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
    area += s.area()
  }
  return area
}

type MultiShape struct {
  shapes []Shape
}
func (m *MultiShape) area() float64 {
  var area float64
  for _, s := range m.shapes {
    area += s.area()
  }
  return area
}

func main() {
  c := Circle{2,4,6};
  fmt.Println(c.area())

  r := Rectangle{0, 0, 10, 10}
  fmt.Println(r.area())

  fmt.Println(totalArea(&c, &r))

  m := new(MultiShape);
  m.shapes[0] = Circle{2,4,6};
  m.shapes[1] = Rectangle{0, 0, 10, 10};

  fmt.Println(m.area())
}
