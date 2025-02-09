package main

import (
	"fmt"

	"github.com/beevik/ntp"
)

func main() {
	exactTime, _ := ntp.Time("pool.ntp.org")
	fmt.Println("time Now:", exactTime)
}