package main

import (
	"fmt"
	"log"

	"github.com/bsstahl/DeterministicAI/LPDemos/PetesPotteryParadise/go/internal/pottery"
)

func main() {
	plan, err := pottery.Solve(24, 16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pete's Pottery Paradise production plan")
	fmt.Printf("Small vases: %d\n", plan.SmallVases)
	fmt.Printf("Large vases: %d\n", plan.LargeVases)
	fmt.Printf("Clay used: %d oz.\n", plan.ClayUsedOz)
	fmt.Printf("Glaze used: %d oz.\n", plan.GlazeUsedOz)
	fmt.Printf("Maximum revenue: $%d.00\n", plan.Revenue)
}
