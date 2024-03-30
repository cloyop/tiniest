package types

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id            string
	Name          string
	Email         string
	Password      string
	Avatar_url    string
	Protected     bool
	Verified      bool
	Joined_Date   int64
	Updates       UserUpdate
	LinksCount    int
	MaxLinksCount int
}

const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

var EmailPattern = regexp.MustCompile(emailRegex)

func (u *User) StringifyJoinedDate() string {
	today := time.Unix(u.Joined_Date, 0)
	return fmt.Sprintf("%v %v, %v", today.Month(), today.Day(), today.Year())
}
func (u *User) GenerateID() {
	u.Id = uuid.Must(uuid.NewRandom()).String()
}
func (u *User) Joined() {
	u.MaxLinksCount = 10
	u.Joined_Date = time.Now().Unix()
}
func (u *User) DefaultAvatar() {
	u.Avatar_url = "https://tiniest.app/static/images/default.webp"
}

func (u *User) Validate(r *http.Request, log bool) (*[]string, bool) {
	var errors []string
	var hasErrors bool
	r.ParseForm()
	u.Email = r.FormValue("email")
	u.Password = r.FormValue("password")
	if len(u.Email) < 6 || len(u.Email) > 255 {
		hasErrors = true
		errors = append(errors, "Invalid Mail Lenght")
	}

	if !EmailPattern.MatchString(u.Email) {
		hasErrors = true
		errors = append(errors, "Invalid Mail")
	}
	if len(u.Password) < 8 || len(u.Password) > 72 {
		hasErrors = true
		errors = append(errors, "Invalid Password Lenght")
	}
	emailcut := strings.Split(u.Email, "@")
	if strings.Contains(u.Password, emailcut[0]) || strings.Contains(u.Password, emailcut[1]) {
		hasErrors = true
		errors = append(errors, "Password is very similar to email ")
	}
	if !log {
		u.Name = r.FormValue("name")
		if len(u.Name) < 4 || len(u.Name) > 30 {
			hasErrors = true
			errors = append(errors, "Invalid Name Lenght")
		}
	}
	return &errors, hasErrors
}
