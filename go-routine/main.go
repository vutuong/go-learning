package main

import (
	"fmt"
	"time"
)

func main() {
	theMine := []string{"rock", "ore", "ore", "rock", "ore"}
	go finder1(theMine)
	go finder2(theMine)
	<-time.After(time.Second * 5)
}

func finder1(tm []string) {
	// ores := []string{}
	for _, mine := range tm {
		if mine == "ore" {
			// ores = append(ores,mine)
			fmt.Println("Finder1 found one!")
		}
	}
}

func finder2(tm []string) {
	// ores := []string{}
	for _, mine := range tm {
		if mine == "ore" {
			// ores = append(ores,mine)
			fmt.Println("Finder2 found one!")
		}
	}
}
