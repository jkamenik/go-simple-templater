package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("%s |", i)
		},
	})

	tmpl, err := template.New("test").Funcs(sprig.FuncMap()).Parse("{{ \"hello!\" | upper | repeat 5 }}")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
