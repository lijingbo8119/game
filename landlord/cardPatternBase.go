package landlord

import (
	"game/card"
)

type cardPatternBase struct {
	card.CardPattern
	cards card.Cards
}

func (r cardPatternBase) Cards() card.Cards {
	return r.cards
}
