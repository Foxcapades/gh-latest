package flags

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"

	"github.com/Foxcapades/Argonaut/v0"
	"github.com/Foxcapades/Argonaut/v0/pkg/argo"
	"github.com/Foxcapades/gh-latest/internal/config"
)

const (
	flagUrlShort = 'u'
	flagUrlLong  = "urls-only"
	flagUrlDesc  = "Print the output as a list of urls rather than a json " +
		"object.\nCannot be used with -t/--tag-only"

	flagTagShort = 't'
	flagTagLong  = "tag-only"
	flagTagDesc  = "Print only the latest release tag rather than a json object." +
		"\nCannot be used with -u/--urls-only"

	flagVerShort = 'v'
	flagVerLong  = "version"
	flagVerDesc  = "Prints this application's version."
)

func ParseArgs(version string, log *logrus.Entry) (out config.Options) {

	_, err := cli.NewCommand().
		Flag(cli.SlFlag(flagUrlShort, flagUrlLong, flagUrlDesc).
			Bind(&out.UrlOnly, false)).
		Flag(cli.SlFlag(flagTagShort, flagTagLong, flagTagDesc).
			Bind(&out.TagOnly, false)).
		Flag(cli.SlFlag(flagVerShort, flagVerLong, flagVerDesc).
			OnHit(func(argo.Flag) { fmt.Println(version); os.Exit(0) })).
		Arg(cli.NewArg().Name("User/Repo").Require().Bind(&out.Slug)).
		Parse()

	if err != nil {
		log.Fatal(err)
	}

	out.Slug = strings.Trim(out.Slug, "/")

	if out.UrlOnly && out.TagOnly {
		log.Fatal("Cannot use urls-only and tag-only modes together.")
	}

	return
}