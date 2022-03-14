package landlord

import (
	"game/card"
)

type patternFactory func(card.Cards) card.CardPattern

var patternFactories = []patternFactory{
	FactoryCardPatternRocket,
	FactoryCardPatternBomb,

	FactoryCardPatternQuadrupletWithCards,
	FactoryCardPatternQuadrupletWithPairs,

	FactoryCardPatternSequenceOfTripletsWithPairs,
	FactoryCardPatternSequenceOfTripletsWithCards,
	FactoryCardPatternSequenceOfTriplets,
	FactoryCardPatternSequenceOfPairs,
	FactoryCardPatternSequence,

	FactoryCardPatternTripletWithCard,
	FactoryCardPatternTripletWithPair,
	FactoryCardPatternTriplet,

	FactoryCardPatternPair,
	FactoryCardPatternSingle,
}

func PatternFactory(cards card.Cards) card.CardPattern {
	for _, f := range patternFactories {
		p := f(cards)
		if p.Valid() {
			return p
		}
	}
	return nil
}
