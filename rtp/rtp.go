package rtp

import (
	"fmt"
	"slot-golang/slots"
	"time"
)

const SPINS = 1000000
const BET_AMOUNT = 1

func calculateWithLines(lines int) {
	fmt.Printf("-- %d lines --\n", lines)

	totalBet := float32(0)
	totalWin := float32(0)

	start := time.Now()
	for i := 0; i < SPINS; i++ {
		totalBet += BET_AMOUNT
		result := slots.SlotV1Bet(BET_AMOUNT, 5)
		totalWin += result.TotalWin
	}
	rtp := totalWin * 100 / totalBet
	elapsed := time.Since(start)

	fmt.Printf("Time: %dms\n", elapsed.Milliseconds())
	fmt.Printf("Total Bet: %.2f\n", totalBet)
	fmt.Printf("Total Win: %.2f\n", totalWin)
	fmt.Printf("RTP: %.2f\n", rtp)
	fmt.Printf("\n")
}

func CalculateRTP() {
	fmt.Println("---- RTP ----")

	calculateWithLines(1)
	calculateWithLines(3)
	calculateWithLines(5)

	fmt.Println("---- END ----")
}
