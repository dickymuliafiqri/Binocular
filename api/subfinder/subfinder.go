package subfinder

import (
	"bytes"
	"context"
	"io"
	"strings"

	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

func FindSubdomain(domain string) ([]string, error) {
	var (
		subdomains    = []string{}
		subfinderOpts = &runner.Options{
			Threads:            50,
			Timeout:            30,
			MaxEnumerationTime: 10,
		}
	)

	subfinder, err := runner.NewRunner(subfinderOpts)
	if err != nil {
		return subdomains, err
	}

	output := &bytes.Buffer{}
	if err = subfinder.EnumerateSingleDomainWithCtx(context.Background(), domain, []io.Writer{output}); err != nil {
		return subdomains, err
	}

	subdomains = strings.Split(output.String(), "\n")

	return subdomains, nil
}
