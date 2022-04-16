// Copyright 2022 Lenny McLennington
// SPDX-License-Identifier: GPL-3.0-only

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "fosspotat <subcommand> [args...]",
	Short: "Libre alternative to the proprietary potatencoder software",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
