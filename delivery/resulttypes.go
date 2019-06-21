package main

// Results from JIRA API response containing a list of issues.
type results struct {
	Issues []issue
}

// Issue (single) in Results from JIRA API response.
type issue struct {
	Expand string
	Self   string
	ID     string
	Key    string
	Fields fields
}

// Fields for an issue.
type fields struct {
	IssueType        issueType
	TimeSpent        int
	Project          project
	Description      string
	Summary          string
	Components       []component
}

// IssueType information about an issue.
type issueType struct {
	Self        string
	ID          string
	Description string
	IconURL     string
	Name        string
	Subtask     bool
}

// Project information about an issue.
type project struct {
	Self           string
	ID             string
	Key            string
	Name           string
	ProjectTypeKey string
}

// Component for a given issue.
type component struct {
	Self string
	ID   string
	Name string
}
