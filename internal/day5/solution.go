package day5

import (
  "fmt"
  "sort"

  "advent/internal/input"
)

func Solve() {
  data := input.NewInput("./input/day5.txt").AsStrings()
  passes := convertToBoardingPass(data)
  fmt.Println("Part 1 Solution was", highestSeatID(passes))
  fmt.Println("Part 2 Solution was", findMissingPass(passes))
}

func convertToBoardingPass(passes []string) []BoardingPass {
  result := []BoardingPass{}
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

func findMissingPass(passes []BoardingPass) int {
  ids := []int{}
  for _, pass := range passes { ids = append(ids, pass.ID()) }
  sort.Ints(ids)
  for i := 0; i < len(ids)-1; i++ {
      next := ids[i + 1]
      current := ids[i]
      if((next - current) > 1) {
        return current + 1
      }
  }
  return 0
}