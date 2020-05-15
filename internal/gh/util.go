package gh

import (
	"fmt"
	"github.com/Foxcapades/Go-ChainRequest"
	"github.com/sirupsen/logrus"
	"net/http"
)

func toUrl(slug string) string {
	return fmt.Sprintf(urlFormat, ghUrl, slug, ghLatest)
}

func validateResponse(res creq.Response, expected int, log *logrus.Entry) {
	code, err := res.GetResponseCode()
	if err != nil {
		log.Fatal("Request failed ", err)
	}
	if int(code) != expected {
		log.Fatalf("Unexpected HTTP status code.  Expected %d, got %d",
			http.StatusFound, res.MustGetResponseCode())
	}
}