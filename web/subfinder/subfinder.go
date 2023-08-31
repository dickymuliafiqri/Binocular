package subfinder

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	httpxRunner "github.com/projectdiscovery/httpx/runner"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

type J map[string]any
type metadata struct {
	PageTitle string
	Result    []httpxRunner.Result
}

func SubfinderHandler(w http.ResponseWriter, r *http.Request) {
	var (
		domain = r.URL.Query().Get("domain")
		result = []httpxRunner.Result{}
	)

	if domain == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(J{
			"error":   true,
			"message": fmt.Sprintf("Failed to initiate subfinder: %v", errors.New("queries not satisfied")),
		})
		return
	}

	// Subdomain enumeration
	subfinderOpts := &runner.Options{
		Threads:            10,
		Timeout:            30,
		MaxEnumerationTime: 10,
	}

	subfinder, err := runner.NewRunner(subfinderOpts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(J{
			"error":   true,
			"message": fmt.Sprintf("Failed to create subfinder runner: %v", err),
		})
		return
	}

	output := &bytes.Buffer{}
	if err = subfinder.EnumerateSingleDomainWithCtx(context.Background(), domain, []io.Writer{output}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(J{
			"error":   true,
			"message": fmt.Sprintf("Failed to enumerate domain: %v", err),
		})
		return
	}

	// Scan each subdomain
	httpxOpts := httpxRunner.Options{
		Methods:         "GET",
		InputTargetHost: strings.Split(output.String(), "\n"),
		Timeout:         30,
		OnResult: func(r httpxRunner.Result) {
			if r.Input != "" {
				result = append(result, r)
			}

		},
	}

	if err = httpxOpts.ValidateOptions(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(J{
			"error":   true,
			"message": fmt.Sprintf("Failed to validate httpx options: %v", err),
		})
		return
	}

	hxRunner, err := httpxRunner.New(&httpxOpts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(J{
			"error":   true,
			"message": fmt.Sprintf("Failed to start httpx runner: %v", err),
		})
		return
	}
	defer hxRunner.Close()

	hxRunner.RunEnumeration()

	tmpl := template.Must(template.New("subfinder").Parse(SubfinderTemplate))
	tmpl.Execute(w, metadata{
		PageTitle: domain,
		Result:    result,
	})
}
