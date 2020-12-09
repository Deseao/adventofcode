package day8

import (
  "fmt"

  "advent/internal/input"
)

func Solve() {
  data := input.NewInput("./input/day8.txt").AsStrings()
  instructions := NewInstructions(data)
  fmt.Println("Part 1 Solution was", instructions.Traverse())
  fmt.Println("Part 2 Solution was", instructions.TraverseTerminal())
}