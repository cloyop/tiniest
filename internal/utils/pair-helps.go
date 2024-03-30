package utils

import (
	"time"

	"github.com/cloyop/tiniest/internal/database"
)

func DeletePublicExceed() {
	go func() {
		pairs, _ := database.GetAllPairsByID("")
		for _, p := range *pairs {
			pDate := time.Unix(p.Date, 0)
			if time.Since(pDate) >= time.Hour*48 {
				database.DeletePair(p.Key, "")
			}
		}
	}()
	time.AfterFunc(time.Hour*4, DeletePublicExceed)
}
func UpdateMaxLinks() {
	go func() {
		users, _ := database.GetAllUsers()
		for _, u := range *users {
			if u.Id != "" {
				joinedDate := time.Unix(u.Joined_Date, 0)
				MaxLinksDeserved := (int((time.Since(joinedDate).Hours()/24)/30) * 5) + 10
				if u.MaxLinksCount != MaxLinksDeserved {
					u.MaxLinksCount = MaxLinksDeserved
					database.UpdateMaxLinks(&u)
				}
			}
		}
	}()
	time.AfterFunc(time.Hour*24, UpdateMaxLinks)
}
