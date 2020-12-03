package day2

type Entry struct {
  Policy Policy
  Password string
}

type Policy struct {
  Minimum int
  Maximum int
  Letter rune
}

func (e Entry) IsValid() bool {
  seen := 0
  for _, v := range e.Password {
    if v == e.Policy.Letter {
      seen++
    }
    if seen > e.Policy.Maximum {
      return false
    }
  }
  if seen < e.Policy.Minimum {
    return false
  }
  return true
}