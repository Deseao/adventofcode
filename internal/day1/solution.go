package day1

import (
  "errors"
)

var ErrNoInput = errors.New("no input")
var ErrNoMatch = errors.New("no match found")
func FindSumPair(input []int, target int) (a int, b int, err error) {
  if len(input) == 0 {
    return 0, 0, ErrNoInput
  }
  //Iterate over the input array
  for i, a := range input {
    //Iterate over the rest of the array from i + 1 onwards
    for j := i + 1; j < len(input); j++ {
      b := input[j]
      //Compare a + b to target
      if a + b == target {
        return a, b, nil
      }
    }
  }
  return 0, 0, ErrNoMatch
}

func FindSumTriplet(input []int, target int) (a int, b int, c int, err error) {
   if len(input) == 0 {
    return 0, 0, 0, ErrNoInput
  }
  //Iterate over the input array
  for i, a := range input {
    //Iterate over the rest of the array from i + 1 onwards
    for j := i + 1; j < len(input); j++ {
      b := input[j]
      for k := j +1; k < len(input); k++ {
        c := input[k]
              //Compare a + b + c to target
        if a + b +c  == target {
          return a, b, c, nil
        }
      }
    }
  }
  return 0, 0, 0, ErrNoMatch
}