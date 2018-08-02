package main

import (
  "fmt"
  "math/rand"
)

func main() {
  x: int;
  y: int;
  x := 3;
  y = 4;
  fmt.Println("x +y ", x + y);
  rand.Seed(50);
  fmt.Println("My favorite number is", rand.Intn(100), rand.Intn(100), rand.Intn(100));
}
