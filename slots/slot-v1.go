package slots

import "test-go/utils"

type SlotV1Result struct {
	BetAmount float32 `json:"betAmount"`
	TotalWin  float32 `json:"totalWin"`
	Reel      [3]Reel `json:"reel"`
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

func fillReel(reel *Reel) {
	for i := 0; i < 3; i++ {
		reel[i] = utils.RandomItemFromInt(symbolBucket)
	}
}

func SlotV1(betAmount float32) SlotV1Result {
	result := SlotV1Result{
		BetAmount: betAmount,
		TotalWin:  0,
		Reel: [3]Reel{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}
	fillReel(&result.Reel[0])
	fillReel(&result.Reel[1])
	fillReel(&result.Reel[2])
	return result
}
