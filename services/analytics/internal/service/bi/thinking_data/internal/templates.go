package internal

import "text/template"

const (
	tdTpl = `{
              "#account_id": "{{.UserID}}",
            "#distinct_id": "{{.UserID}}",
            "#type":"{{.EventType}}",
            "#ip": "{{.IP}}",
            "#time": "{{.Time}}",
            "#event_name":"{{.EventName}}"
            {{if .Properties}}
            ,"properties":{{.Properties}}
            {{end}}}`
)

var tdTemplate *template.Template

func init() {
	tdTemplate = template.Must(template.New("thinkingdata").Parse(tdTpl))
}
