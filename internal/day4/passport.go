package day4

import (
  "log"
  "strconv"
  "strings"
  "regexp"
)

type Passport struct {
  BirthYear string  `yaml:"byr"`
  IssueYear string  `yaml:"iyr"`
  ExpirationYear string `yaml:"eyr"`
  Height string `yaml:"hgt"`
  HairColor string `yaml:"hcl"`
  EyeColor string `yaml:"ecl"`
  ID string `yaml:"pid"`
  Country string `yaml:"cid"`
}

func (p Passport) IsValid() bool {
  if p.BirthYear == "" || p.IssueYear == "" || p.ExpirationYear == "" || p.Height == "" || p.HairColor == "" || p.EyeColor == "" || p.ID == "" {
    return false
  }
  return true
}

func (p Passport) IsStrictValid() bool {
  if !p.IsValid() { return false } //All required fields are still required
  if len(p.BirthYear) != 4 { return false}
  birthYear, err := strconv.ParseInt(p.BirthYear, 10, 64)
  if err != nil {
    log.Fatalf("converting birth year to int: %s", err)
  }
  if birthYear < 1920 || birthYear > 2002 { return false }
  if len(p.IssueYear) != 4 {return false}
  issueYear, err := strconv.ParseInt(p.IssueYear, 10, 64)
  if err != nil {
    log.Fatalf("converting issue year to int: %s", err)
  }
  if issueYear < 2010 || issueYear > 2020 { return false }
  if len(p.ExpirationYear) != 4 {return false}
  expirationYear, err := strconv.ParseInt(p.ExpirationYear, 10, 64)
  if err != nil {
    log.Fatalf("converting expiration year to int: %s", err)
  }
  if expirationYear < 2020 || expirationYear > 2030 { return false }

  if !strings.HasSuffix(p.Height, "cm") && !strings.HasSuffix(p.Height, "in") { return false}
  if strings.HasSuffix(p.Height, "cm") {
    height, err := strconv.ParseInt(p.Height[:len(p.Height)-2], 10, 64)
    if err != nil {
      log.Fatalf("converting cm height to int: %s", err)
    }
    if height < 150 || height > 193 { return false}
  }
  if strings.HasSuffix(p.Height, "in") {
    height, err := strconv.ParseInt(p.Height[:len(p.Height)-2], 10, 64)
    if err != nil {
      log.Fatalf("converting in height to int: %s", err)
    }
    if height < 59 || height > 76 { return false}
  }
  hairColorRegex := regexp.MustCompile(`#[0-9a-f]{6}`)
  if !hairColorRegex.MatchString(p.HairColor) { return false }
  validEyeColors := []string{"amb","blu", "brn", "gry", "grn", "hzl", "oth"}
  if !contains(validEyeColors, p.EyeColor) { return false }
  if len(p.ID) != 9 { return false}

  return true
}

func contains(vals []string, target string) bool {
  for _, v := range vals {
    if v == target { return true}
  }
  return false
}