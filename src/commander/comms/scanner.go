package comms

import (
	"bufio"
	"os"
	"strings"
)

func Scn() string{
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	s := scanner.Text()

	strim := strings.TrimSpace(s)

	return strim
}