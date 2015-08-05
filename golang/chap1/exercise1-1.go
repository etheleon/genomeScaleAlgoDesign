/*
Task1: Write a program that lists all the DNA sequences that encode a given protein sequence
*/

/*
codon table
    "UCU":"S", "UCC":"s", "UCA":"S", "UCG":"S",
    "UAU":"Y", "UAC":"Y", "UAA":"STOP", "UAG":"STOP",
    "UGU":"C", "UGC":"C", "UGA":"STOP", "UGG":"W",


    "CCU":"P", "CCC":"P", "CCA":"P", "CCG":"P",
    "CAU":"H", "CAC":"H",
    "AUU":"I", "AUC":"I", "AUA":"I",
    "ACU":"T", "ACC":"T", "ACA":"T", "ACG":"T",
    "AAU":"N", "AAC":"N", "AAA":"K", "AAG":"K",
    "AGU":"S", "AGC":"S",
    "GUU":"V", "GUC":"V", "GUA":"V", "GUG":"V",
    "GCU":"A", "GCC":"A", "GCA":"A", "GCG":"A",
    "GAU":"D", "GAC":"D", "GAA":"E", "GAG":"E",
    "GGU":"G", "GGC":"G", "GGA":"G", "GGG":"G",

*/

/*
eg. protein sequence MRQFL
*/

package main

import "fmt"

func main() {
	codonTable := map[string][]string{
		"m": []string{"AUG"},
		"r": []string{"CGG", "CGA", "CGC", "CGU", "AGG", "AGA"},
		"q": []string{"CAG", "CAA"},
		"f": []string{"UUU", "UUC"},
		"l": []string{"UUA", "UUG", "CUU", "CUC", "CUA", "CUG"},
	}

	for k := range codonTable {
		fmt.Println(k)
	}
}
