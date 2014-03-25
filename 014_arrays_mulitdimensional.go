package main

import "fmt"

func main() {
  var grid [2][3]string

  for row := 0; row < 2; row++ {
    for col := 0; col < 3; col++ {
      grid[row][col] = fmt.Sprintf("row %d - column %d", row + 1, col + 1)
    }
  }

  fmt.Printf("%q", grid)
}
