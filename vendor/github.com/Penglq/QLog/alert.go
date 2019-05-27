package QLog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type AlertApiConfig struct {
	AppId       string
	URL         string
	ContentType string
	Format      string
	IsAlert     bool
}

type alertApiData struct {
	AppId    string `json:"appId"`
	Content  string `json:"content"`
	Priority string `json:"priority"`
}

//"alert"
func (l *logger) alert(s string, priority string) {
	if l.alertURL != "" {
		client := http.Client{}
		data := &alertApiData{l.alertAppId, s, priority}
		str, err := json.Marshal(data)
		resp, err := client.Post(l.alertURL, l.alertContentType, strings.NewReader(string(str)))
		if err != nil {
			l.Error("alertApi alert", err.Error())
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				l.Error("alert service read body error", err.Error())
			} else {
				l.Info("alert service body", string(body))
			}

			if resp.StatusCode != http.StatusOK {
				l.Info("alert service http error", resp.StatusCode, "body", string(body))
			}
		}
	}
}
