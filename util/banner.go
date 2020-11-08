package util

import "fmt"

func PrintBanner() {
	banner := `

  pwn! Thread go n pwn

`
	fmt.Printf("%v\nVersion: %v (%v) - %v - %v\n\n", banner, Version, GitCommit, BuildDate, Author)
}
