package day3

import (
  "fmt"

  "advent/internal/input"
)

func Solve() {
  data := input.NewInput("./input/day3.txt").AsStrings()
  fmt.Println("Part 1 Solution was", countTrees(data, 3, 1))
  fmt.Println("Part 2 Solution was", checkOtherSlopes(data))
}

const treeIndicator = '#'

func countTrees(terrain []string, slopeRight int, slopeDown int) int {
  count := 0
  index := 0
  for row := 0; row < len(terrain); row = row + slopeDown {
    if row == 0 {
      continue
    }
    path := terrain[row]
    index = index + slopeRight //Move over
    if  index > len(path) - 1 { 
      index = index - len(path)
    } //rotate back around
    if isTree(rune(path[index])) { 
      count++
    }
  }
  return count
}

func isTree(char rune) bool {
  return char == treeIndicator
}

func checkOtherSlopes(terrain []string) int {
  trees1 :=  countTrees(terrain, 1, 1)
  trees3 := countTrees(terrain, 3, 1)
  trees5 := countTrees(terrain, 5, 1)
  trees7 := countTrees(terrain, 7, 1)
  treesDown :=  countTrees(terrain, 1, 2)
  return trees1 * trees3 * trees5 * trees7 * treesDown
}

