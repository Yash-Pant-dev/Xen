package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"xen/xeni"
)

/* CL Tool for the Xen Interpreter. */

const CLToolVersion = 1.0

func main() {

	fmt.Println("Xen CLTool v", CLToolVersion)

	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("[E] ", err)
			panic(fmt.Sprintf("[%v] %v", E_CLTool_ReadLine, err))
		}

		parse(input)
	}
}

func parse(input string) {

	elements := strings.Fields(input)

	switch elements[0] {
	case "r": // run interpreter
		xeni.OrchestrateStages(elements[1:])
	default:
		panic(fmt.Sprintf("[%v]", E_CLTool_UnknownOperation))
	}
}
