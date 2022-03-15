package view

type ViewBooleanClosure = func(View) bool
type ViewBooleanIndexClosure = func(View, int) bool

type Views []View

func (r Views) Length() int {
	return len(r)
}

func (r Views) Exists(closure ViewBooleanClosure) bool {
	return r.First(closure) != nil
}

func (r Views) First(closure ...ViewBooleanClosure) View {
	_closure := func() ViewBooleanClosure {
		if len(closure) > 0 {
			return closure[0]
		}
		return func(c View) bool { return true }
	}()
	for _, c := range r {
		if _closure(c) {
			return c
		}
	}
	return nil
}

func (r Views) Each(closure ViewBooleanIndexClosure) {
	for i, c := range r {
		if !closure(c, i) {
			break
		}
	}
}

func (r Views) FindActiveView() View {
	return r.First(func(v View) bool {
		return v.IsActive()
	})
}

func (r Views) SetActive(v View, active bool) {
	for _, v2 := range r {
		v2.SetActive(false)
		if v2.Name() == v.Name() {
			v2.SetActive(active)
		}
	}
}

func (r Views) Last(closure ...ViewBooleanClosure) View {
	_closure := func() ViewBooleanClosure {
		if len(closure) > 0 {
			return closure[0]
		}
		return func(c View) bool { return true }
	}()
	for i := r.Length() - 1; i >= 0; i-- {
		if _closure(r[i]) {
			return r[i]
		}
	}
	return nil
}

func (r Views) Count(closure ...ViewBooleanClosure) int {
	_closure := func() ViewBooleanClosure {
		if len(closure) > 0 {
			return closure[0]
		}
		return func(c View) bool { return true }
	}()
	res := 0
	for i := r.Length() - 1; i >= 0; i-- {
		if _closure(r[i]) {
			res++
		}
	}
	return res
}

func (r *Views) Append(Views ...View) {
	*r = append(*r, Views...)
}

func (r *Views) Remove(player View) bool {
	found := false
	temp := Views{}
	for _, p := range *r {
		if p == player {
			found = true
			continue
		}
		temp = append(temp, p)
	}
	*r = temp
	return found
}
