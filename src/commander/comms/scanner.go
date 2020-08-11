package comms

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func Scn() string{
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	s := scanner.Text()

	strim := strings.TrimSpace(s)

	//lazy implementation for a global shutdown
	if strim == "shutdown" {
		fmt.Println("Goodbye!")
		os.Exit(0)
	}

	return strim
}