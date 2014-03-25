package main

import "fmt"

func main() {
  // var a [2]string
  // a[0] = "Hello"
  // a[1] = "World"
  // fmt.Println(a[0], a[1])

  // Setting entries as array is declared
  // a:= [2]string{"Hello", "World!"}

  // Implicit length
  a:= [...]string{"Hello", "World!"}

  fmt.Printf("%q\n", a)
}
