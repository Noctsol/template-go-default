package utils

import (
	"encoding/csv"
	"os"
)

// WriteCSV writes data to a CSV file with the given headers and records.
// filename: Path to the CSV file.
// headers: A slice of column headers.
// records: A slice of records, where each record is a slice of strings.
func WriteCSV(filename string, headers []string, records [][]string) error {
	// Create or overwrite the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush() // Ensure data is written to the file

	// Write the header row
	if len(headers) > 0 {
		if err := writer.Write(headers); err != nil {
			return err
		}
	}

	// Write the data rows
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
