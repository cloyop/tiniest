package utils

import (
	"math/rand"

	"github.com/a-h/templ"
	"github.com/cloyop/tiniest/internal/database"
	"github.com/cloyop/tiniest/internal/templates/partials"
	"github.com/cloyop/tiniest/internal/types"
	"golang.org/x/crypto/bcrypt"
)

const characters = "abcdefghijkmlnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const firstLetterCharacters = "abcdefghijkmlnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ShortURLGen() string {
	sl := make([]byte, 6)
	is := true
	for is {
		for i := range sl {
			if i == 0 {
				sl[i] = firstLetterCharacters[rand.Intn(len(firstLetterCharacters))]
			} else {
				sl[i] = characters[rand.Intn(len(characters))]
			}
		}
		_, is, _ = database.GetPair(string(sl))
	}
	str := string(sl)
	return str
}
func HashPwd(s string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 8)
	return string(bs), err
}
func CompPwd(h, p string) error {
	return bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
}

func MapFilters(visits []types.Visit) (dateMap, fromMap map[string]templ.Component) {
	dates := make(map[string]templ.Component)
	from := make(map[string]templ.Component)
	for _, v := range visits {
		if _, ok := dates[v.Date]; !ok {
			dates[v.Date] = partials.Opt(v.Date)
		}
		if _, ok := from[v.Visit_from]; !ok {
			from[v.Visit_from] = partials.Opt(v.Visit_from)
		}
	}
	return dates, from
}
