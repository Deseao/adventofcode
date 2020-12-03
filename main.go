package main

import (
  "fmt"
  "errors"
)

var ErrNoInput = errors.New("no input")
var ErrNoMatch = errors.New("no match found")
func main() {
  a, b, err := FindSumPair(puzzleInput, targetYear)
  if err != nil {
    panic(err)
  }
  result := a * b
  fmt.Printf("Solution was %d\n", result)
}

func FindSumPair(input []int, target int) (a int, b int, err error) {
  if len(input) == 0 {
    return 0, 0, ErrNoInput
  }
  //Iterate over the input array
  for i, a := range input {
    //Iterate over the rest of the array from i + 1 onwards
    for j := i + 1; j < len(input); j++ {
      b := input[j]
      fmt.Printf("Comparing %d and %d\n", a, b)
      //Compare a + b to target
      if a + b == target {
        fmt.Printf("Found a match. %d and %d sum to %d\n", a, b, target)
        return a, b, nil
      }
    }
  }
  return 0, 0, ErrNoMatch
}