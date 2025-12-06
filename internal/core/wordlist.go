package core

import (
	"bufio"
	"errors"
	"net/http"
	"strings"
	"sync"
)

var (
	// Default Wordlist URLs - using EFF wordlists as examples or placeholders
	// In a real scenario, we might want to ensure these exist or host our own.
	// For this task, we'll assume these generic paths are placeholders to be configured.
	// We'll use a map to store defaults.
	DefaultWordlists = map[string]string{
		"en": "https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt", // A popular wordlist
		"fi": "https://raw.githubusercontent.com/jocxfin/pwgen/refs/heads/main/wordlist_fi.txt",
		"fr": "https://raw.githubusercontent.com/jocxfin/pwgen/refs/heads/main/wordlist_fr.txt",
	}
)

type WordlistLoader struct {
	cache map[string][]string
	mutex sync.RWMutex
}

var GlobalWordlistLoader = &WordlistLoader{
	cache: make(map[string][]string),
}

func (wl *WordlistLoader) GetWords(lang, customURL string) ([]string, error) {
	wl.mutex.RLock()
	// Create a cache key. If customURL is provided, use it as key, otherwise lang.
	key := lang
	if customURL != "" {
		key = customURL
	}
	
	if words, ok := wl.cache[key]; ok {
		wl.mutex.RUnlock()
		return words, nil
	}
	wl.mutex.RUnlock()

	wl.mutex.Lock()
	defer wl.mutex.Unlock()

	// Double check locking
	if words, ok := wl.cache[key]; ok {
		return words, nil
	}

	url := customURL
	if url == "" {
		var ok bool
		url, ok = DefaultWordlists[lang]
		if !ok {
			return nil, errors.New("unsupported language and no custom url provided")
		}
	}

	// Security check for URL
	if !isValidURL(url) {
		return nil, errors.New("invalid wordlist url: must start with https://raw.githubusercontent.com/")
	}

	words, err := fetchWords(url)
	if err != nil {
		return nil, err
	}

	if len(words) == 0 {
		return nil, errors.New("wordlist is empty")
	}

	wl.cache[key] = words
	return words, nil
}

func isValidURL(url string) bool {
	// Simple check as per requirements. 
	// Can be disabled via config strictly speaking, but implemented as safe default here.
	// The requirement: "By default URLs are required to start with https://raw.githubusercontent.com/"
	return strings.HasPrefix(url, "https://raw.githubusercontent.com/")
}

func fetchWords(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch wordlist: " + resp.Status)
	}

	var words []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		w := strings.TrimSpace(scanner.Text())
		if w != "" {
			words = append(words, w)
		}
	}
	return words, scanner.Err()
}
