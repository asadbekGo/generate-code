package helper

import "os"

func ReadFile(filename string) ([]byte, error) {
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return body, nil
}
