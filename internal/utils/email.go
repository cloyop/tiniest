package utils

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/cloyop/tiniest/internal/types"
)

var emailsToVerify = make(map[string]UserToVerify)

type UserToVerify struct {
	User          types.User
	CodeCreatedAt time.Time
	Code          string
}

func genEmailVerificationCode() string {
	sl := make([]byte, 6)
	is := true
	for is {
		for i := range sl {
			sl[i] = characters[rand.Intn(len(characters))]
		}
		_, is = emailsToVerify[string(sl)]
	}
	return string(sl)
}
func SendingMail(to, subject string, emailContent templ.Component) {
	go func() {
		var buf bytes.Buffer
		auth := smtp.PlainAuth("", os.Getenv("MAIL"), os.Getenv("MAIL_PWD"), "smtpout.secureserver.net")
		emailContent.Render(context.Background(), &buf)
		message := fmt.Sprintf("To: %v\r\nSubject: %v\r\nFrom:%v\r\nContent-Type: text/html; charset=utf-8\r\n\r\n%s", to, subject, os.Getenv("MAIL"), buf.String())
		if err := smtp.SendMail("smtpout.secureserver.net:587", auth, os.Getenv("MAIL"), []string{to}, []byte(message)); err != nil {
			fmt.Println("Error Sending Mail to", to, "cuz ", err.Error())
		}
	}()
}
func NewUserToVerify(user *types.User) *UserToVerify {
	code := genEmailVerificationCode()
	utv := UserToVerify{User: *user, CodeCreatedAt: time.Now(), Code: code}
	emailsToVerify[user.Id] = utv
	return &utv
}
func ValidateUserVerifyCode(keyId, code string) (*types.User, error) {
	value, exist := emailsToVerify[keyId]
	if !exist {
		return &types.User{}, fmt.Errorf("error: verification form do not exists")
	}
	if time.Since(value.CodeCreatedAt) > time.Second*135 || value.Code != code {
		return &types.User{}, fmt.Errorf("error: invalid code")
	}
	delete(emailsToVerify, keyId)
	return &value.User, nil
}
func UpdateCode(keyId string) (*UserToVerify, error) {
	var nutv UserToVerify
	value, exist := emailsToVerify[keyId]
	if !exist {
		return &nutv, fmt.Errorf("error: no verification form , try again")
	}
	nutv = *NewUserToVerify(&value.User)
	return &nutv, nil
}
func CleanEmailsInterval() {
	go func() {
		for key, value := range emailsToVerify {
			if time.Since(value.CodeCreatedAt) >= time.Hour*1 {
				delete(emailsToVerify, key)
			}
		}
		time.AfterFunc(time.Hour, CleanEmailsInterval)
	}()
}
