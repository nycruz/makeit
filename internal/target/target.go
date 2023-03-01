package target

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type target struct {
	Name        string
	Description string
}

func NewTarget() *target {
	return &target{}
}

func (t target) GetMakefileTargets(makefileName string) ([]*target, error) {
	var targets []*target

	readFile, err := os.Open(makefileName)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %v", makefileName, err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		before, after, found := strings.Cut(line, ":")

		if found {
			name := strings.TrimSpace(before)

			if strings.HasPrefix(name, ".") {
				continue // skips .PHONY
			}

			_, description, _ := strings.Cut(after, "#")

			target := &target{
				Name:        name,
				Description: strings.TrimSpace(description),
			}

			targets = append(targets, target)
		}
	}

	return targets, nil
}