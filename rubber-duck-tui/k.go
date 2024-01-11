package main

import (
	"strings"
	"text/template"

	"github.com/rs/zerolog/log"
)

func k() {
	var commands []string
	var prefixes []string
	var suffixes []string
	for _, command := range CLI.K.Commands {
		if command == "o" {
			commands = append(commands, "-o {{.FileType}}")
			suffixes = append(suffixes, "> {{.FileName}}.{{.FileType}}")
		} else {
			commands = append(commands, command)
		}
	}

	commands = append(prefixes, commands...)
	commands = append(commands, suffixes...)

	for _, command := range commands {
		log.Trace().Msg(command)
	}

	command := strings.Join(commands, " ")
	log.Info().Msg(command)

	params := struct {
		FileType string
		FileName string
	}{
		FileType: "yaml",
		FileName: "test",
	}
	tmpl, err := template.New("k").Parse(command)
	if err != nil {
		panic(err)
	}
	buf := new(strings.Builder)
	err = tmpl.Execute(buf, params)
	if err != nil {
		panic(err)
	}

	command = buf.String()
	log.Info().Msg(command)
}
