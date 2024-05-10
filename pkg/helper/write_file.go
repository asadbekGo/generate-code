package helper

import "os"

func WriteFile(filename string, data string) error {
	// Write the string data to the file
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return err
	}

	return nil
}
