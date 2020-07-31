package comms

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func Scn() string{
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	s := scanner.Text()

	strim := strings.TrimSpace(s)

	fmt.Println(strim)
	return s
}