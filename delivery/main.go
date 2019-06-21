package main

import (
	"encoding/base64"
	"log"
	"net/url"
	"os"
)

func main() {
	var err error
	var results results
	fileName := os.Getenv("FILE_NAME")
	jqlQuery := os.Getenv("JQL_QUERY")
	domain := os.Getenv("JIRA_DOMAIN")
	if len(jqlQuery) == 0 {
		logEnd("Exiting program. ", "JQL_QUERY not specified.", 1)
	}

	queryString := url.QueryEscape(jqlQuery)

	url := domain + "/rest/api/2/search?jql=" + queryString
	encoded := base64.StdEncoding.EncodeToString([]byte(os.Getenv("JIRA_AUTHORIZATION")))
	err = getJSON(url, encoded, &results)

	if err != nil {
		logEnd("ERROR: %s. Exiting program.", err.Error(), 1)
	}

	fileStr, err := buildFileStr(results.Issues, fileName)

	if err != nil {
		logEnd("ERROR: %s. Exiting program.", err.Error(), 1)
	}

	writeFile(fileStr, fileName)
	logEnd("SUCCESS: %s. Exiting program.", "file created", 0)
}

func logEnd(msg string, errMsg string, exitCode int) {
	if exitCode == 0 {
		log.SetFlags(log.Ltime)
	} else {
		log.SetFlags(log.Ltime | log.Lshortfile)
	}

	log.Printf(msg, errMsg)
	os.Exit(exitCode)
}
