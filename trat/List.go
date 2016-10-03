package trat

import (
	"encoding/json"
	"fmt"
	"os"

	"../util"
	"github.com/boltdb/bolt"
	"github.com/olekukonko/tablewriter"
)

func printTable(entries *bolt.Bucket) {
	// Create an array of arrays of strings to hold
	// the table data
	var data [][]string
	// Loop over each entry
	entries.ForEach(func(_, entryValue []byte) error {
		// Unmarshal entry
		var entry Entry
		err := json.Unmarshal(entryValue, &entry)
		util.Check(err)
		// Format dates
		start := util.FormatDate(entry.Start)
		end := util.FormatDate(entry.End)
		// Create array of strings (row)
		row := []string{start, end, "total time here", entry.Note}
		// append to data
		data = append(data, row)

		return nil
	})

	// Create new tablewriter
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Start", "End", "Total", "Note"})
	for _, value := range data {
		table.Append(value)
	}
	table.Render()
}

func List(db *bolt.DB) {
	err := db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, bucket *bolt.Bucket) error {
			fmt.Printf("%s entries\n", name)
			// Get entries bucket
			entries := bucket.Bucket([]byte("entries"))
			// Print out a table of entries
			printTable(entries)
			return nil
		})
	})
	util.Check(err)
}
