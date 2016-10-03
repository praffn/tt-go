package trat

import (
	"encoding/json"
	"fmt"
	"time"

	"../util"
	"github.com/boltdb/bolt"
)

func Start(db *bolt.DB, name string, start string, end string) error {
	// Open an update transaction to the database
	err := db.Update(func(tx *bolt.Tx) error {
		// Get job bucket, or create one if it doesnt exist
		job, err := tx.CreateBucketIfNotExists([]byte(name))
		util.Check(err)
		// Get entries sub-bucket
		entries, err := job.CreateBucketIfNotExists([]byte("entries"))
		util.Check(err)
		// Get last entry
		_, lastValue := entries.Cursor().Last()
		// If last value is not empty, check if it has an endtime
		if len(lastValue) > 0 {
			var lastEntry Entry
			err = json.Unmarshal(lastValue, &lastEntry)
			util.Check(err)
			// Check if it has ended
			if lastEntry.End.IsZero() {
				// If the entry is unfinished, return
				fmt.Println("There is an unfinished entry")
				return nil
			}
		}

		// either there was no previous entry, or the previous
		// entry has ended, and thus we can create a new one
		id, _ := entries.NextSequence()
		entry := Entry{
			Start: time.Now(),
		}
		// Marshal it into a []byte
		entryValue, err := json.Marshal(entry)
		util.Check(err)
		// And now save it to the entries bucket

		return entries.Put(util.Itob(id), entryValue)

	})
	util.Check(err)

	return nil
}
