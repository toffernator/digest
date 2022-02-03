package guardian

import (
	"log"
	"net/url"
	"os"
	"time"

	"github.com/toffernator/digest/utils"
)

const URL = "https://content.guardianapis.com"

func Get() (Response, error) {
	target := "https://content.guardianapis.com/search?" + apiKeyPayload().Encode()

	var resp Response
	err := utils.Get(target, &resp)
	if err != nil {
		return Response{}, err
	}

	return resp, err
}

func apiKeyPayload() url.Values {
	key, exists := os.LookupEnv("GUARDIAN_API_KEY")
	if !exists {
		log.Fatalf("Could not find an API key for the guardian")
	}

	payload := url.Values{}
	payload.Add("api-key", key)

	return payload
}

// Auto-generated using: https://transform.tools/json-to-go
type Response struct {
	Data struct {
		Status      string `json:"status"`
		UserTier    string `json:"userTier"`
		Total       int    `json:"total"`
		StartIndex  int    `json:"startIndex"`
		PageSize    int    `json:"pageSize"`
		CurrentPage int    `json:"currentPage"`
		Pages       int    `json:"pages"`
		OrderBy     string `json:"orderBy"`
		Results     []struct {
			ID                 string    `json:"id"`
			Type               string    `json:"type"`
			SectionID          string    `json:"sectionId"`
			SectionName        string    `json:"sectionName"`
			WebPublicationDate time.Time `json:"webPublicationDate"`
			WebTitle           string    `json:"webTitle"`
			WebURL             string    `json:"webUrl"`
			APIURL             string    `json:"apiUrl"`
			IsHosted           bool      `json:"isHosted"`
			PillarID           string    `json:"pillarId"`
			PillarName         string    `json:"pillarName"`
		} `json:"results"`
	} `json:"response"`
}
