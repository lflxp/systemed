package main

import (
	"fmt"
	"strings"

	"github.com/lflxp/systemed"
)

func test() {
	systemed := systemed.NewSystemed("systemctl")
	systemed.SetArgs("list-unit-files").SetArgs("--type service").SetArgs("--plain").SetArgs("|grep -vE 'STATE|listed'|tr -s '\n'|awk '{print $1\" \" $2}'")
	rs, err := systemed.Exec()
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(rs))
	for i, x := range strings.Split(string(rs), "\n") {
		if x != "" && x != " " {
			fmt.Println(i, x)
		}

	}

	defer fmt.Println(systemed.Command)
}

// strings.Split("list-unit-files --type service --plain", " "
func main() {
	rs, err := systemed.ListUnitFiles()
	if err != nil {
		panic(err)
	}

	for i, x := range strings.Split(string(rs), "\n") {
		if x != "" && x != " " {
			fmt.Println(i, x)
		}
	}
}
