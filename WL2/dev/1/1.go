package main

import (
	"beevik/ntp"
	"fmt"
	"os"
)

func main() {

	if time, err := ntp.Time("0.beevik-ntp.pool.ntp.org"); err != nil {
		fmt.Fprintf(os.Stderr, "TimeERREOR: %v", err)
		os.Exit(1)
	} else {
		fmt.Println(time)
	}
}
