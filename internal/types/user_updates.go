package types

import (
	"strings"
	"time"
)

type UpdateStatus struct {
	Locked      bool
	ReleaseDate string
}
type UserUpdate struct {
	Email_Last_Update          int64
	Name_Last_Update           int64
	Password_Last_Update       int64
	Email_Update_Permission    UpdateStatus
	Name_Update_Permission     UpdateStatus
	Password_Update_Permission UpdateStatus
}

func (u_u *UserUpdate) FullFill() {
	if u_u.Email_Last_Update > 0 {
		last := time.Unix(u_u.Email_Last_Update, 0)
		since := time.Since(last)
		lastStr := strings.Split(strings.Split(last.UTC().String(), ".")[0], " ")[0]
		releaseD := strings.Split(strings.Split(last.Add(time.Hour*360).UTC().String(), ".")[0], " ")[0]
		if since < time.Hour*360 && lastStr != releaseD {
			u_u.Email_Update_Permission.Locked = true
			u_u.Email_Update_Permission.ReleaseDate = releaseD
		}
	}
	if u_u.Name_Last_Update > 0 {
		last := time.Unix(u_u.Name_Last_Update, 0)
		since := time.Since(last)
		lastStr := strings.Split(strings.Split(last.UTC().String(), ".")[0], " ")[0]
		releaseD := strings.Split(strings.Split(last.Add(time.Hour*360).UTC().String(), ".")[0], " ")[0]
		if since < time.Hour*360 && lastStr != releaseD {
			u_u.Name_Update_Permission.Locked = true
			u_u.Name_Update_Permission.ReleaseDate = releaseD
		}
	}
	if u_u.Password_Last_Update > 0 {
		last := time.Unix(u_u.Password_Last_Update, 0)
		since := time.Since(last)
		lastStr := strings.Split(strings.Split(last.UTC().String(), ".")[0], " ")[0]
		releaseD := strings.Split(strings.Split(last.Add(time.Hour*360).UTC().String(), ".")[0], " ")[0]
		if since < time.Hour*360 && lastStr != releaseD {
			u_u.Password_Update_Permission.Locked = true
			u_u.Password_Update_Permission.ReleaseDate = releaseD
		}
	}
}
