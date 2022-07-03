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
  u "github.com/MikunoNaka/MAL2Go/v3/user"
  "fmt"
  "os"
)

func GetUserInfo() u.UserInfo {
  userInfo, err := userClient.GetSelfUserInfo()
  if err != nil {
    fmt.Println("Error fetching User Information:", err)
    os.Exit(1)
  }
  return userInfo
}
