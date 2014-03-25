package main

import "fmt"

func main() {
  cities := []string{}
  cities = append(cities, "Boston", "Cambridge", "Salem")

  moreCities := []string{"Meriden", "Chesire", "Hamden"}
  cities = append(cities, moreCities...)

  fmt.Printf("%q\n", cities)
}
