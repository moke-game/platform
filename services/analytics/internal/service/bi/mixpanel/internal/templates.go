package internal

import "text/template"

const (
	mpTrackTpl = `{ 
          "event": "{{.EventName}}",
          "properties":{
			"distinct_id": "{{.UserID}}",
    		"token":"{{.Token}}",
			"time":"{{.Time}}"
			{{if .Properties}},{{.Properties}}{{end}}
			}}`
	mpEngageTpl = `{ 
          	"$token": "{{.Token}}",
			"$distinct_id": "{{.UserID}}",
			"$ip": "{{.IP}}",
			"{{.EventName}}":{
				"$name": "{{.UserID}}"
				{{if .Properties}},{{.Properties}}{{end}}
			}}`
)

var mpTrackTemplate, mpEngageTemplate *template.Template

func init() {
	mpTrackTemplate = template.Must(template.New("mixpanel_track").Parse(mpTrackTpl))
	mpEngageTemplate = template.Must(template.New("mixpanel_engage").Parse(mpEngageTpl))
}
