// Copyright 2022 Lenny McLennington
// SPDX-License-Identifier: GPL-3.0-only

package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var encodeCommand = &cobra.Command{
	Use:   "encode [options] <text>",
	Short: "Encode text into potat",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bits := 8

		if cmd.Flags().Lookup("compat").Value.String() == "true" {
			bits = 7
		}

		r := strings.NewReader(args[0])

		out := ""

		for {
			b, err := r.ReadByte()
			if err == io.EOF {
				break
			}

			for i := 0; i < bits; i++ {
				out += func() string {
					if (b>>(bits-i-1))&1 == 1 {
						return ":potat:"
					} else {
						return ":pofat:"
					}
				}()
			}
		}

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
	encodeCommand.Flags().BoolP("compat", "f", false, "use proprietary potatencoder format (does not support utf-8)")
	encodeCommand.Flags().BoolP("copy", "c", false, "copy command output to the clipboard")
	rootCmd.AddCommand(encodeCommand)
}
