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

var eventHandlerTemplate = template.Must(template.New("eventHandler").Funcs(template.FuncMap{
	"constName":      constName,
	"isDiscordEvent": isDiscordEvent,
	"privateName":    privateName,
}).Parse(`package main


// Event types used to match values sent by Discord
const ({{range .}}
  {{privateName .}}EventType = "{{constName .}}"{{end}}
)
{{range .}}
// {{privateName .}}EventHandler is an event handler for {{.}} events.
type {{privateName .}}EventHandler func(*Client, *Event{{.}})
// Type returns the event type for {{.}} events.
func (eventhandler {{privateName .}}EventHandler) Type() string {
  return {{privateName .}}EventType
}
{{if isDiscordEvent .}}
// New returns a new instance of {{.}}.
func (eventhandler {{privateName .}}EventHandler) New() interface{} {
  return &Event{{.}}{}
}{{end}}
// Handle is the handler for {{.}} events.
func (eventhandler {{privateName .}}EventHandler) Handle(c *Client, i interface{}) {
  if t, ok := i.(*Event{{.}}); ok {
    eventhandler(c, t)
  }
}
{{end}}
func handlerForInterface(handler interface{}) EventHandler {
  switch v := handler.(type) { {{range .}}
  case func(*Client, *Event{{.}}):
    return {{privateName .}}EventHandler(v){{end}}
  }
  return nil
}

// EventHandler represents any EventHandler
type EventHandler interface {
	Type() string
	New() interface{}
}

var eventHandlers = map[string]EventHandler{}

func addEventHandler(eventhandler EventHandler) {
	eventHandlers[eventhandler.Type()] = eventhandler
}

// AddEventHandlers maps all event handlers
func AddEventHandlers() { {{range .}}{{if isDiscordEvent .}}
  addEventHandler({{privateName .}}EventHandler(nil)){{end}}{{end}}
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
	eventHandlerTemplate.Execute(&buf, names)

	src, err := format.Source(buf.Bytes())
	if err != nil {
		log.Println("warning: internal error: invalid Go generated:", err)
		src = buf.Bytes()
	}

	err = ioutil.WriteFile(filepath.Join(dir, strings.ToLower("eventhandlers_generated.go")), src, 0644)
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
