package main

import (
	"fmt"

	"modular_app_frameworks/core"
	"modular_app_frameworks/modules"
)

func main() {
	container := core.NewContainer()

	all := map[string]core.Module{
		"core":       &modules.CoreModule{},
		"logging":    &modules.LoggingModule{},
		"validation": &modules.ValidationModule{},
		"export":     &modules.ExportModule{},
		"report":     &modules.ReportModule{},
	}

	names, err := core.LoadConfig("./config/modules.json")
	if err != nil {
		panic(err)
	}

	ordered, err := core.BuildOrder(all, names)
	if err != nil {
		panic(err)
	}

	for _, m := range ordered {
		m.Register(container)
	}

	for _, m := range ordered {
		m.Init(container)
	}

	fmt.Println("Запуск действий")

	actions := container.GetMany("action.")
	for _, a := range actions {
		act := a.(*modules.Action)
		fmt.Println("Действие:", act.Title)
		act.Execute()
	}
}
