package modules

import (
	"time"

	"modular_app_frameworks/core"
)

type CoreModule struct{}

func (m *CoreModule) Name() string       { return "Core" }
func (m *CoreModule) Requires() []string { return []string{} }

func (m *CoreModule) Register(c *core.Container) {
	c.AddSingleton("clock", func(c *core.Container) interface{} {
		return func() string {
			return time.Now().Format(time.RFC3339)
		}
	})

	c.AddSingleton("storage", func(c *core.Container) interface{} {
		var data []string
		return &struct {
			Add func(string)
			All func() []string
		}{
			Add: func(v string) { data = append(data, v) },
			All: func() []string { return append([]string{}, data...) },
		}
	})
}

func (m *CoreModule) Init(c *core.Container) error { return nil }
