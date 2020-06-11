package main

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

var eventTypeHandlerTemplate = template.Must(template.New("eventTypeHandler").Funcs(template.FuncMap{
	"constName":      constName,
	"isDiscordEvent": isDiscordEvent,
	"privateName":    privateName,
}).Parse(`package main


// Event types used to match values sent by Discord
const ({{range .}}
  {{privateName .}}EventType = "{{constName .}}"{{end}}
)
{{range .}}
// {{privateName .}}EventTypeHandler is an event handler for {{.}} events.
type {{privateName .}}EventTypeHandler func(*Client, *Event{{.}})
// Type returns the event type for {{.}} events.
func (eventTypeHandler {{privateName .}}EventTypeHandler) Type() string {
  return {{privateName .}}EventType
}
{{if isDiscordEvent .}}
// New returns a new instance of {{.}}.
func (eventTypeHandler {{privateName .}}EventTypeHandler) New() interface{} {
  return &Event{{.}}{}
}{{end}}
// Handle is the handler for {{.}} events.
func (eventTypeHandler {{privateName .}}EventTypeHandler) Handle(c *Client, i interface{}) error {
  if t, ok := i.(*Event{{.}}); ok {
    eventhandler(c, t)
  }
}
{{end}}
func handlerForInterface(handler interface{}) EventTypeHandler {
  switch v := handler.(type) { {{range .}}
  case func(*Client, *Event{{.}}):
    return {{privateName .}}EventTypeHandler(v){{end}}
  }
  return nil
}

// EventTypeHandler represents any EventTypeHandler
type EventTypeHandler interface {
	Type() string
	New() interface{}
	Handle(c *Client, i interface{}) error
}

var eventTypeHandlers = map[string]EventTypeHandler{}

func addEventTypeHandler(eventTypeHandler EventTypeHandler) {
	eventTypeHandlers[eventTypeHandler.Type()] = eventTypeHandler
}

// AddEventTypeHandlers maps all event handlers
func AddEventTypeHandlers() { {{range .}}{{if isDiscordEvent .}}
  addEventTypeHandler({{privateName .}}EventTypeHandler(nil)){{end}}{{end}}
}
`))

func main() {
	var buf bytes.Buffer
	dir := filepath.Dir(".")

	fs := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fs, "events.go", nil, 0)
	if err != nil {
		log.Fatalf("warning: internal error: could not parse events.go: %s", err)
		return
	}

	names := []string{}
	for object := range parsedFile.Scope.Objects {
		names = append(names, object[5:])
	}
	sort.Strings(names)
	eventTypeHandlerTemplate.Execute(&buf, names)

	src, err := format.Source(buf.Bytes())
	if err != nil {
		log.Println("warning: internal error: invalid Go generated:", err)
		src = buf.Bytes()
	}

	err = ioutil.WriteFile(filepath.Join(dir, strings.ToLower("eventtypehandlers_generated.go")), src, 0644)
	if err != nil {
		log.Fatal(buf, "writing output: %s", err)
	}
}

var constRegexp = regexp.MustCompile("([a-z])([A-Z])")

func constCase(name string) string {
	return strings.ToUpper(constRegexp.ReplaceAllString(name, "${1}_${2}"))
}

func isDiscordEvent(name string) bool {
	switch {
	case name == "Connect", name == "Disconnect", name == "Event", name == "RateLimit", name == "Interface":
		return false
	default:
		return true
	}
}

func constName(name string) string {
	if !isDiscordEvent(name) {
		return "__" + constCase(name) + "__"
	}

	return constCase(name)
}

func privateName(name string) string {
	return strings.ToLower(string(name[0])) + name[1:]
}
