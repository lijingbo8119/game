package mahjong

import "fmt"

type Suit uint8

const (
	SuitNone Suit = iota
	SuitFlower
	SuitDragon
	SuitWind
	SuitDot
	SuitBamboo
	SuitCharacter
)

type Rank uint8

const (
	RankNone Rank = iota
	RankOne
	RankTwo
	RankThree
	RankFour
	RankFive
	RankSix
	RankSeven
	RankEight
	RankNine
)

type Tile struct {
	s Suit
	r Rank
}

func (r Tile) Suit() Suit {
	return r.s
}

func (r Tile) Rank() Rank {
	return r.r
}

func (r Tile) String() string {
	unicodeMap := map[string]string{
		fmt.Sprintf("%d%d", SuitNone, RankNone): "🀫",

		fmt.Sprintf("%d%d", SuitFlower, RankOne):   "🀢",
		fmt.Sprintf("%d%d", SuitFlower, RankTwo):   "🀣",
		fmt.Sprintf("%d%d", SuitFlower, RankThree): "🀤",
		fmt.Sprintf("%d%d", SuitFlower, RankFour):  "🀥",
		fmt.Sprintf("%d%d", SuitFlower, RankFive):  "🀦",
		fmt.Sprintf("%d%d", SuitFlower, RankSix):   "🀧",
		fmt.Sprintf("%d%d", SuitFlower, RankSeven): "🀨",
		fmt.Sprintf("%d%d", SuitFlower, RankEight): "🀩",

		fmt.Sprintf("%d%d", SuitDragon, RankTwo):   "🀄",
		fmt.Sprintf("%d%d", SuitDragon, RankThree): "🀅",
		fmt.Sprintf("%d%d", SuitDragon, RankFour):  "🀆",

		fmt.Sprintf("%d%d", SuitWind, RankTwo):   "🀀",
		fmt.Sprintf("%d%d", SuitWind, RankThree): "🀁",
		fmt.Sprintf("%d%d", SuitWind, RankFour):  "🀂",
		fmt.Sprintf("%d%d", SuitWind, RankFive):  "🀃",

		fmt.Sprintf("%d%d", SuitDot, RankOne):   "🀙",
		fmt.Sprintf("%d%d", SuitDot, RankTwo):   "🀚",
		fmt.Sprintf("%d%d", SuitDot, RankThree): "🀛",
		fmt.Sprintf("%d%d", SuitDot, RankFour):  "🀜",
		fmt.Sprintf("%d%d", SuitDot, RankFive):  "🀝",
		fmt.Sprintf("%d%d", SuitDot, RankSix):   "🀞",
		fmt.Sprintf("%d%d", SuitDot, RankSeven): "🀟",
		fmt.Sprintf("%d%d", SuitDot, RankEight): "🀠",
		fmt.Sprintf("%d%d", SuitDot, RankNine):  "🀡",

		fmt.Sprintf("%d%d", SuitBamboo, RankOne):   "🀐",
		fmt.Sprintf("%d%d", SuitBamboo, RankTwo):   "🀑",
		fmt.Sprintf("%d%d", SuitBamboo, RankThree): "🀒",
		fmt.Sprintf("%d%d", SuitBamboo, RankFour):  "🀓",
		fmt.Sprintf("%d%d", SuitBamboo, RankFive):  "🀔",
		fmt.Sprintf("%d%d", SuitBamboo, RankSix):   "🀕",
		fmt.Sprintf("%d%d", SuitBamboo, RankSeven): "🀖",
		fmt.Sprintf("%d%d", SuitBamboo, RankEight): "🀗",
		fmt.Sprintf("%d%d", SuitBamboo, RankNine):  "🀘",

		fmt.Sprintf("%d%d", SuitCharacter, RankOne):   "🀇",
		fmt.Sprintf("%d%d", SuitCharacter, RankTwo):   "🀈",
		fmt.Sprintf("%d%d", SuitCharacter, RankThree): "🀉",
		fmt.Sprintf("%d%d", SuitCharacter, RankFour):  "🀊",
		fmt.Sprintf("%d%d", SuitCharacter, RankFive):  "🀋",
		fmt.Sprintf("%d%d", SuitCharacter, RankSix):   "🀌",
		fmt.Sprintf("%d%d", SuitCharacter, RankSeven): "🀍",
		fmt.Sprintf("%d%d", SuitCharacter, RankEight): "🀎",
		fmt.Sprintf("%d%d", SuitCharacter, RankNine):  "🀏",
	}
	display, ok := unicodeMap[fmt.Sprintf("%d%d", r.s, r.r)]
	if !ok {
		panic(r)
	}
	return display
}
