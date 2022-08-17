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
	"os"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	queryOnlyMode, mangaMode bool
	entryId int

	// auth
	saveClientId string = "yes"
	// searching
    promptLength, searchLength, searchOffset int = 5, 10, 0
	searchNsfw bool = false
	// lists
	listOffset, listLength int = 0, 15
	listIncludeNsfw bool = false
)

var rootCmd = &cobra.Command{
	Use: "macli",
	Short: "macli - Unofficial CLI-Based MyAnimeList Client.",
	Long: "macli is an unofficial MyAnimeList Client for use inside the terminal.",
}

func Execute() {
	viper.SetConfigName("macli")
	viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AddConfigPath("$HOME/.config/macli")
    viper.AddConfigPath("/etc/macli")

	// dont show error if file not found
	// macli doesnt need a config file to work properly
    if err := viper.ReadInConfig(); err != nil {
		// error if config file found but has errors
	    if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Error while reading macli config file:", err)
			fmt.Println("Exiting... Please check the macli config file.")
			os.Exit(1)
	    }
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
