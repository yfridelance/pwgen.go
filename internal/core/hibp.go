package core

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const hibpAPI = "https://api.pwnedpasswords.com/range/"

var hibpClient = &http.Client{
	Timeout: 5 * time.Second,
}

// CheckHIBP checks if the password has been exposed in a data breach.
// Returns true if pwned, false otherwise.
func CheckHIBP(password string) (bool, error) {
	if password == "" {
		return false, nil
	}

	// SHA-1 hash
	h := sha1.New()
	io.WriteString(h, password)
	hash := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))

	prefix := hash[:5]
	suffix := hash[5:]

	url := hibpAPI + prefix
	resp, err := hibpClient.Get(url)
	if err != nil {
		return false, fmt.Errorf("hibp api error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("hibp api returned status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Response is a list of suffixes and counts like:
	// 0018A45C4D9F8D...:2
	// Look for our suffix
	bodyStr := string(body)
	if strings.Contains(bodyStr, suffix) {
		return true, nil
	}

	return false, nil
}
