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
	"os"
	"strings"
	"strconv"
	"errors"
	"github.com/MikunoNaka/macli/ui"
	"github.com/MikunoNaka/macli/mal"
	"github.com/MikunoNaka/macli/util"
	a "github.com/MikunoNaka/MAL2Go/v4/anime"
	m "github.com/MikunoNaka/MAL2Go/v4/manga"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var scoreCmd = &cobra.Command{
	Use:   "score",
	Short: "Set an anime/manga's status",
	Long: "Set an anime's status\n" +
	"\n" +
    "Example Usage:\n" +
	" - \x1b[33m`macli status <anime-name>`\x1b[0m For interactive prompt (anime-name can be omitted)\n" +
	" - \x1b[33m`macli status -s \x1b[34mwatching|plan_to_watch|dropped|on_hold|completed\x1b[33m <anime-name>`\x1b[0m to specify status from command\n",
	Run: func(cmd *cobra.Command, args []string) {
		mal.Init()
		searchInput := strings.Join(args, " ")

		scoreInput, err := cmd.Flags().GetString("set-value")
		if err != nil {
			fmt.Println("Error while reading \x1b[33m--set-value\x1b[0m flag.", err.Error())
			os.Exit(1)
		}

		mangaMode, err := cmd.Flags().GetBool("manga")
		if err != nil {
			fmt.Println("Error while reading \x1b[33m--manga\x1b[0m flag.", err.Error())
			os.Exit(1)
		}

		if mangaMode {
			setMangaScore(scoreInput, searchInput)
		} else {
			setAnimeScore(scoreInput, searchInput)
		}

	},
}

func validateScore(input string, currentScore int) (int, error) {
	parsedScore := util.ParseNumeric(input, currentScore)
	if parsedScore > 10 {
        return currentScore, errors.New("\x1b[31mScore out of range (" + strconv.Itoa(parsedScore) + " > 10)\x1b[0m")
	}
	if parsedScore < 0 {
        return currentScore, errors.New("\x1b[31mScore out of range (" + strconv.Itoa(parsedScore) + " < 0)\x1b[0m")
	}
	return parsedScore, nil
}

func setAnimeScore(scoreInput, searchInput string) {
	var selectedAnime a.Anime
	if entryId > 0 {
	    selectedAnime = mal.GetAnimeData(entryId, []string{"my_list_status"})
	}

	if searchInput == "" && entryId < 1 {
	    var promptText string
		if queryOnlyMode {
			promptText = "Search Anime to Get Score of: "
		} else {
			promptText = "Search Anime to Set Score of: "
		}
	   	searchInput = ui.TextInput(promptText, "Search can't be blank.")
	}

	if entryId < 1 {
		anime := ui.AnimeSearch("Select Anime: ", searchInput)
		selectedAnime = mal.GetAnimeData(anime.Id, []string{"my_list_status"})
	}

	if queryOnlyMode {
    	fmt.Println(ui.CreateScoreUpdateConfirmationMessage(selectedAnime.Title, selectedAnime.MyListStatus.Score, -1))
		os.Exit(0)
	}

    if scoreInput == "" {
		ui.AnimeScoreInput(selectedAnime)
    } else {
		score, err := validateScore(scoreInput, selectedAnime.MyListStatus.Score)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
    	resp := mal.SetAnimeScore(selectedAnime.Id, score)
    	fmt.Println(ui.CreateScoreUpdateConfirmationMessage(selectedAnime.Title, resp.Score, selectedAnime.MyListStatus.Score))
    }
}

func setMangaScore(scoreInput, searchInput string) {
	var selectedManga m.Manga
	if entryId > 0 {
	    selectedManga = mal.GetMangaData(entryId, []string{"my_list_status"})
	}

	if searchInput == "" && entryId < 1 {
	    var promptText string
		if queryOnlyMode {
			promptText = "Search Manga to Get Score of: "
		} else {
			promptText = "Search Manga to Set Score of: "
		}
	   	searchInput = ui.TextInput(promptText, "Search can't be blank.")
	}

	if entryId < 1 {
		manga := ui.MangaSearch("Select Manga: ", searchInput)
		selectedManga = mal.GetMangaData(manga.Id, []string{"my_list_status"})
	}

	if queryOnlyMode {
    	fmt.Println(ui.CreateScoreUpdateConfirmationMessage(selectedManga.Title, selectedManga.MyListStatus.Score, -1))
		os.Exit(0)
	}

    if scoreInput == "" {
		ui.MangaScoreInput(selectedManga)
    } else {
		score, err := validateScore(scoreInput, selectedManga.MyListStatus.Score)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
    	resp := mal.SetMangaScore(selectedManga.Id, score)
    	fmt.Println(ui.CreateScoreUpdateConfirmationMessage(selectedManga.Title, resp.Score, selectedManga.MyListStatus.Score))
    }
}

func init() {
	rootCmd.AddCommand(scoreCmd)
    scoreCmd.Flags().StringP("set-value", "s", "", "Score to be set")
    scoreCmd.Flags().IntVarP(&ui.PromptLength, "prompt-length", "l", 5, "Length of select prompt")
    scoreCmd.Flags().IntVarP(&mal.SearchLength, "search-length", "n", 10, "Amount of search results to load")
    scoreCmd.Flags().IntVarP(&mal.SearchOffset, "search-offset", "o", 0, "Offset for the search results")
    scoreCmd.Flags().BoolVarP(&mal.SearchNSFW, "search-nsfw", "", false, "Include NSFW-rated items in search results")
    scoreCmd.Flags().BoolVarP(&mangaMode, "manga", "m", false, "Use manga mode")
    scoreCmd.Flags().BoolVarP(&queryOnlyMode, "query", "q", false, "Query only (don't update data)")
    scoreCmd.Flags().IntVarP(&entryId, "id", "i", -1, "Manually specify the ID of anime/manga (overrides search)")
}
