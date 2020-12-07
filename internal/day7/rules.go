package day7

type Rules []Rule

type Rule struct {
	Name     string
	Quantity int
	Children []Rule
}
