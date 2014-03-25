package main

import "fmt"

func main() {
  slice := []int{2, 3, 5, 7, 11, 13}

  fmt.Println(slice)
  fmt.Println(slice[1:4])
  fmt.Println(slice[:3])
  fmt.Println(slice[5:])
}
