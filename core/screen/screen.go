package screen

import (
	"fmt"
	"os"
	"os/exec"
)

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
