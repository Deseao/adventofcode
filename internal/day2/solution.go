package day2

import (
  "errors"
  "fmt"
  "log"
  "strings"
  "strconv"

  "advent/internal/input"
)

const inputSeparator string = ":"
const policySeperator string = " "
const limitsSeperator string = "-"

func Solve() {
  data := input.NewInput("./input/day2.txt").AsStrings()
  entries, err := ParseInput(data)
  if err != nil {
    log.Fatalf("failed parsing input: %s", err)
  }
  fmt.Printf("Part 1 Solution was %d\n", sumValidEntries(entries)) 
  fmt.Printf("Part 2 Solution was %d\n", sumPositionallyValidEntries(entries))
}

func sumValidEntries(entries []Entry) int {
  sum := 0
  for _, v := range entries {
    if v.IsValid() { sum++ }
  }
  return sum
}

func sumPositionallyValidEntries(entries []Entry) int {
  sum := 0
  for _, v := range entries {
    if v.IsPositionallyValid() { sum++ }
  }
  return sum
}

func ParseInput(input []string) (entries []Entry, err error) {
  entries = make([]Entry, len(input))
  for i, v := range input {
    policyInput, password, err := splitInput(v)
    if err != nil {
      return nil, fmt.Errorf("splitting %s: %w", v, err)
    }
    entry := Entry{
      Password:password,
    }

    policy, err := buildPolicy(policyInput)
    if err != nil {
      return nil, fmt.Errorf("building policy %s: %w", policyInput, err)
    }
    entry.Policy = policy

    entries[i] = entry
  }
  return entries, nil
}

func splitInput(input string) (policy, password string, err error) {
  results := strings.Split(input, inputSeparator)
  if len(results) != 2 {
    return "", "", errors.New("splitting policy and password")
  }
  return results[0], strings.Trim(results[1], " "), nil
}

func buildPolicy(input string) (policy Policy, err error) {
  results := strings.Split(input, policySeperator)
  if len(results) != 2 {
    return Policy{}, errors.New("splitting policy into range and letter")
  }
  policy = Policy{
    Letter: []rune(results[1])[0],
  }
  limits := strings.Split(results[0], limitsSeperator)
  if len(limits) != 2 {
    return Policy{}, errors.New("splitting limits by -")
  }
  min, err := strconv.ParseInt(limits[0], 10, 64)
  if err != nil {
    return Policy{}, fmt.Errorf("parsing %s as an int: %w", limits[0], err)
  }
  policy.Minimum = int(min)
  max, err := strconv.ParseInt(limits[1], 10, 64)
  if err != nil {
    return Policy{}, fmt.Errorf("parsing %s as an int: %w", limits[1], err)
  }
  policy.Maximum = int(max)
  return policy, nil
}