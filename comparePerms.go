package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/Ramshackle-Jamathon/go-quickPerm"
	"github.com/ernestosuarez/itertools"
)

func main() {

	input := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "a", "b"}

	//1. itertools
	counter := 0
	fmt.Println("[+] Running test for github.com/ernestosuarez/itertools...")
	start := time.Now()
	for _ = range itertools.PermutationsStr(input, len(input)) {
		//fmt.Println(p)
		counter++
	}
	elapsed := time.Since(start)
	fmt.Printf(" \\--> Finished. Took %s to calculate %d permutations.\n\n", elapsed, counter)

	//2. quickperm
	counter = 0
	start = time.Now()
	fmt.Println("[+] Running test for github.com/Ramshackle-Jamathon/go-quickPerm...")
	for _ = range quickPerm.GeneratePermutationsString(input) {
		//fmt.Println(p)
		counter++
	}
	elapsed = time.Since(start)
	fmt.Printf(" \\--> Finished. Took %s to calculate %d permutations.\n\n", elapsed, counter)

	//Test Python's itertools
	start = time.Now()
	fmt.Println("[+] Running test for Python's itertools...")

	cmd := exec.Command("./testPy.py")
	out, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}
	elapsed = time.Since(start)
	fmt.Printf(" \\--> Finished. Took %s to calculate %s permutations.\n\n", elapsed, strings.TrimSuffix(string(out), "\n"))

}
