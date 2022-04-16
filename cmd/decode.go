// Copyright 2022 Lenny McLennington
// SPDX-License-Identifier: GPL-3.0-only

package cmd

import (
	"bytes"
	"fmt"
	"os"
	"regexp"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var decodeCommand = &cobra.Command{
	Use:   "decode [options] <text>",
	Short: "Decode text from potat",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bits := 8

		if cmd.Flags().Lookup("compat").Value.String() == "true" {
			bits = 7
		}

		potatre, _ := regexp.Compile(":po[tf]at:")
		matches := potatre.FindAllString(args[0], -1)

		buff := bytes.Buffer{}

		for i, match := range matches {
			if i%bits == 0 {
				buff.WriteByte(0)
			}

			if match[3] == 't' {
				buff.Bytes()[i/bits] |= 1 << (bits - (i % bits) - 1)
			}
		}

		out := buff.String()

		fmt.Println(out)

		if cmd.Flags().Lookup("copy").Value.String() == "true" {
			err := clipboard.WriteAll(out)
			if err != nil {
				fmt.Println("Error copying output to clipboard: ", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	decodeCommand.Flags().BoolP("compat", "f", false, "use proprietary potatencoder format (does not support utf-8)")
	decodeCommand.Flags().BoolP("copy", "c", false, "copy command output to the clipboard")
	rootCmd.AddCommand(decodeCommand)
}
