package types

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type PairLinks struct {
	Key       string
	Short     string
	Long      string
	User_id   string
	Date      int64
	Protected bool
	Password  string
	Deleted   bool
	Title     string
}

const urlRegex = `^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`

var urlPattern = regexp.MustCompile(urlRegex)

func MatchUrlRegex(url string) bool {
	return urlPattern.MatchString(url)
}
func (p *PairLinks) DoShort(h string) {
	p.Short = "https://" + h + "/" + p.Key
}
func (p *PairLinks) CreatedAt() {
	p.Date = time.Now().Unix()
}
func (p *PairLinks) ParsedDate() string {
	date := time.Unix(p.Date, 0)
	return strings.Split(strings.Split(date.String(), ".")[0], " ")[0]
}
func (p *PairLinks) DeleteQR() {
	os.Remove("./assets/images/qrs/" + p.Key + ".jpeg")
}

func (p *PairLinks) Validate(r *http.Request) ([]string, bool) {
	var errors []string
	var hasErrors bool
	r.ParseForm()
	p.Long = r.FormValue("url")
	p.Title = r.FormValue("title")
	if !MatchUrlRegex(p.Long) {
		hasErrors = true
		errors = append(errors, "Invalid Url")
	}
	if !strings.HasPrefix(p.Long, "https") {
		hasErrors = true
		errors = append(errors, "Cant short unsecured pages")
	}
	if strings.Contains(p.Long, r.Host) {
		hasErrors = true
		errors = append(errors, "Cant Short Host URL")
	}
	if EmailPattern.Match([]byte(p.Long)) {
		hasErrors = true
		errors = append(errors, "Cannot Short Mails")
	}
	if r.FormValue("protect") == "on" {
		p.Protected = true
		p.Password = r.FormValue("password")
		if len(p.Password) < 10 || len(p.Password) > 30 {
			hasErrors = true
			errors = append(errors, "Invalid Password length")
		}
	}
	return errors, hasErrors
}
func (p *PairLinks) GenerateQR() error {
	qrc, err := qrcode.New(p.Short + "?qr=true")
	if err != nil {
		return err
	}

	w, err := standard.New("./assets/images/qrs/" + p.Key + ".jpeg")
	if err != nil {
		return err
	}

	// save file
	if err = qrc.Save(w); err != nil {
		return err
	}

	return nil
}
func (p *PairLinks) GetTittle() (string, error) {
	var title string
	c := colly.NewCollector()
	c.SetRequestTimeout(time.Second * 10)
	c.OnHTML("head > title", func(e *colly.HTMLElement) {
		title = e.Text
	})
	if err := c.Visit(p.Long); err != nil {
		return "", err
	}
	return title, nil
}
func (p *PairLinks) LookUpHost() error {
	c := colly.NewCollector()
	c.SetRequestTimeout(time.Second * 10)
	if err := c.Visit(p.Long); err != nil {
		if strings.Contains(err.Error(), "no such host") {
			return fmt.Errorf("error: " + p.Long + " appear not exist")
		}
	}
	return nil
}
