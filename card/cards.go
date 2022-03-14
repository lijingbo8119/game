package card

import (
	"fmt"
	"math/rand"
	"time"
)

type CardBooleanClosure = func(*Card) bool
type CardCountsClosure = func(val CardValue, count int) bool

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

func (r Cards) Append(cards ...*Card) Cards {
	r = append(r, cards...)
	return r
}

func (r Cards) Counts(closure ...CardCountsClosure) map[CardValue]int {
	_closure := func() CardCountsClosure {
		if len(closure) > 0 {
			return closure[0]
		}
		return func(val CardValue, count int) bool { return count > 0 }
	}()
	temp := map[CardValue]int{}
	for _, c := range r {
		if _, ok := temp[c.Value()]; !ok {
			temp[c.Value()] = 0
		}
		temp[c.Value()]++
	}
	result := map[CardValue]int{}
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
		fmt.Println(time.Now(), i, j)
		if i == j {
			continue
		}
		(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
	}
}

func (r Cards) Sort(v ...CardValues) {
	values := func() CardValues {
		if len(v) > 0 {
			return v[0]
		}
		return CardValueSortRanks
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
