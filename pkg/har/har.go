package har

import (
	"encoding/json"
	"fmt"
	"har2sequence/pkg/config"
	"har2sequence/pkg/diagram"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
)

type HAR struct {
	Log struct {
		Entries []struct {
			Request struct {
				Method string `json:"method"`
				URL    string `json:"url"`
			} `json:"request"`
			Response struct {
				Status int `json:"status"`
			} `json:"response"`
			StartedDateTime string `json:"startedDateTime"`
		} `json:"entries"`
	} `json:"log"`
}

func LoadHAR(filePath string) (HAR, error) {
	var har HAR
	file, err := os.Open(filePath)
	if err != nil {
		return har, fmt.Errorf("failed to open HAR file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return har, fmt.Errorf("failed to read HAR file: %v", err)
	}

	err = json.Unmarshal(byteValue, &har)
	if err != nil {
		return har, fmt.Errorf("failed to parse HAR file: %v", err)
	}

	return har, nil
}

func (har HAR) removeQueryParams() HAR {
	newHar := har
	for i := range newHar.Log.Entries {
		u, err := url.Parse(newHar.Log.Entries[i].Request.URL)
		if err == nil {
			u.RawQuery = ""
			newHar.Log.Entries[i].Request.URL = u.String()
		}
	}
	return newHar
}

func (har HAR) filterEntries(config config.Config) HAR {
	newHar := har
	var filteredEntries []struct {
		Request struct {
			Method string `json:"method"`
			URL    string `json:"url"`
		} `json:"request"`
		Response struct {
			Status int `json:"status"`
		} `json:"response"`
		StartedDateTime string `json:"startedDateTime"`
	}
	for _, entry := range newHar.Log.Entries {
		u, err := url.Parse(entry.Request.URL)
		if err != nil {
			log.Printf("Failed to parse URL %s: %v", entry.Request.URL, err)
			continue
		}

		domain := u.Host

		// Check if domain is in the participants list
		isParticipant := len(config.Participants) == 0
		for _, participant := range config.Participants {
			if domain == participant {
				isParticipant = true
				break
			}
		}

		// Check if URL is in the exclude list
		isExcluded := false
		for _, excludeURL := range config.ExcludePaths {
			if strings.Contains(entry.Request.URL, excludeURL) {
				isExcluded = true
				break
			}
		}

		if isParticipant && !isExcluded {
			filteredEntries = append(filteredEntries, entry)
		}
	}
	newHar.Log.Entries = filteredEntries
	return newHar
}

func (har HAR) GenerateSequenceDiagram(config config.Config) diagram.SequenceDiagram {
	var processedHar = har.removeQueryParams().filterEntries(config)
	var browser = "Browser"

	sequenceDiagram := diagram.SequenceDiagram{
		Participants:    []string{browser},
		Messages:        []diagram.Message{},
		MessagePrefixes: config.MessagePrefixes,
	}

	domainSet := make(map[string]bool)
	for _, entry := range processedHar.Log.Entries {
		request := entry.Request
		response := entry.Response

		u, err := url.Parse(request.URL)
		if err != nil {
			log.Printf("Failed to parse URL %s: %v", request.URL, err)
			continue
		}

		domain := u.Host
		path := u.Path

		if _, exists := domainSet[domain]; !exists {
			domainSet[domain] = true
			sequenceDiagram.Participants = append(sequenceDiagram.Participants, domain)
		}

		message := diagram.Message{
			From: browser,
			To:   domain,
			Request: diagram.Request{
				Method: request.Method,
				Path:   path,
			},
			Response: diagram.Response{
				Status: response.Status,
			},
		}

		sequenceDiagram.Messages = append(sequenceDiagram.Messages, message)
	}

	return sequenceDiagram
}
