package main

import (
	"database/sql"
	"encoding/json"
	"os"

	_ "github.com/lib/pq"
)

// SeenAds is the type used to represent seen ad ids
type SeenAds map[string]bool

const filename = "./seen_ad_ids"

// ReadSeenAdIds returns a map of ad ids we've already seen.
// Not an array since Go doesn't have `contains` for arrays out the box 😬 😎
func ReadSeenAdIds() (SeenAds, error) {
	db, connectionError := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if connectionError != nil {
		return nil, connectionError
	}
	defer db.Close()

	var data []byte
	queryError := db.QueryRow("SELECT last_seen_ids FROM state").Scan(&data)
	if queryError != nil {
		return nil, queryError
	}

	var ids SeenAds
	jsonError := json.Unmarshal(data, &ids)

	if jsonError != nil {
		return nil, jsonError
	}
	return ids, nil
}

// WriteSeenAdIds writes the list of ids to Postgres.
func WriteSeenAdIds(seenAds SeenAds) error {
	json, jsonError := json.Marshal(seenAds)
	if jsonError != nil {
		return jsonError
	}

	db, connectionError := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if connectionError != nil {
		return connectionError
	}
	defer db.Close()

	_, deleteError := db.Exec("DELETE FROM state")
	if deleteError != nil {
		return deleteError
	}

	stmt, prepareError := db.Prepare("INSERT INTO state(last_seen_ids) VALUES($1)")
	if prepareError != nil {
		return prepareError
	}

	_, insertError := stmt.Exec(json)
	return insertError
}
