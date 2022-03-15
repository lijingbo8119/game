package landlord

import "game/internal/poker"

type patternFactory func(poker.Cards) poker.CardPattern

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

func PatternFactory(cards poker.Cards) poker.CardPattern {
	for _, f := range patternFactories {
		p := f(cards)
		if p.Valid() {
			return p
		}
	}
	return nil
}
