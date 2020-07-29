/*
Copyright © 2020 Mike de Heij

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"waarnemer/di"

	"github.com/spf13/cobra"
)

// checkCmd shows known checks
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "List checks",
	Run: func(cmd *cobra.Command, args []string) {
		cr := di.InitializeCheckRepository()

		for x, check := range cr.FindAllChecks() {
			fmt.Printf("#%d %s [%s]\n", x, check.Identifier, check.Type)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
