package main

import "fmt"

func fibo(x uint) uint {
  if x <= 0 {
    return 0
  } else if x == 1 {
    return 1
  }
  return fibo(x - 1) + fibo(x - 2)
}

func main() {
  fmt.Println(fibo(10));
}
