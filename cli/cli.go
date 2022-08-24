package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IInputOutput interface {
	ReadString(msg string) string
	ReadInt(msg string) int
	ReadFloat(msg string) float64
	ReadKey()
	Clear()
	Stringfy(data interface{})
}

type CLI struct {
}

func (c *CLI) ReadString(msg string) string {
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	return text

}

func (c *CLI) ReadInt(msg string) int {
reAsk:
	text := c.ReadString(msg)
	i, err := strconv.Atoi(text)
	if err != nil {
		fmt.Printf("Given value is not a valid int %q ", text)
		goto reAsk
	}
	return i

}

func (c *CLI) ReadFloat(msg string) float64 {
reAsk:
	text := c.ReadString(msg)
	i, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Printf("Given value is not a valid float %q ", text)
		goto reAsk
	}
	return i
}

func (c *CLI) ReadKey() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func (c *CLI) Clear() {
	fmt.Print("\033[H\033[2J")
}

func (c *CLI) Stringfy(data interface{}) {
	str, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(string(str))
}
