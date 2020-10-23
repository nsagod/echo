package echo

import (
	"fmt"
	"os"
	"strings"
)

// Echo3 SS
func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func Slice() {
	a := make([]string, 0)
	a = append(a, "1", "2", "3", "4")
	a = append(a[0:1], a[1+1:]...)
	c := make([]string, 4)
	copy(c, a)
	c[0] = "a"
	for _, v := range c {
		fmt.Println(v)
	}
	for _, v := range a {
		fmt.Println(v)
	}
}
