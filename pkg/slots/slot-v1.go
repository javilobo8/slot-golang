package slots

import (
	"fmt"
	"slot-golang/pkg/utils"
)

const N_REELS = 3
const N_SYMBOLS_PER_REEL = 3

type SlotV1Result struct {
	BetAmount     float64       `json:"betAmount"`
	TotalWin      float64       `json:"totalWin"`
	NumberOfLines int           `json:"numberOfLines"`
	Reels         [N_REELS]Reel `json:"reel"`
}

type Reel [N_SYMBOLS_PER_REEL]int

type Symbol struct {
	Symbol     int
	Chance     float64
	Multiplier float64
}

var SymbolList = []Symbol{
	{Symbol: 1, Chance: 0.01, Multiplier: 0.40},   // Combination
	{Symbol: 2, Chance: 0.01, Multiplier: 0.80},   // Cerezas
	{Symbol: 3, Chance: 0.01, Multiplier: 1.60},   // Naranjas
	{Symbol: 4, Chance: 0.01, Multiplier: 2.40},   // Limones
	{Symbol: 5, Chance: 0.01, Multiplier: 3.20},   // Fresas
	{Symbol: 6, Chance: 0.01, Multiplier: 4.00},   // Campanas
	{Symbol: 7, Chance: 0.01, Multiplier: 10.00},  // 7 verde
	{Symbol: 8, Chance: 0.01, Multiplier: 20.00},  // 7 Rojo
	{Symbol: 9, Chance: 0.01, Multiplier: 100.00}, // 7 Azul
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
	for i := 0; i < N_REELS; i++ {
		for j := 0; j < N_SYMBOLS_PER_REEL; j++ {
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

func calcTotalWinLines(result *SlotV1Result) float64 {
	var totalWin float64 = 0

	for i := 0; i < result.NumberOfLines; i++ {
		var win bool = true
		currentLine := LinePositions[i]

		for j := 1; j < len(currentLine); j++ {
			if win && result.Reels[j][currentLine[j]] != result.Reels[j-1][currentLine[j-1]] {
				win = false
			}
		}

		if win {
			winSymbol := result.Reels[0][currentLine[0]] // First symbol, first match
			s := findSymbol(winSymbol)
			totalWin += result.BetAmount * s.Multiplier
		}
	}

	return totalWin
}

func printReels(reels *[N_REELS]Reel) {
	// TODO: Print dinamically
	fmt.Println(reels[0][2], reels[1][2], reels[2][2])
	fmt.Println(reels[0][1], reels[1][1], reels[2][1])
	fmt.Println(reels[0][0], reels[1][0], reels[2][0])
}

func SlotV1Bet(betAmount float64, numberOfLines int) SlotV1Result {
	result := SlotV1Result{
		BetAmount:     betAmount,
		NumberOfLines: numberOfLines,
		TotalWin:      0,
		Reels:         [N_REELS]Reel{},
	}
	fillReels(&result.Reels)
	// printReels(&result.Reels)
	result.TotalWin = calcTotalWinLines(&result)
	return result
}
