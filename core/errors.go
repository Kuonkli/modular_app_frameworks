package core

type ModuleLoadError struct {
	Message string
}

func (e *ModuleLoadError) Error() string {
	return e.Message
}
