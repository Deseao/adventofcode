package main

import (
  "fmt"

  "advent/internal/day1"
)

func main() {
  a, b, c, err := day1.FindSumTriplet(day1.ExampleInput, day1.TargetYear)
  if err != nil {
    panic(err)
  }
  result := a * b * c
  fmt.Printf("Solution was %d\n", result)
}