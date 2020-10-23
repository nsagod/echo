package echo

import (
	"fmt"
	"os"
	"strconv"
)

func Echo2() {
	var p = ""
	for t, arg := range os.Args[2:] {
		p = strconv.Itoa(t)
		fmt.Println(p + " " + arg)
	}
}
