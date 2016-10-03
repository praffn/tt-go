package trat

import (
	"encoding/json"
	"fmt"
	"time"

	"../util"
	"github.com/boltdb/bolt"
)

func Stop(db *bolt.DB, name string, end string, note string) error {
	// Open an update transaction to the database
	err := db.Update(func(tx *bolt.Tx) error {
		// Get the job bucket
		job := tx.Bucket([]byte(name))
		// Check if job bucket exists
		if job == nil {
			fmt.Println("no job with name", name)
			return nil
		}
		// Get entries sub-bucket
		entries := job.Bucket([]byte("entries"))
		// Check if entries exists
		if entries == nil {
			fmt.Println("no entries in job", name)
			return nil
		}
		// Get last entry
		key, lastValue := entries.Cursor().Last()
		var lastEntry Entry
		err := json.Unmarshal(lastValue, &lastEntry)
		util.Check(err)
		// Check if last entry has end time
		if !lastEntry.End.IsZero() {
			fmt.Println("Latest entry has already stopped")
			return nil
		}
		// Set end time to now
		lastEntry.End = time.Now()
		// marshal it back into a []byte
		entryValue, err := json.Marshal(lastEntry)
		util.Check(err)
		// And save it back into the entries bucket
		return entries.Put(key, entryValue)
	})
	util.Check(err)
	return nil
}
