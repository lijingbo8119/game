package poker

import (
	"math/rand"
	"time"
)

type CardBooleanClosure = func(*Card) bool
type CardCountsClosure = func(val Value, count int) bool

type Cards []*Card

func (r Cards) Length() int {
	return len(r)
}

func (r Cards) Exists(closure CardBooleanClosure) bool {
	return r.First(closure) != nil
}

func (r Cards) First(closure ...CardBooleanClosure) *Card {
	_closure := func() CardBooleanClosure {
		if len(closure) > 0 {
			return closure[0]
		}
		return func(c *Card) bool { return true }
	}()
	for _, c := range r {
		if _closure(c) {
			return c
		}
	}
	return nil
}

func (r Cards) Last(closure ...CardBooleanClosure) *Card {
	_closure := func() CardBooleanClosure {
		if len(closure) > 0 {
			return closure[0]
		}
		return func(c *Card) bool { return true }
	}()
	for i := r.Length() - 1; i >= 0; i-- {
		if _closure(r[i]) {
			return r[i]
		}
	}
	return nil
}

func (r Cards) Count(closure ...CardBooleanClosure) int {
	_closure := func() CardBooleanClosure {
		if len(closure) > 0 {
			return closure[0]
		}
		return func(c *Card) bool { return true }
	}()
	res := 0
	for i := r.Length() - 1; i >= 0; i-- {
		if _closure(r[i]) {
			res++
		}
	}
	return res
}

func (r *Cards) Append(cards ...*Card) {
	*r = append(*r, cards...)
}

func (r *Cards) Pop() *Card {
	c := r.Last()
	if c == nil {
		return nil
	}
	r.Remove(c)
	return c
}

func (r *Cards) Remove(cards ...*Card) int {
	count := 0
	temp := Cards{}
	for _, r := range *r {
		for _, card := range cards {
			if r == card {
				count++
				goto NextLoop
			}
		}
		temp = append(temp, r)
	NextLoop:
	}
	*r = temp
	return count
}

func (r Cards) Counts(closure ...CardCountsClosure) map[Value]int {
	_closure := func() CardCountsClosure {
		if len(closure) > 0 {
			return closure[0]
		}
		return func(val Value, count int) bool { return count > 0 }
	}()
	temp := map[Value]int{}
	for _, c := range r {
		if _, ok := temp[c.Value]; !ok {
			temp[c.Value] = 0
		}
		temp[c.Value]++
	}
	result := map[Value]int{}
	for val, count := range temp {
		if _closure(val, count) {
			result[val] = count
		}
	}
	return result
}

func (r Cards) SubCards(closure CardBooleanClosure) Cards {
	res := Cards{}
	for i := r.Length() - 1; i >= 0; i-- {
		if closure(r[i]) {
			res.Append(r[i])
		}
	}
	return res
}

func (r *Cards) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < r.Length(); i++ {
		j := rand.Intn(r.Length())
		if i == j {
			continue
		}
		(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
	}
}

func (r Cards) Sort(v ...Values) {
	values := func() Values {
		if len(v) > 0 {
			return v[0]
		}
		return ValueSortRanks
	}()
	for i := 0; i < len(r); i++ {
		for j := i + 1; j < len(r); j++ {
			if values.Rank(r[i]) > values.Rank(r[j]) {
				r[i], r[j] = r[j], r[i]
			}
		}
	}
}

func (r Cards) String() string {
	res := ""
	for i, c := range r {
		res += c.String()
		if i < r.Length()-1 {
			res += " "
		}
	}
	return res
}

func NewDeckCards() Cards {
	suits := []Suit{
		SuitHeart,
		SuitDiamond,
		SuitClub,
		SuitSpade,
	}
	values := []Value{
		ValueAce,
		ValueTwo,
		ValueThree,
		ValueFour,
		ValueFive,
		ValueSix,
		ValueSeven,
		ValueEight,
		ValueNine,
		ValueTen,
		ValueJack,
		ValueQueen,
		ValueKing,
	}
	cards := Cards{}
	cards.Append(NewCard(SuitNone, ValueColoredJoker))
	cards.Append(NewCard(SuitNone, ValueJoker))
	for _, suit := range suits {
		for _, value := range values {
			cards.Append(NewCard(suit, value))
		}
	}
	return cards
}
