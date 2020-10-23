package echo

import (
	"fmt"
	"os"
	"strconv"
)

func Echo1() {
	var s string
	var p string
	for i := 0; i < len(os.Args); i++ {
		s = " " + os.Args[i]
		p = strconv.Itoa(i)
		fmt.Println(p + s)
		// fmt.Printf("index: %d value: %s", i, s)
	}
}
