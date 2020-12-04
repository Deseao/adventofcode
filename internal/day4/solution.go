package day4

import (
  "fmt"
  "strings"
  "log"
  "regexp"
  
  "gopkg.in/yaml.v2"

  "advent/internal/input"
)

func Solve() {
  data := input.NewInput("./input/day4.txt").AsStrings()
  passports, err := splitSliceIntoPassports(data)
  if err != nil {
    log.Fatalf("converting data into passports:", err)
  }
  fmt.Println("Part 1 Solution was", countValid(passports))
  fmt.Println("Part 2 Solution was", countStrictlyValid(passports))
}

func splitSliceIntoPassports(data []string) ([]Passport, error) {
  passports := []Passport{}
  passport := []string{}
  for _, v := range data {
    if v == "" {
      result, err := passportStringsToPassport(passport)
      if err != nil {
        return []Passport{}, fmt.Errorf("Converting passport %s: %w", passport, err)
      }
      passports = append(passports, result)
      passport = passport[:0]
      continue
    }
    passport = append(passport, v)
  }
  result, err := passportStringsToPassport(passport)
      if err != nil {
        return []Passport{}, fmt.Errorf("Converting passport %s: %w", passport, err)
      }
      passports = append(passports, result)
  return passports, nil
}

func passportStringsToPassport(data []string) (Passport, error) {
  flattened := strings.Join(data[:], " ")
  //replace spaces with newlines and add spaces after colons to make it yaml parseable
  newLined := strings.ReplaceAll(flattened, " ", "\n")
  spaced := strings.ReplaceAll(newLined, ":", ": ")
  //Wrap all # symbols in quotes so they're not treated as a comment
  uncommented := fixComments(spaced)

  var pssprt Passport
 
  err := yaml.Unmarshal([]byte(uncommented), &pssprt)
  if err != nil {
    return Passport{}, fmt.Errorf("unmarshalling yaml: %w", err)
  }
  return pssprt, nil
}

func fixComments(input string) string {
  comments := regexp.MustCompile(`(#\S*)`)
  return comments.ReplaceAllString(input, "\"$1\"\n")
}

func countValid(passports []Passport) int {
  total := 0
  for _, pass := range passports {
    if pass.IsValid() { total++ }
  }
  return total
}

func countStrictlyValid(passports []Passport) int {
   total := 0
  for _, pass := range passports {
    if pass.IsStrictValid() { total++ }
  }
  return total
}

