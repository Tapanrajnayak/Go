package main

import (

	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_ExecuteAnyCommand(t *testing.T) {
	
	actual := new(bytes.Buffer)
	i := RootCmd
	_ = i
	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs([]string{"A", "a1"})
	RootCmd.Execute()

	expected := "This-is-command-a1"

	assert.Equal(t, actual.String(), expected, "actual is not expected")
}

