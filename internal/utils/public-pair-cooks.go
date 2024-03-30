package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cloyop/tiniest/internal/database"
	"github.com/cloyop/tiniest/internal/types"
)

func NewPublicPairCookie(lastSessionStr, toAddStr string) *http.Cookie {
	newStr := lastSessionStr + toAddStr
	return &http.Cookie{
		Name: "pairs", Value: newStr, HttpOnly: true, Secure: true, Path: "/", SameSite: http.SameSiteLaxMode,
	}
}
func GetPairCookieString(r *http.Request) string {
	cookie, err := r.Cookie("pairs")
	if err != nil || len(cookie.Value) < 7 {
		return ""
	}
	return cookie.Value
}
func ProcessCookiesString(cookStr string) *[]string {
	strs := strings.Split(cookStr, "-")
	return &strs
}
func GetPairsFromCookie(pairKeys *[]string) ([]types.PairLinks, string) {
	var str string
	for i, pairKey := range *pairKeys {
		if pairKey != "" {
			if i == 0 {
				str += fmt.Sprintf("'%v'", pairKey)
				continue
			}
			str += fmt.Sprintf(",'%v'", pairKey)
		}
	}
	pairs, err := database.GetAllPairsByKeys(str)
	if err != nil {
		return *pairs, ""
	}
	str = ""
	for _, v := range *pairs {
		str += v.Key + "-"
	}
	return *pairs, str
}
func MigratePublicToNewUser(u types.User, w http.ResponseWriter, r *http.Request) {
	cookStr := GetPairCookieString(r)
	if cookStr == "" {
		return
	}
	strs := ProcessCookiesString(cookStr)
	for _, pairKey := range *strs {
		database.MigratePairOwnership(pairKey, &u)
	}
	http.SetCookie(w, NewPublicPairCookie("", ""))
}
