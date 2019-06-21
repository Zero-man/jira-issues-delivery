package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// GetJSON prepares a request with an authorization header, and decodes the response.
func getJSON(url string, encoded string, target *results) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	var client = &http.Client{Timeout: 10 * time.Second, Transport: tr}
	log.SetFlags(log.Ltime)
	log.Printf("Making request: %s...", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", "Basic "+encoded)

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	log.Printf("Response Status Code: %s", res.Status)

	if strings.TrimSpace(res.Status) != "200" {
		return fmt.Errorf("request failed")
	}

	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

// BuildFileStr returns a markdown-formatted string of the issue data.
func buildFileStr(issues []issue, version string) (string, error) {
	if len(issues) <= 0 {
		return "", fmt.Errorf("no issues in response")
	}

	var sb strings.Builder
	today := time.Now()
	formatted := today.Format("January 02 2006")

	sb.WriteString("## " + version + " -- " + formatted + "\n")
	sb.WriteString("| Issue Key | Summary | Issue Type | Component(s) |\n")
	sb.WriteString("| --- | --- | --- | --- |\n")
	for _, v := range issues {
		sb.WriteString("|" + v.Key)
		sb.WriteString("|" + v.Fields.Summary)
		sb.WriteString("|" + v.Fields.IssueType.Name)

		componentTxt := ""
		for _, cv := range v.Fields.Components {
			componentTxt += cv.Name
		}

		sb.WriteString("|" + componentTxt + "|\n")
	}

	return sb.String(), nil
}

// WriteFile creates a markdown file in the markdown directory.
func writeFile(txt string, name string) {
	dir := "markdown"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	filename := name + ".md"
	ioutil.WriteFile(dir+"/"+filename, []byte(txt), os.ModePerm)
}
