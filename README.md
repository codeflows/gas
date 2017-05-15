# gas

Reads the latest ads from [Muusikoiden.net Tori](https://muusikoiden.net/tori/)
and posts them to Slack. Keeps track of already seen ads in a Postgres JSONB column.

## Usage

```bash
SLACK_WEBHOOK_URL="https://hooks.slack.com/..." go run *.go
```

## Gotchas

Only loads the first page of results, so if the script isn't run often enough it will miss new ads.
Database code is fast and loose.
