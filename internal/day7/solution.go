package day7

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"advent/internal/input"
)

func Solve() {
	data := input.NewInput("./input/day7.txt").AsStrings()
	rules := buildRules(data)
	fmt.Println("Part 1 Solution was", countCanContain("shiny gold", rules))
}

func countCanContain(name string, rules Rules) int {
	sum := 0
	for _, rule := range rules {
		if canContain(name, rule, rules) {
			sum++
		}
	}
	return sum
}

func canContain(name string, rule Rule, rules Rules) bool {
	for _, child := range rule.Children {
		if name == child.Name {
			return true
		}
		for _, childRule := range rules {
			if childRule.Name == child.Name {
				if canContain(name, childRule, rules) {
					return true
				}
				break
			}
		}
	}
	return false
}

func buildRules(data []string) Rules {
	childMatch := regexp.MustCompile(`(\d+) (.*) bags?`)
	rules := []Rule{}
	for _, line := range data {
		split := strings.Split(line, "bags contain")
		rule := Rule{
			Name:     strings.Trim(split[0], " "),
			Quantity: 1,
		}
		children := []Rule{}
		childrenStr := strings.Split(split[1], ",")
		for _, childStr := range childrenStr {
			matches := childMatch.FindAllSubmatch([]byte(childStr), -1)
			if len(matches) == 0 {
				continue
			}
			quantity, err := strconv.ParseInt(string(matches[0][1]), 10, 64)
			if err != nil {
				log.Fatalf("Converting %s to int: %s", string(matches[0][1]), err)
			}
			childRule := Rule{
				Name:     strings.Trim(string(matches[0][2]), " "),
				Quantity: int(quantity),
			}
			children = append(children, childRule)
		}
		rule.Children = children
		rules = append(rules, rule)
	}
	return rules
}
