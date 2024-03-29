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

package ui

import (
  "vidhukant.com/mg"
  "github.com/jedib0t/go-pretty/v6/table"
  "fmt"
  "os"
)

func AnimeList(animes []mg.Anime) {
  t := table.NewWriter()
  t.SetOutputMirror(os.Stdout)

  t.AppendHeader(table.Row{"#", "Title", "ID", "Score", "Type", "Status", "Progress"})

  for index, anime := range animes {
    status := anime.ListStatus.Status
    score := anime.ListStatus.Score

    formattedStatus := GetColorCodeByStatus(status) + FormatStatus(status) + "\x1b[0m"
    formattedScore := FormatScore(score)

    // TODO: format it
    formattedMediaType := anime.MediaType

    progress := fmt.Sprintf("%d/%d", anime.ListStatus.EpWatched, anime.NumEpisodes)

    t.AppendRow([]interface{}{
      index + 1, anime.Title, anime.Id, formattedScore, formattedMediaType, formattedStatus, progress,
    })
  }

  t.Render()
}

func MangaList(mangas []mg.Manga) {
  t := table.NewWriter()
  t.SetOutputMirror(os.Stdout)

  t.AppendHeader(table.Row{"#", "Title", "ID", "Score", "Type", "Status", "Chapters", "Volumes"})

  for index, manga := range mangas {
    status := manga.ListStatus.Status
    score := manga.ListStatus.Score

    formattedStatus := GetColorCodeByStatus(status) + FormatStatus(status) + "\x1b[0m"
    formattedScore := FormatScore(score)

    // TODO: format it
    formattedMediaType := manga.MediaType

    chapterProgress := fmt.Sprintf("%d/%d", manga.ListStatus.ChaptersRead, manga.NumChapters)
    volumeProgress := fmt.Sprintf("%d/%d", manga.ListStatus.VolumesRead, manga.NumVolumes)

    t.AppendRow([]interface{}{
      index + 1, manga.Title, manga.Id, formattedScore, formattedMediaType, formattedStatus, chapterProgress, volumeProgress,
    })
  }

  t.Render()
}
