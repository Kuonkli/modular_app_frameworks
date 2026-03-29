package modules

type Action struct {
	Title   string
	Execute func()
}
