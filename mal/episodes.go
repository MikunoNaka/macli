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

package mal

import (
  "fmt"
  "os"
  "strconv"
)

func SetEpisodes(animeId int, ep string) {
  epValue, err := strconv.Atoi(ep)
  if err != nil {
    fmt.Println("Error while parsing episode input", err)
    os.Exit(1)
  }

  sign := ep[0:1]
  if sign == "+" || sign == "-" {
    fmt.Printf("Cannot increment/decrement watched episodes by %d\n. Currently that doesn't wokr", epValue)
    os.Exit(2)
  }

  userAnimeClient.SetWatchedEpisodes(animeId, epValue)
}


func SetChapters(mangaId int, ch string) {
  chValue, err := strconv.Atoi(ch)
  if err != nil {
    fmt.Println("Error while parsing chapter input", err)
    os.Exit(1)
  }

  sign := ch[0:1]
  if sign == "+" || sign == "-" {
    fmt.Printf("Cannot increment/decrement read chapters by %d\n. Currently that doesn't wokr", chValue)
    os.Exit(2)
  }

  userMangaClient.SetChaptersRead(mangaId, chValue)
}
