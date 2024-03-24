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
	Command     string
}

func New() *target {
	return &target{}
}
func (t target) GetMakefileTargets(makefileName string) ([]*target, error) {
	var targets []*target
	var currentCommandLines []string

	readFile, err := os.Open(makefileName)
	if err != nil {
		return nil, fmt.Errorf("target: could not open file '%s': %w", makefileName, err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lastTarget *target
	for fileScanner.Scan() {
		line := fileScanner.Text()

		// Check for target line
		before, after, found := strings.Cut(line, ":")
		if found {
			name := strings.TrimSpace(before)

			// skips .PHONY and similar
			if strings.HasPrefix(name, ".") {
				continue
			}

			description := ""
			if descIndex := strings.Index(after, "#"); descIndex != -1 {
				description = strings.TrimSpace(after[descIndex+1:])
			}

			// Save any pending commands to the last target
			if lastTarget != nil {
				lastTarget.Command = strings.Join(currentCommandLines, "\n")
				currentCommandLines = []string{} // Reset command lines
			}

			target := &target{
				Name:        name,
				Description: description,
			}

			targets = append(targets, target)
			lastTarget = target // Update last target to the new one
		} else if lastTarget != nil && strings.HasPrefix(line, "\t") { //commands are indented with a tab
			// Collect command lines for the last target
			currentCommandLines = append(currentCommandLines, strings.TrimSpace(line))
		}
	}

	if lastTarget != nil {
		lastTarget.Command = strings.Join(currentCommandLines, "\n")
	}

	return targets, nil
}

// GetMakefileTargets reads a Makefile and returns a slice of targets
// func (t target) GetMakefileTargets(makefileName string) ([]*target, error) {
// 	var targets []*target
//
// 	readFile, err := os.Open(makefileName)
// 	if err != nil {
// 		return nil, fmt.Errorf("target: could not open file '%s': %w", makefileName, err)
// 	}
// 	defer readFile.Close()
//
// 	fileScanner := bufio.NewScanner(readFile)
// 	fileScanner.Split(bufio.ScanLines)
//
// 	for fileScanner.Scan() {
// 		line := fileScanner.Text()
// 		before, after, found := strings.Cut(line, ":")
//
// 		if found {
// 			name := strings.TrimSpace(before)
//
// 			// skips .PHONY
// 			if strings.HasPrefix(name, ".") {
// 				continue
// 			}
//
// 			_, description, _ := strings.Cut(after, "#")
//
// 			target := &target{
// 				Name:        name,
// 				Description: strings.TrimSpace(description),
// 			}
//
// 			targets = append(targets, target)
// 		}
// 	}
//
// 	return targets, nil
// }
