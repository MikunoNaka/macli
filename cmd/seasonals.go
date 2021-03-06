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
	// "os"
	"fmt"
	// "strings"
	"github.com/MikunoNaka/macli/ui"
	"github.com/MikunoNaka/macli/util"
	"github.com/MikunoNaka/macli/mal"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var seasonalsCmd = &cobra.Command{
	Use:   "seasonals",
	Short: "Get seasonal animes",
	Long: "" +
    "" +
	"",

	Run: func(cmd *cobra.Command, args []string) {
		mal.Init()
		season := util.GetCurrentSeason()

		sort, _ := cmd.Flags().GetString("sort")
		seasonInput, _ := cmd.Flags().GetString("season")
		yearInput, _ := cmd.Flags().GetInt("year")

		if seasonInput != "" {season.Name = seasonInput}
		if yearInput > 0 {season.Year = yearInput}

		res := mal.GetSeasonalAnime(season, sort)
		for _, i := range res {
			fmt.Println(i.Title)
		}
	},
}


func init() {
	rootCmd.AddCommand(seasonalsCmd)
    seasonalsCmd.Flags().IntVarP(&ui.PromptLength, "prompt-length", "l", 5, "Length of select prompt")
    seasonalsCmd.Flags().BoolVarP(&queryOnlyMode, "query", "q", false, "Query only (don't update data)")

    seasonalsCmd.Flags().IntVarP(&mal.SearchLength, "results-length", "n", 10, "Amount of results to load")
    seasonalsCmd.Flags().BoolVarP(&mal.SearchNSFW, "include-nsfw", "", false, "Include NSFW-rated items in results")
    seasonalsCmd.Flags().IntVarP(&mal.SearchOffset, "results-offset", "o", 0, "Offset for the results")

    seasonalsCmd.Flags().StringP("sort", "", "anime_num_list_users", "sort")
    seasonalsCmd.Flags().StringP("season", "", "", "")
    seasonalsCmd.Flags().IntP("year", "", 0, "")
}
