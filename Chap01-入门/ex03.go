// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses string.Join. (Section 1.6
// illustrates part of the time package, and Section 11.4 shows how to write
// benchmark tests for systematic performance evaluation.

package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func BenchmarkIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s, sep string
		for _, arg := range os.Args[1:] {
			s += s + arg + sep
			sep = " "
		}
	}
}

func BenchmarkJoin(b *testing.B) {
    for i := 0; i < b.N; i++ {
    	strings.Join(os.Args[1:], " ")
	}
}

// other method
func echo1() string {
	start := time.Now()
	defer func() {
		fmt.Printf("echo1: %v ns\n", time.Since(start).Nanoseconds())
	}()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep +os.Args[i]
		sep = " "
	}
	return s
}

func echo2() string {
	start := time.Now()
	defer func() {
		fmt.Printf("echo2: %v ns\n", time.Since(start).Nanoseconds())
	}()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3() string {
	start := time.Now()
	defer func() {
		fmt.Printf("echo3: %v ns\n", time.Since(start).Nanoseconds())
	}()

	return strings.Join(os.Args[1:], " ")
}

func main() {
	echo1()
	echo2()
	echo3()
}


