package game

import (
	"github.com/gofrs/uuid"
	"github.com/samber/lo"
)

type gameBase struct {
	id     uuid.UUID
	Frames []frame
}

func (r gameBase) Id() uuid.UUID {
	if r.id.IsNil() {
		r.id = uuid.Must(uuid.NewV4())
	}
	return r.id
}

// without framePlayerTalk
func (r gameBase) lastFrame(closure ...func(f frame) bool) frame {
	var _closure func(f frame) bool
	if len(closure) > 0 {
		_closure = closure[0]
	}
	for i := len(r.Frames) - 1; i >= 0; i-- {
		if _, ok := r.Frames[i].(framePlayerTalk); ok {
			continue
		}
		if _closure == nil || _closure(r.Frames[i]) {
			return r.Frames[i]
		}
	}
	return nil
}

func (r gameBase) framesCountBy(closure func(f frame) bool) int {
	return lo.CountBy(r.Frames, closure)
}
