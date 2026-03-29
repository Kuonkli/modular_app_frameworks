package core

type Module interface {
	Name() string
	Requires() []string
	Register(c *Container)
	Init(c *Container) error
}
