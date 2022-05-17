package main

import (
	"fmt"
	"strconv"

	"github.com/syndagate/cliutils/color"
)

func main() {
	fmt.Println("\x1b[" + strconv.Itoa(color.FgRed) + "m" + "test" + "\x1b[0m")
	fmt.Println("\x1b[31mtest\x1b[0m")
}
