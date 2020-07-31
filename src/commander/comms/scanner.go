package comms

import (
	"bufio"
	"os"
)

func Scn() string{
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	s := scanner.Text()

	return s
}