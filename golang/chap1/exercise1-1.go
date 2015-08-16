/*
Task1: Write a program that lists all the DNA sequences that encode a given protein sequence
*/

package main

import "fmt"
import "strings"

var codonTable = map[string][]string{
	"stop": []string{"UAA", "UAG", "UGA"},
	"e":    []string{"GAA", "GAG"},
	"d":    []string{"GAU", "GAC"},
	"g":    []string{"GGU", "GGC", "GGA", "GGG"},
	"v":    []string{"GUU", "GUC", "GUA", "GUG"},
	"t":    []string{"ACU", "ACC", "ACA", "ACG"},
	"i":    []string{"AUU", "AUC", "AUA"},
	"h":    []string{"CAU", "CAC"},
	"a":    []string{"GCU", "GCC", "GCA", "GCG"},
	"w":    []string{"UGG"},
	"c":    []string{"UGU", "UGC"},
	"y":    []string{"UAU", "UAC"},
	"s":    []string{"UCU", "UCC", "UCA", "UCG", "AGU", "AGC"},
	"m":    []string{"AUG"},
	"n":    []string{"AAU", "AAC"},
	"r":    []string{"CGG", "CGA", "CGC", "CGU", "AGG", "AGA"},
	"p":    []string{"CCU", "CCC", "CCA", "CCG"},
	"q":    []string{"CAG", "CAA"},
	"f":    []string{"UUU", "UUC"},
	"l":    []string{"UUA", "UUG", "CUU", "CUC", "CUA", "CUG"},
	"k":    []string{"AAA", "AAG"},
}

var peptide = []string{"M", "R", "F", "L"}

var grandList = []string{}

func recurse(prevSeq string, aa []string) int {
	if len(aa) == 0 {
		for _, stopCodon := range codonTable["stop"] {
			finalSeq := strings.Join(append([]string{prevSeq}, stopCodon), " ")
			grandList = append(grandList, finalSeq)
		}

		return 1
	} else {
		for _, codon := range codonTable[strings.ToLower(aa[0])] {
			newSeq := strings.Join(append([]string{prevSeq}, codon), " ")
			recurse(newSeq, aa[1:])
		}
	}
	return 1
}

func main() {
	recurse("", peptide)
	for i, seq := range grandList {
		i = i + 1
		fmt.Print("Seq ")
		fmt.Printf("%0.2d", i)
		fmt.Println(":", seq)
	}
}
