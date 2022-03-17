package client

import (
	"github.com/samber/lo"
)

var views []view

func goToView(name string) {
	lo.ForEach(views, func(v view, i int) {
		v.SetActive(false)
		if v.Name() == name {
			v.SetActive(true)
		}
	})
}

func activeView() view {
	v, _ := lo.Find(views, func(v view) bool {
		return v.IsActive()
	})
	return v
}

func viewsInit() {
	views = []view{
		(&viewSignin{}).init(),
		(&viewSignup{}).init(),
	}
	views[0].SetActive(true)
}
