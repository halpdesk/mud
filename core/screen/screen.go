package screen

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var nonAlphanumericRegex1 = regexp.MustCompile(`\033\[[0-9];[0-9]{1,2}m`)
var nonAlphanumericRegex2 = regexp.MustCompile(`\033\[[0-9][0-9]m`)

func New() Screen {
	return Screen{}
}

type Screen struct{}

func (s Screen) Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (s Screen) WriteLn(format string, a ...any) {
	line := fmt.Sprintf(format, a...)
	fmt.Printf("%s\n", line)
}

func (s Screen) Write(format string, a ...any) {
	line := fmt.Sprintf(format, a...)
	fmt.Printf("%s", line)
}

func (s Screen) Write80(format string, a ...any) {
	str := fmt.Sprintf(format, a...)
	strs := strings.Split(str, "\n")
	for _, str := range strs {
		lines := chunkString(str, 100)
		fmt.Printf("  %s\n", strings.Join(lines, "\n  "))
	}
}

func chunkString(str string, chunkSize int) []string {
	var lines []string
	words := strings.Split(str, " ")
	length := 0
	line := ""
	for i := 0; i < len(words); i++ {
		length += len(clearString(words[i])) + 1
		if length <= chunkSize {
			line = fmt.Sprintf("%s%s ", line, words[i])
		}
		// fmt.Printf("line: %s [clear: %s] (%d)\n", line, clearString(words[i]), length)
		if length > chunkSize {
			length = 0
			lines = append(lines, line)
			line = ""
			i--
		}
		if i == len(words)-1 {
			lines = append(lines, line)
		}
	}
	return lines
}

func clearString(str string) string {
	return nonAlphanumericRegex1.ReplaceAllString(
		nonAlphanumericRegex2.ReplaceAllString(str, " "),
		" ",
	)
}
