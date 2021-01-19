package main

import (
	"github.com/BOPR/simulator"
	"github.com/spf13/cobra"
)

func runSimulator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "simulator",
		Short: "Simulates rollup activity",
		RunE: func(cmd *cobra.Command, args []string) error {
			return simulator.Run()
		},
	}
	return cmd
}
