package slots

import (
	"fmt"
	"slot-golang/utils"
)

const N_REELS = 3

type SlotV1Result struct {
	BetAmount     float32       `json:"betAmount"`
	TotalWin      float32       `json:"totalWin"`
	NumberOfLines int           `json:"numberOfLines"`
	Reels         [N_REELS]Reel `json:"reel"`
}

type Reel [3]int

type Symbol struct {
	Symbol     int
	Chance     float32
	Multiplier float32
}

var SymbolList = []Symbol{
	{Symbol: 1, Chance: 0.10, Multiplier: 0.20},
	{Symbol: 2, Chance: 0.10, Multiplier: 0.50},
	{Symbol: 3, Chance: 0.10, Multiplier: 0.75},
	{Symbol: 4, Chance: 0.10, Multiplier: 1.00},
	{Symbol: 5, Chance: 0.10, Multiplier: 2.00},
	{Symbol: 6, Chance: 0.10, Multiplier: 5.00},
}

var symbolBucket = generateBucket()

func findSymbol(symbol int) *Symbol {
	for _, item := range SymbolList {
		if symbol == item.Symbol {
			return &item
		}
	}
	return nil
}

func generateBucket() []int {
	bucket := []int{}
	for i := 0; i < len(SymbolList); i++ {
		chanceLen := int(SymbolList[i].Chance * 100)
		for j := 0; j < chanceLen; j++ {
			bucket = append(bucket, SymbolList[i].Symbol)
		}
	}
	return bucket
}

func fillReels(reels *[N_REELS]Reel) {
	for i := 0; i < len(reels); i++ {
		for j := 0; j < 3; j++ {
			reels[i][j] = utils.RandomItemFromInt(symbolBucket)
		}
	}
}

var LinePositions = [][3]int{
	{1, 1, 1},
	{0, 0, 0},
	{2, 2, 2},
	{0, 1, 2},
	{2, 1, 0},
}

func getTotalWin(result *SlotV1Result) float32 {
	var totalWin float32 = 0

	for i := 0; i < result.NumberOfLines; i++ {
		var win bool = true
		currentLine := LinePositions[i]

		for j := 1; j < len(currentLine); j++ {
			if win {
				symbolA := result.Reels[j][currentLine[j]]
				symbolB := result.Reels[j-1][currentLine[j-1]]
				if symbolA != symbolB {
					win = false
				}
			}
		}

		if win {
			fmt.Println("Win line: ", i)
			winSymbol := result.Reels[0][currentLine[0]] // First symbol, first match
			s := findSymbol(winSymbol)
			totalWin += result.BetAmount * s.Multiplier
		}
	}

	return totalWin
}

func printReels(reels *[N_REELS]Reel) {
	fmt.Println(reels[0][2], reels[1][2], reels[2][2])
	fmt.Println(reels[0][1], reels[1][1], reels[2][1])
	fmt.Println(reels[0][0], reels[1][0], reels[2][0])
}

func SlotV1Bet(betAmount float32, numberOfLines int) SlotV1Result {
	result := SlotV1Result{
		BetAmount:     betAmount,
		NumberOfLines: numberOfLines,
		TotalWin:      0,
		Reels: [N_REELS]Reel{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}
	fillReels(&result.Reels)
	printReels(&result.Reels)
	result.TotalWin = getTotalWin(&result)
	return result
}
