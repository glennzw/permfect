//A Go tool for creating permutations of alphanumerics

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ernestosuarez/itertools"
)

var counter int64
var debug = false

// Calculate permutations and write results into channel
func calculatePermutations(lines []string, r int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range itertools.PermutationsStr(lines, r) {
		counter++
		ch <- strings.Join(v, "")
	}
	close(ch)
	if debug == true {
		fmt.Printf(" \\---> Finished calculating permutations of length %d\n", r)
	}
}

//Keep writing to disk until channel closes
func writePermsToDisk(fname string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	//Open file of length of perm
	file, err := os.Create(fname)
	if err != nil {
		fmt.Println("[!] Unable to write to ", fname)
		panic(err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	//Iterate over channel. Will keep consuming until channel is closed by calculatePermutations
	for data := range ch {
		fmt.Fprintln(w, data)
	}
	w.Flush()
	if debug == true {
		fmt.Printf(" \\---> Finished writing to file %s\n", fname)
	}

}

func main() {

	//Parse command line options
	inputFile := flag.String("infile", "", "Input file.")
	outputBase := flag.String("outbase", "perms", "Base for output files.")
	minLength := flag.Int("min", 1, "Minimum permutation length")
	maxLength := flag.Int("max", -1, "Maximum permutation length")
	debugFlag := flag.Bool("debug", false, "Output extra debugging info.")
	flag.Parse()
	if *inputFile == "" {
		fmt.Println("[!] No input file specified. Use -h for help.")
		os.Exit(1)
	}
	debug = *debugFlag

	var wg sync.WaitGroup
	start := time.Now()
	lines, err := readLines(*inputFile)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	if *maxLength == -1 {
		*maxLength = len(lines)
	}

	if *maxLength > len(lines) {
		fmt.Println("[!] Error: max length must be less than or equal to number of items in input file.")
		os.Exit(1)
	}

	fmt.Printf("[+] Calculating permutations from input file '%s' and writing to %d files (output base of '%s')...\n", *inputFile, (*maxLength - *minLength + 1), *outputBase)
	for r := *minLength; r <= *maxLength; r++ {
		wg.Add(2)
		ch := make(chan string)
		go calculatePermutations(lines, r, ch, &wg)
		go writePermsToDisk(strconv.Itoa(r)+*outputBase, ch, &wg)
	}

	//Wait for goroutines to finish
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("[+] Finished. Took %s to calculate %d permutations.\n", elapsed, counter)

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
