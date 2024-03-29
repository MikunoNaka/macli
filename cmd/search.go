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
	"vidhukant.com/macli/mal"
	"vidhukant.com/macli/ui"
	"vidhukant.com/macli/util"
  "vidhukant.com/mg"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command {
	Use:   "search",
	Short: "Search for an anime/manga",
	Long: "Search for an anime or a manga on MyAnimeList\n" +
	"\n" +
    "Example Usage:\n" +
	"\t\x1b[33m`macli search <anime-name>`\x1b[0m searches for an anime\n" +
	"\t\x1b[33m`macli search -m <manga-name>`\x1b[0m searches for a manga\n" +
	"\t\x1b[33m`macli search`\x1b[0m interactively asks for an anime to search for (same for manga with -m/--manga flag)\n",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := util.BindSearchConfig(cmd.Flags())
		if err != nil {
			fmt.Println("Error while parsing flags.", err.Error())
			os.Exit(1)
		}
		mal.SearchLength = conf.SearchLength
		mal.SearchOffset = conf.SearchOffset
		mal.SearchNSFW = conf.SearchNSFW
		mal.AutoSel = conf.AutoSel
		ui.PromptLength = conf.PromptLength
		mal.Init()

		// read searchInput from command
		searchInput := strings.Join(args, " ")
		mangaMode, err := cmd.Flags().GetBool("manga")
		if err != nil {
			fmt.Println("Error while reading \x1b[33m--manga\x1b[0m flag.", err.Error())
			os.Exit(1)
		}

		if mangaMode {
			searchManga(searchInput)
		} else {
			searchAnime(searchInput)
		}
	},
}

func searchManga(searchInput string) {
	var selectedManga mg.Manga
	mangaId := entryId
	fields := []string{}

	if entryId < 1 {
		if searchInput == "" {
			searchInput = ui.TextInput("Search Manga: ", "Search can't be blank.")
		}
		manga := ui.MangaSearch("Select Manga:", searchInput)
		mangaId = manga.Id
		fields = []string{"my_list_status", "num_chapters"}
	}

	selectedManga = mal.GetMangaData(mangaId, fields)

	if queryOnlyMode {
		util.PrintManga(selectedManga)
		os.Exit(0)
	}

	if entryId > 1 {
		fmt.Println("Selected: \x1b[35m" + selectedManga.Title + "\x1b[0m")
	}
	ui.MangaActionMenu(selectedManga.MyListStatus.Status != "")(selectedManga)
}

func searchAnime(searchInput string) {
	var selectedAnime mg.Anime
	animeId := entryId
	fields := []string{}

	if entryId < 1 {
		if searchInput == "" {
			searchInput = ui.TextInput("Search Anime: ", "Search can't be blank.")
		}
		anime := ui.AnimeSearch("Select Anime:", searchInput)
		animeId = anime.Id
		fields = []string{"my_list_status", "num_episodes"}
	}

	selectedAnime = mal.GetAnimeData(animeId, fields)

	if queryOnlyMode {
		util.PrintAnime(selectedAnime)
		os.Exit(0)
	}

	if entryId > 1 {
		fmt.Println("Selected: \x1b[35m" + selectedAnime.Title + "\x1b[0m")
	}
	ui.AnimeActionMenu(selectedAnime.MyListStatus.Status != "")(selectedAnime)
}

func init() {
	rootCmd.AddCommand(searchCmd)
    searchCmd.Flags().BoolVarP(&mangaMode, "manga", "m", false, "Use manga mode")
    searchCmd.Flags().IntVarP(&entryId, "id", "i", -1, "Manually specify the ID of anime/manga (overrides search)")
    searchCmd.Flags().BoolVarP(&queryOnlyMode, "query", "q", false, "Query only (don't update data)")
    searchCmd.Flags().StringVarP(&mal.Secret, "authentication-token", "t", "", "MyAnimeList authentication token to use (overrides system keyring if any)")

    searchCmd.Flags().IntP("auto-select", "S", 0, "Automatically select nth value")
    searchCmd.Flags().IntP("prompt-length", "l", 5, "Length of select prompt")
    searchCmd.Flags().IntP("search-length", "n", 10, "Amount of search results to load")
    searchCmd.Flags().IntP("search-offset", "o", 0, "Offset for the search results")
    searchCmd.Flags().BoolP("search-nsfw", "", false, "Include NSFW-rated items in search results")
}
