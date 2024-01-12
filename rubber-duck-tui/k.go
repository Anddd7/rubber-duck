package main

import (
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
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
		log.Trace(command)
	}

	command := strings.Join(commands, " ")
	log.Info(command)

	params := struct {
		FileType string
		FileName string
	}{
		FileType: "yaml",
		FileName: "test",
	}
	tmpl, err := template.New("k").Parse(command)
	if err != nil {
		log.Panic(err)
	}
	buf := new(strings.Builder)
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Panic(err)
	}

	command = buf.String()
	log.Info(command)
}
