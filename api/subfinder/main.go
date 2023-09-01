package subfinder

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dickymuliafiqri/Binocular/common/slice"
	"github.com/dickymuliafiqri/Binocular/db"
	httpx "github.com/projectdiscovery/httpx/runner"
)

type J map[string]any

func SubfinderHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err            error
		now            = time.Now()
		domain         = r.URL.Query().Get("domain")
		isLibraryValid = false
		subdomains     = []string{}
		expiredData    = []db.DomainScheme{}
		result         = []httpx.Result{}
	)

	// Check required queries
	if domain == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(J{
			"error":   true,
			"message": fmt.Sprintf("Failed to initiate subfinder: %v", errors.New("queries not satisfied")),
		})
		return
	}

	// Check is library still valid
	isLibraryValid = db.Postgrest.IsLibraryValid(domain)

	// Subdomain enumeration
	if !isLibraryValid {
		subdomains, err = FindSubdomain(domain)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Get domains data from database
	domains, err := db.Postgrest.SelectByDomain(domain)
	if err == nil {
		for _, data := range domains {
			lastUpdated, err := time.Parse("2006-01-02", data.LastUpdated)
			if err != nil {
				fmt.Println("Failed to parse last_updated date: ", err)
				continue
			}

			daysPassed := now.Sub(lastUpdated).Hours() / 24
			if daysPassed > 3 {
				expiredData = append(expiredData, data)
				continue
			}

			subdomains = slice.RemoveAllByString(subdomains, data.Domain)
		}
	} else {
		fmt.Println("Failed to get data from database: ", err)
	}

	// Scan each subdomain
	if len(subdomains) > 0 {
		result, err = FetchDomains(subdomains)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Clasify data
	var (
		updatedData = []db.DomainScheme{}
		newData     = []db.DomainScheme{}
	)

	for _, res := range result {
		var (
			isExpired = false
			isNew     = true
			port, _   = strconv.Atoi(res.Port)
		)

		domainData := db.DomainScheme{
			Domain:       res.Input,
			Port:         port,
			StatusCode:   res.StatusCode,
			WebServer:    res.WebServer,
			Host:         res.Host,
			ResponseTime: res.ResponseTime,
			Raw:          res,
		}

		for _, data := range expiredData {
			if data.Domain == domainData.Domain {
				domainData.Id = data.Id
				domainData.LastUpdated = "now()"

				updatedData = append(updatedData, domainData)

				isExpired = true
				break
			}
		}

		if !isExpired {
			for _, data := range domains {
				if data.Domain == domainData.Domain {
					isNew = false
					break
				}
			}

			if isNew {
				domainData.LastUpdated = "now()"
				newData = append(newData, domainData)
			}
		}
	}

	// Save domain data to database
	if len(updatedData) > 0 {
		db.Postgrest.InsertMany(updatedData)
	}
	if len(newData) > 0 {
		db.Postgrest.InsertMany(newData)
	}
	if (len(updatedData)+len(newData)) > 0 || !isLibraryValid {
		db.Postgrest.UpdateLibrary(domain)
	}

	// Response API
	json.NewEncoder(w).Encode(J{
		"error":  false,
		"result": append(newData, domains...),
	})
}
