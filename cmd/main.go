package main

import (
	"encoding/json"
	"fmt"
	"github.com/Foxcapades/Go-Chainrequest/simple"
	"github.com/Foxcapades/gh-download-latest/internal/flags"
	"net/http"
	"os"
	"path"
	"regexp"
)

const version = "snapshot"

const (
	ghUrl    = "https://github.com/"
	ghLatest = "/releases/latest"
)

var pattern = regexp.MustCompile(`"[/\w\-.%]+/download/.+?"`)

func main() {
	opts := flags.ParseArgs(version)

	if opts.UrlOnly && opts.TagOnly {
		panic("Cannot use urls-only and tag-only modes together.")
	}

	res := simple.GetRequest(fmt.Sprintf("%s%s%s", ghUrl, opts.Slug, ghLatest)).
		DisableRedirects().
		Submit()

	if res.MustGetResponseCode() != http.StatusFound {
		panic(fmt.Sprintf("Unexpected HTTP status code.  Expected %d, got %d",
			http.StatusFound, res.MustGetResponseCode()))
	}

	out := new(result)
	if url, ok := res.MustLookupHeader("Location"); !ok {
		panic("Response missing Location header")
	} else {
		out.Version = path.Base(url)
		if opts.TagOnly {
			fmt.Println(out.Version)
			os.Exit(0)
		}

		res = simple.GetRequest(url).Submit()
	}

	if res.MustGetResponseCode() != http.StatusOK {
		panic(fmt.Sprintf("Unexpected HTTP status code.  Expected %d, got %d",
			http.StatusOK, res.MustGetResponseCode()))
	}

	body := string(res.MustGetBody())
	for _, match := range pattern.FindAllStringIndex(body, -1) {
		subPath := body[match[0]+2 : match[1]-1]
		out.Files = append(out.Files, ghUrl+subPath)
	}

	if opts.UrlOnly {
		for _, url := range out.Files {
			fmt.Println(url)
		}
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(out)
}

type result struct {
	Version string   `json:"version"`
	Files   []string `json:"files"`
}
