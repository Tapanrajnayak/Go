package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func execute(args string) string {

	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs(strings.Split(args, " "))
	RootCmd.Execute()

	return actual.String()
}


func Test_ExecuteParticularCommandDynamically(t *testing.T) {

	//run all the child command reside under root command
	for _, child := range RootCmd.Commands() {

		//check if args exist or not
		childArgs, exists := child.Annotations["args"]
		if exists {
			actual := execute(childArgs)
			expected, exists := child.Annotations["output"]
			if !exists{
				t.Errorf("Output doesn't exists in [%s] command", child.Use)
			}
			assert.Equal(t, actual, expected, "actual is not expected")
		}

		if !child.HasSubCommands() {
			continue
		}

		for _, grandChild := range child.Commands() {
			grandChildArgs, exists := grandChild.Annotations["args"]
			if exists {
				actual := execute(grandChildArgs)
				expected, exists := grandChild.Annotations["output"]
				if !exists{
					t.Errorf("Output doesn't exists in [%s] command", child.Use)
				}
				assert.Equal(t, actual, expected, "actual is not expected")
			}
		}

	}

}
