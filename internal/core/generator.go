package core

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

// PasswordConfig holds options for password generation
type PasswordConfig struct {
	Length            int
	IncludeUppercase  bool
	IncludeDigits     bool
	IncludeSpecial    bool
	ExcludeHomoglyphs bool
	ExcludeChars      string
}

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars   = "0123456789"
	specialChars = "!@#$%^&*()_+-=[]{}|;:,.<>?"
	homoglyphs   = "1lIo0O" // Basic homoglyphs to exclude
)

// GeneratePassword generates a random secure password
func GeneratePassword(cfg PasswordConfig) (string, error) {
	if cfg.Length <= 0 {
		return "", errors.New("length must be greater than 0")
	}
	if cfg.Length > 256 {
		return "", errors.New("length cannot exceed 256 characters")
	}

	charSet := lowerChars
	if cfg.IncludeUppercase {
		charSet += upperChars
	}
	if cfg.IncludeDigits {
		charSet += digitChars
	}
	if cfg.IncludeSpecial {
		charSet += specialChars
	}

	if cfg.ExcludeHomoglyphs {
		charSet = removeChars(charSet, homoglyphs)
	}
	if cfg.ExcludeChars != "" {
		charSet = removeChars(charSet, cfg.ExcludeChars)
	}

	if charSet == "" {
		return "", errors.New("character set is empty")
	}

	return generateRandomString(charSet, cfg.Length)
}

func removeChars(str, charsToRemove string) string {
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(charsToRemove, r) {
			return -1
		}
		return r
	}, str)
}

func generateRandomString(charSet string, length int) (string, error) {
	b := make([]byte, length)
	maxIdx := big.NewInt(int64(len(charSet)))
	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, maxIdx)
		if err != nil {
			return "", err
		}
		b[i] = charSet[idx.Int64()]
	}
	
	// Ensure complexity requirements are met (heuristically)
	// For strict compliance we might need to retry or force inject, 
	// but for now random selection from a full set is standard.
	// Improvements could be made to ensure at least one of each selected type exists.

	return string(b), nil
}

// Helper to check if string contains char from set
func containsAny(s, chars string) bool {
	for _, r := range s {
		if strings.ContainsRune(chars, r) {
			return true
		}
	}
	return false
}

// PassphraseConfig holds options for passphrase generation
type PassphraseConfig struct {
	WordCount     int
	Separator     string
	Capitalize    bool
	IncludeNumber bool // Append a number to one of the words
	Language      string // en, fi, fr
	CustomWordListURL string
}

// GeneratePassphrase generates a mnemonic passphrase
func GeneratePassphrase(cfg PassphraseConfig) (string, error) {
	if cfg.WordCount <= 0 {
		return "", errors.New("word count must be greater than 0")
	}

	words, err := GlobalWordlistLoader.GetWords(cfg.Language, cfg.CustomWordListURL)
	if err != nil {
		return "", err
	}

	selectedWords := make([]string, cfg.WordCount)
	maxIdx := big.NewInt(int64(len(words)))

	for i := 0; i < cfg.WordCount; i++ {
		idx, err := rand.Int(rand.Reader, maxIdx)
		if err != nil {
			return "", err
		}
		word := words[idx.Int64()]
		
		if cfg.Capitalize {
			word = strings.Title(strings.ToLower(word))
		}
		selectedWords[i] = word
	}

	// Insert number/special char if requested
	// This is a simplified implementation: appending to a random word
	if cfg.IncludeNumber {
		num, _ := rand.Int(rand.Reader, big.NewInt(1000)) // 0-999
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(cfg.WordCount)))
		selectedWords[idx.Int64()] += fmt.Sprintf("%d", num)
	}

	return strings.Join(selectedWords, cfg.Separator), nil
}

