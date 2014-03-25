package main

import "fmt"

func main() {
  // Creates an empty slice of length 3
  cities := make([]string, 3)

  // Populate the empty slice
  cities[0] = "Boston"
  cities[1] = "Salem"
  cities[2] = "Newton"

  fmt.Printf("%q\n", cities)
}
