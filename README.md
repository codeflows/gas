# gas

Reads the latest ads from [Muusikoiden.net Tori](https://muusikoiden.net/tori/)
and posts them to Slack. Keeps track of already seen ads in a Postgres JSONB column.

## Usage

```bash
SLACK_WEBHOOK_URL="https://hooks.slack.com/..." DATABASE_URL="postgres://..." go run *.go
```
