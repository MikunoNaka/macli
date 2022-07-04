/*
macli - Unofficial CLI-Based MyAnimeList Client
Copyright © 2022 Vidhu Kant Sharma <vidhukant@vidhukant.xyz>

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
  "runtime"
	"github.com/spf13/cobra"
)

const version string = "v1.7.2"

var versionCmd = &cobra.Command {
	Use:   "version",
	Short: "Shows current version",
	Long: "Shows current version of macli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("macli version", version, runtime.GOOS + "/" + runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
