package vial

import "go.uber.org/dig"

var registeredConstructors []interface{}

func Add(constructors ...interface{}) {
	registeredConstructors = append(registeredConstructors, constructors...)
}

func Run(executable interface{}) error {
	container := dig.New()

	for _, constructor := range registeredConstructors {
		if err := container.Provide(constructor); err != nil {
			return err
		}
	}

	return container.Invoke(executable)
}
