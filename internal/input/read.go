package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Input struct {
	filepath string
}

func NewInput(file string) *Input {
	return &Input{filepath: file}
}

func (i *Input) AsStrings() []string {
	file, err := os.Open(i.filepath)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	return txtlines
}

func (i *Input) AsInts() []int {
	strs := i.AsStrings()
	var result []int
	for _, str := range strs {
		val, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Fatalf("Error converting string %s to int:%s", str, err)
		}
		result = append(result, int(val))
	}
	return result
}
