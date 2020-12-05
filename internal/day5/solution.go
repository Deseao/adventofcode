package day5

import (
  "fmt"

  "advent/internal/input"
)

func Solve() {
  data := input.NewInput("./input/day5.txt").AsStrings()
  passes := convertToBoardingPass(data)
  fmt.Println("Day 5 Solution was", highestSeatID(passes))
}

func convertToBoardingPass(passes []string) []BoardingPass {
  result := make([]BoardingPass, len(passes))
  for _, v := range passes {
    result = append(result, NewBoardingPass(v))
  }
  return result
}

func highestSeatID(passes []BoardingPass) int {
  max := 0
  for _, v := range passes {
    if v.ID() > max {
      max = v.ID()
    }
  }
  return max
}