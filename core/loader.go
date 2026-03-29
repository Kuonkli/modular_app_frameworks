package core

import (
	"encoding/json"
	"os"
	"strings"
)

type Config struct {
	Modules []string `json:"modules"`
}

func LoadConfig(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return cfg.Modules, nil
}

func BuildOrder(all map[string]Module, enabledNames []string) ([]Module, error) {
	enabled := make(map[string]Module)

	for _, name := range enabledNames {
		key := strings.ToLower(name)
		m, ok := all[key]
		if !ok {
			return nil, &ModuleLoadError{"Модуль не найден: " + name}
		}
		enabled[key] = m
	}

	for _, m := range enabled {
		for _, r := range m.Requires() {
			if _, ok := enabled[strings.ToLower(r)]; !ok {
				return nil, &ModuleLoadError{
					"Не хватает зависимости: " + m.Name() + " требует " + r,
				}
			}
		}
	}

	indeg := make(map[string]int)
	edges := make(map[string][]string)

	for k := range enabled {
		indeg[k] = 0
		edges[k] = []string{}
	}

	for k, m := range enabled {
		for _, r := range m.Requires() {
			rk := strings.ToLower(r)
			edges[rk] = append(edges[rk], k)
			indeg[k]++
		}
	}

	var queue []string
	for k, v := range indeg {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	var result []Module

	for len(queue) > 0 {
		k := queue[0]
		queue = queue[1:]

		result = append(result, enabled[k])

		for _, to := range edges[k] {
			indeg[to]--
			if indeg[to] == 0 {
				queue = append(queue, to)
			}
		}
	}

	if len(result) != len(enabled) {
		var stuck []string
		for k, v := range indeg {
			if v > 0 {
				stuck = append(stuck, enabled[k].Name())
			}
		}
		return nil, &ModuleLoadError{
			"Циклическая зависимость: " + strings.Join(stuck, ", "),
		}
	}

	return result, nil
}
