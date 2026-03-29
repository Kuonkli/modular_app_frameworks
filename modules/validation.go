package modules

import "modular_app_frameworks/core"

type ValidationModule struct{}

func (m *ValidationModule) Name() string       { return "Validation" }
func (m *ValidationModule) Requires() []string { return []string{"Core"} }

func (m *ValidationModule) Register(c *core.Container) {
	c.AddSingleton("action.validation", func(c *core.Container) interface{} {
		storage := c.Get("storage").(*struct {
			Add func(string)
			All func() []string
		})

		return &Action{
			Title: "Валидация",
			Execute: func() {
				val := "example"
				if len(val) < 3 {
					panic("короткое значение")
				}
				storage.Add(val)
			},
		}
	})
}

func (m *ValidationModule) Init(c *core.Container) error { return nil }
