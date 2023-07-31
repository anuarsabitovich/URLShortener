package getlink

import (
	"fmt"
	"net/http"
)

func GetFullLink(shortLink string) (string, error) {
	resp, err := http.Get(shortLink)
	if err != nil {
		fmt.Println("Error:", err)
		return "", nil
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code", resp.StatusCode)
		return "", nil
	}

	finalURL := resp.Request.URL.String()
	return finalURL, nil
}
