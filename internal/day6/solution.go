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
  fmt.Println("Part 2 Solution was", sumCombinedAnswers(groups))
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

func sumCombinedAnswers(groups [][]string) int {
  sum := 0
  for _, v := range groups {
    sum = sum + countCombinedAnswersInGroup(v)
  }
  return sum
}

func countCombinedAnswersInGroup(group []string) int {
  total := 0
  targetQuestions := group[0]
  for _, answer := range targetQuestions {
    allContained := true
    for i := 1; i < len(group); i++ {
      if !strings.Contains(group[i], string(answer)) {
        allContained = false
        break
      }
    }
    if allContained { total++ }
  }
  return total
}
