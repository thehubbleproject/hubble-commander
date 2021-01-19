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
			flags := cmd.Flags()
			n, err := flags.GetInt64("iterations")
			if err != nil {
				return err
			}
			return simulator.Run(n)
		},
	}
	cmd.Flags().Int64P("iterations", "i", 0, "iteration count")
	err := cmd.MarkFlagRequired("iterations")
	if err != nil {
		panic(err)
	}
	return cmd
}
