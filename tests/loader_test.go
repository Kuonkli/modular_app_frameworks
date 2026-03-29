package tests

import (
	"testing"

	"modular_app_frameworks/core"
)

type mockModule struct {
	name     string
	requires []string
}

func (m *mockModule) Name() string                 { return m.name }
func (m *mockModule) Requires() []string           { return m.requires }
func (m *mockModule) Register(c *core.Container)   {}
func (m *mockModule) Init(c *core.Container) error { return nil }

func TestOrder(t *testing.T) {
	all := map[string]core.Module{
		"a": &mockModule{"A", []string{}},
		"b": &mockModule{"B", []string{"A"}},
		"c": &mockModule{"C", []string{"B"}},
	}

	order, err := core.BuildOrder(all, []string{"A", "B", "C"})
	if err != nil {
		t.Fatal(err)
	}

	if order[0].Name() != "A" {
		t.Fail()
	}
}

func TestMissingModule(t *testing.T) {
	all := map[string]core.Module{
		"a": &mockModule{"A", []string{}},
	}

	_, err := core.BuildOrder(all, []string{"A", "B"})
	if err == nil {
		t.Fail()
	}
}

func TestCycle(t *testing.T) {
	all := map[string]core.Module{
		"a": &mockModule{"A", []string{"B"}},
		"b": &mockModule{"B", []string{"A"}},
	}

	_, err := core.BuildOrder(all, []string{"A", "B"})
	if err == nil {
		t.Fail()
	}
}

func TestInjection(t *testing.T) {
	container := core.NewContainer()

	container.AddSingleton("clock", func(c *core.Container) interface{} {
		return "mock-clock"
	})

	val := container.Get("clock")
	if val != "mock-clock" {
		t.Fatalf("expected injected value, got %v", val)
	}

	// transient создаёт новый объект
	counter := 0
	container.AddTransient("count", func(c *core.Container) interface{} {
		counter++
		return counter
	})

	a := container.Get("count").(int)
	b := container.Get("count").(int)
	if a == b {
		t.Fatalf("transient did not create new instance, got same value %d", a)
	}
}
