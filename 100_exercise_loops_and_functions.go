package main

import "fmt"

func Sqrt(x float64) float64 {
  approximation := func(z, x float64) float64 {
    return z - ((z * z) - x) / (2 * z)
  }

  z := 1.0
  for i := 0; i  < 10; i++ {
    z = approximation(z, x)
  }

  return z
}

func main() {
  fmt.Println(Sqrt(16))
}
