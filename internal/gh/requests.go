package gh

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"path"

	"github.com/Foxcapades/Go-ChainRequest/simple"
)

const (
	urlFormat = "%s%s%s"

	ghUrl    = "https://github.com/"
	ghLatest = "/releases/latest"
)

func GetReleasePath(slug string, log *logrus.Entry) (string, string) {
	url := toUrl(slug)
	log = log.WithField("url", url)

	res := simple.GetRequest(url).DisableRedirects().Submit()

	validateResponse(res, http.StatusFound, log)

	url, ok := res.MustLookupHeader("Location")
	if !ok {
		log.Fatal("Response missing Location header")
	}

	version := path.Base(url)

	if version == "releases" {
		log.Fatalf("Repository %s has no releases.", slug)
	}

	return version, url
}

func GetReleasePage(url string, log *logrus.Entry) string {
	log = log.WithField("url", url)
	res := simple.GetRequest(url).Submit()

	validateResponse(res, http.StatusOK, log)

	return string(res.MustGetBody())
}
