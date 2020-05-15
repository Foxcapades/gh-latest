package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/Foxcapades/gh-latest/internal/flags"
	"github.com/Foxcapades/gh-latest/internal/gh"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
)

const version = "snapshot"

const (
	ghUrl    = "https://github.com/"
)

var pattern = regexp.MustCompile(`"[/\w\-.%]+/download/.+?"`)

func main() {
	logrus.SetFormatter(new(prefixed.TextFormatter))
	log := logrus.WithFields(logrus.Fields{
		"app":     "gh-latest",
		"version": version,
	})
	opts := flags.ParseArgs(version, log)

	out := new(result)
	url := ""
	out.Version, url = gh.GetReleasePath(opts.Slug, log)

	if opts.TagOnly {
		fmt.Println(out.Version)
		return
	}

	body := gh.GetReleasePage(url, log)

	for _, match := range pattern.FindAllStringIndex(body, -1) {
		subPath := body[match[0]+2 : match[1]-1]
		out.Files = append(out.Files, ghUrl+subPath)
	}

	if opts.UrlOnly {
		for _, url := range out.Files {
			fmt.Println(url)
		}
		return
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(out)
}

type result struct {
	Version string   `json:"version"`
	Files   []string `json:"files"`
}
