package modules

import (
	"os"
	"strings"

	"modular_app_frameworks/core"
)

type ExportModule struct{}

func (m *ExportModule) Name() string       { return "Export" }
func (m *ExportModule) Requires() []string { return []string{"Core", "Validation"} }

func (m *ExportModule) Register(c *core.Container) {
	c.AddSingleton("action.export", func(c *core.Container) interface{} {
		storage := c.Get("storage").(*struct {
			Add func(string)
			All func() []string
		})

		return &Action{
			Title: "Экспорт",
			Execute: func() {
				data := storage.All()
				os.WriteFile("export.txt", []byte(strings.Join(data, "\n")), 0644)
			},
		}
	})
}

func (m *ExportModule) Init(c *core.Container) error { return nil }
