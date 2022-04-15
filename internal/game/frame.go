package game

import (
	"encoding/json"
	"fmt"
	"time"
)

type frame interface {
	fmt.Stringer
	json.Marshaler
	Name() string
	time() time.Time
	beforeUpdate(Game) error
	update(Game) error
}
