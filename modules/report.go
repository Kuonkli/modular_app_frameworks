package modules

import (
	"fmt"

	"modular_app_frameworks/core"
)

type ReportModule struct{}

func (m *ReportModule) Name() string       { return "Report" }
func (m *ReportModule) Requires() []string { return []string{"Core", "Export"} }

func (m *ReportModule) Register(c *core.Container) {
	c.AddSingleton("action.report", func(c *core.Container) interface{} {
		storage := c.Get("storage").(*struct {
			Add func(string)
			All func() []string
		})

		clock := c.Get("clock").(func() string)

		return &Action{
			Title: "Отчёт",
			Execute: func() {
				fmt.Println("Отчёт:", clock(), len(storage.All()))
			},
		}
	})
}

func (m *ReportModule) Init(c *core.Container) error { return nil }
