package day6

import (
  "fmt"
  "strings"

  "advent/internal/input"
)

func Solve() {
  data := input.NewInput("./input/day6.txt").AsStrings()
  groups := splitIntoGroups(data)
  fmt.Println("Part 1 Solution was", sumUniqueAnswers(groups))
}

func splitIntoGroups(answers []string) [][]string {
  groups := [][]string{}
  groupIndex := 0
  for _, v := range answers {
    if v == "" {
      groupIndex++
      continue
    }
    if len(groups) <= groupIndex {
      groups = append(groups, []string{v})
    } else {
    groups[groupIndex] = append(groups[groupIndex], v)
    }
  }
  return groups
}

func sumUniqueAnswers(groups [][]string) int {
  sum := 0
  for _, v := range groups {
    sum = sum + countUniqueAnswersInGroup(v)
  }
  return sum
}

func countUniqueAnswersInGroup(group []string) int {
  flattened := strings.Join(group, "")
  alphabet := map[rune]bool{}
  for _,v := range flattened {
    alphabet[v]=true
  }
  count := 0
  for _, v := range alphabet {
    if v == true {
      count++
    }
  }
  return count
}