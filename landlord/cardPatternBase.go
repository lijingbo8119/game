package landlord

import "game/poker"

type cardPatternBase struct {
	poker.CardPattern
	cards poker.Cards
}

func (r cardPatternBase) Cards() poker.Cards {
	return r.cards
}
