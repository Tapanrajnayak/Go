package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {

	RootCmd.AddCommand(Acmd)
	Acmd.AddCommand(A1cmd)
	Acmd.AddCommand(A2cmd)
	Acmd.AddCommand(A3cmd)

	RootCmd.AddCommand(Bcmd)
	Bcmd.AddCommand(B1cmd)

	RootCmd.AddCommand(Ccmd)
}

var RootCmd = &cobra.Command{
	Use: "root",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("I-Am-Root")
		return nil
	},
}

var Acmd = &cobra.Command{
	Use: "A",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintf(cmd.OutOrStdout(), "This-is-command-A")
		return nil
	},
}

var A1cmd = &cobra.Command{
	Use: "a1",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintf(cmd.OutOrStdout(), "This-is-command-a1")
		return nil
	},
}

var A2cmd = &cobra.Command{
	Use: "a2",
	Annotations: map[string]string{
		"args": "A a2",
		"output": "This-is-command-a2",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintf(cmd.OutOrStdout(), "This-is-command-a2")
		return nil
	},
}
var A3cmd = &cobra.Command{
	Use: "a3",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintf(cmd.OutOrStdout(), "This-is-command-a3")
		return nil
	},
}

var Bcmd = &cobra.Command{
	Use: "B",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintf(cmd.OutOrStdout(), "This-is-command-B")
		return nil
	},
}

var B1cmd = &cobra.Command{
	Use: "b1",
	Annotations: map[string]string{
		"args": "B b1",
		"output": "This-is-command-b1",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintf(cmd.OutOrStdout(), "This-is-command-b1")
		return nil
	},
}

var Ccmd = &cobra.Command{
	Use: "C",
	Annotations: map[string]string{
		"args": "C",
		"output": "This-is-command-C",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintf(cmd.OutOrStdout(), "This-is-command-C")
		return nil
	},
}
