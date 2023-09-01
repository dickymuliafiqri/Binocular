package db

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/supabase/postgrest-go"
)

var Postgrest = PostgrestClass{}

type PostgrestClass struct {
	conn *postgrest.Client
}

func (c PostgrestClass) Connect() *postgrest.Client {
	if c.conn == nil {
		c.conn = postgrest.NewClient("SUPABASE_REST_ENDPOINT", "", map[string]string{
			"apiKey": "SUPABASE_API_KEY",
		})
	}

	return c.conn
}

func (c PostgrestClass) IsLibraryValid(domain string) bool {
	var (
		now    = time.Now()
		result = LibraryScheme{}
	)
	libraryData, _, err := c.Connect().From("library").Select("", "", false).Eq("domain", domain).Single().Execute()
	if err != nil {
		return false
	}

	json.Unmarshal(libraryData, &result)

	lastUpdated, err := time.Parse("2006-01-02", result.LastScanned)
	if err != nil {
		return false
	}

	daysPassed := now.Sub(lastUpdated).Hours() / 24

	return daysPassed <= 3
}

func (c PostgrestClass) UpdateLibrary(domain string) {
	var (
		result = LibraryScheme{}
	)

	libraryData, _, err := c.Connect().From("library").Select("", "", false).Eq("domain", domain).Single().Execute()
	if err == nil {
		json.Unmarshal(libraryData, &result)
	}

	result.Domain = domain
	result.LastScanned = "now()"
	c.Connect().From("library").Insert(result, true, "", "", "").Execute()
}

func (c PostgrestClass) InsertMany(insertDomains []DomainScheme) error {
	_, _, err := c.Connect().From("domains").Insert(insertDomains, true, "", "", "").Execute()
	if err != nil {
		return err
	}

	return nil
}

func (c PostgrestClass) SelectByDomain(searchDomain string) ([]DomainScheme, error) {
	var (
		result = []DomainScheme{}
	)

	domains, _, err := c.Connect().From("domains").Select("", "", false).Like("domain", "%"+searchDomain+"%").Execute()
	if err != nil {
		return result, err
	}

	r := []DomainScheme{}
	json.Unmarshal(domains, &r)

	for _, domain := range r {
		if strings.HasSuffix(domain.Domain, searchDomain) {
			result = append(result, domain)
		}
	}

	return result, nil
}
