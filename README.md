# JIRA Issues Delivery

This program generates and emails a markdown formatted table of JIRA issues provided a custom JQL query using sSMTP. For example, it could be used to deliver all issues tagged in a specific build.

## Requirements

Docker and docker-compose.

## Configure Environment Variables
Before attempting to use this project, you must define the following variables in `docker-compose.yml`:
* `JIRA_DOMAIN`: The JIRA domain for your organization, e.g.: `"https://your-org.atlassian.net"`
* `JQL_QUERY`: The JQL query.
* `FILE_NAME`: A name for the table header and output file. For example, the JQL query might request all of the issues in a specific build or release. In this case, a build version would be a good candidate.
* `JIRA_AUTHORIZATION`: Basic authorization string: `"username:password"`

And the following ssmtp configuration settings, as well as the recipient email/group:
* `MAIL_ROOT`
* `MAILHUB`
* `PORT`
* `MAIL_USER`
* `MAIL_PASSWORD`
* `MAIL_HOSTNAME`
* `RECIPIENT`

You can read more about sSMTP settings here: https://wiki.debian.org/sSMTP

## Run

In the root directory of the project, run `docker-compose up`.