package repository

type Repo[T any] interface {
	Find() T
	List() []T
}

type RepoTest struct{}

func (RepoTest) Find() int {
	return 1
}

func (RepoTest) List() []int {
	return []int{1}
}
