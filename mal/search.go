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
  "log"
  a "github.com/MikunoNaka/MAL2Go/anime"
)

func SearchAnime(searchString string) []a.Anime {
  // TODO: load limit, offset and (maybe) fields from config
  limit, offset := 10, 0
  fields := []string{"title", "id"}

  res, err := animeClient.SearchAnime(searchString, limit, offset, fields)
  if err != nil {
    log.Println(err)
    return []a.Anime{}
  }

  return res.Animes
}
