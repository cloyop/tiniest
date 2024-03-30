package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

type Visit struct {
	Visit_from    string
	Title         string
	Key           string
	Url_visited   string
	User_owner    string
	Qr            bool
	DateTimestamp int64
	Date          string
	Date_hour     string
}
type ipLocationResponse struct {
	Status       string `json:"status"`
	Country      string `json:"country"`
	ErrorMessage string `json:"message"`
	City         string `json:"city"`
	Query        string `json:"query"`
}

func getIP(r *http.Request) (string, error) {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1", nil
		}
		return ip, nil
	}

	return "", errors.New("IP not found")
}

func IpLocation(ip string) (*ipLocationResponse, error) {
	var payload ipLocationResponse

	r, _ := http.NewRequest("GET", "http://ip-api.com/json/"+ip, nil)
	response, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Response Error", err.Error())
		return &payload, err
	}
	readed, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Read Error", err.Error())
		return &payload, err
	}
	if err = json.Unmarshal([]byte(readed), &payload); err != nil {
		return &payload, err
	}
	return &payload, nil
}

func (v *Visit) Visited_At() {
	now := time.Now().UTC()
	month := now.Month().String()
	if month != "April" && month != "May" && month != "June" && month != "July" {
		month = string(month[0]) + string(month[1]) + string(month[2])
	}
	v.DateTimestamp = now.Unix()
	v.Date_hour = fmt.Sprintf("%v:%v:%v", now.Hour(), now.Minute(), now.Second())
	v.Date = fmt.Sprintf("%v %v, %v", month, now.Day(), now.Year())
}
func NewVisit(pair *PairLinks, qr string, r *http.Request) (*Visit, error) {
	visit := &Visit{}
	ip, err := getIP(r)
	if err != nil {
		fmt.Println(err, "Extracting ip")
	}
	ipInfo, err := IpLocation(ip)
	if err != nil {
		return visit, err
	}
	if ipInfo.Status != "success" {
		fmt.Println("ip errro from", ip)
		return visit, fmt.Errorf("error on response " + ipInfo.Status + " " + ipInfo.ErrorMessage)
	}
	visit.Visit_from = ipInfo.Country
	visit.Key = pair.Key
	visit.Title = pair.Title
	visit.Url_visited = pair.Long
	visit.User_owner = pair.User_id
	visit.Visited_At()
	if qr != "" {
		visit.Qr = true
	}
	return visit, nil
}
