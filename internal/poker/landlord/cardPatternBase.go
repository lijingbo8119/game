package landlord

import "game/internal/poker"

type cardPatternBase struct {
	poker.CardPattern
	cards poker.Cards
}

func (r cardPatternBase) Cards() poker.Cards {
	return r.cards
}
