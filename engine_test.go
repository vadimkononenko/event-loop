package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/vadimkononenko/event-loop/engine"
	"reflect"
	"testing"
)

func TestCat(t *testing.T) {
	commandString := "cat amo gus"
	command := engine.Parse(commandString)
	exampleAdd := engine.CatCommand("amo", "gus")

	if assert.NotNil(t, command) {
		assert.IsType(t, reflect.TypeOf(exampleAdd), reflect.TypeOf(command))
	}
}

func TestPrint(t *testing.T) {
	commandString := "print very sus"
	command := engine.Parse(commandString)
	examplePrint := engine.PrintCommand("very sus")

	if assert.NotNil(t, command) {
		assert.IsType(t, reflect.TypeOf(examplePrint), reflect.TypeOf(command))
	}
}

func TestDefault(t *testing.T) {
	commandString := ""
	command := engine.Parse(commandString)
	examplePrint := engine.PrintCommand("very sus")

	if assert.NotNil(t, command) {
		assert.IsType(t, reflect.TypeOf(examplePrint), reflect.TypeOf(command))
	}
}
