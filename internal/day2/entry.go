package day2

type Entry struct {
	Policy   Policy
	Password string
}

type Policy struct {
	Minimum int
	Maximum int
	Letter  rune
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

func (e Entry) IsPositionallyValid() bool {
	first := rune(e.Password[e.Policy.Minimum-1]) == e.Policy.Letter
	second := rune(e.Password[e.Policy.Maximum-1]) == e.Policy.Letter
	return first != second
}
