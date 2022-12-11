package client

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/halpdesk/mud/core/screen"
)

type Client struct {
	lastInput string
}

func New() Client {
	return Client{}
}

func (c *Client) LastInput() string {
	return c.lastInput
}

func (c *Client) Input() string {
	fmt.Print(screen.Color("> ", screen.BlueFg, screen.BlackBg))
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Could not read that: %s", err.Error())
	}
	// strip and endl
	text = strings.ReplaceAll(text, "\n", "")
	c.lastInput = text
	return text
}
