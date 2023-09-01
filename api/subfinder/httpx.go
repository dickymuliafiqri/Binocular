package subfinder

import (
	httpx "github.com/projectdiscovery/httpx/runner"
)

func FetchDomains(domains []string) ([]httpx.Result, error) {
	result := []httpx.Result{}
	httpxOpts := httpx.Options{
		Methods:         "GET",
		InputTargetHost: domains,
		Timeout:         10,
		Threads:         200,
		OnResult: func(r httpx.Result) {
			if r.Input != "" {
				result = append(result, r)
			}
		},
	}

	if err := httpxOpts.ValidateOptions(); err != nil {
		return result, err
	}

	hxRunner, err := httpx.New(&httpxOpts)
	if err != nil {
		return result, err
	}
	defer hxRunner.Close()

	hxRunner.RunEnumeration()

	return result, nil
}
