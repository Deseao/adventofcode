package input

import (
  "bufio"
	"log"
	"os"
)

type Input struct {
  filepath string
}

func NewInput(file string) (*Input) {
 return &Input{filepath: file,}
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