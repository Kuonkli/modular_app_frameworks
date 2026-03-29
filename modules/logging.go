package modules

import (
	"fmt"

	"modular_app_frameworks/core"
)

type LoggingModule struct{}

func (m *LoggingModule) Name() string       { return "Logging" }
func (m *LoggingModule) Requires() []string { return []string{"Core"} }

func (m *LoggingModule) Register(c *core.Container) {
	c.AddSingleton("action.logging", func(c *core.Container) interface{} {
		return &Action{
			Title: "Логирование",
			Execute: func() {
				fmt.Println("Сообщение из логгера")
			},
		}
	})
}

func (m *LoggingModule) Init(c *core.Container) error { return nil }
