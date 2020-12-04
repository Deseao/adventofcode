package day4

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