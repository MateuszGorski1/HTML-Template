package templatehtml

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/rs/zerolog/log"
	"gorski.mateusz/htmltemplate/person"
)

func checkForError(err error) {
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func GenerateTemplate(tpl string, Person person.Person) {
	path := filepath.Join("assets", "person.html")
	f, err := os.Create(path)
	checkForError(err)
	t, err := template.New("webpage").Parse(tpl)
	checkForError(err)
	err = t.Execute(f, Person)
	checkForError(err)
}
