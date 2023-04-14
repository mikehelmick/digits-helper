package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mikehelmick/digits-helper/pkg/digits"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 7 {
		log.Fatalf("7 arguments are required `target` and the 6 digits")
	}

	intArgs := make([]int, len(args))
	for i, v := range args {
		converted, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			log.Fatalf("Unable to convert argument: %v reason: %v", v, err)
		}
		intArgs[i] = int(converted)
	}

	s, err := digits.New(intArgs[0], intArgs[1:])
	if err != nil {
		log.Fatalf("unable to initialize solver: %v", err)
	}

	solution, err := s.Solve()
	if err != nil {
		log.Fatalf("unable to solve: %v", err)
	}
	fmt.Printf("*** SOLUTION ***\n%v\n", strings.Join(solution, "\n"))
}
