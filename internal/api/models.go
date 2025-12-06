package api

// GenerateRequest defines the payload for generating passwords/passphrases
type GenerateRequest struct {
	Type              string `json:"type"` // "password" or "passphrase"
	Count             int    `json:"count"`
	
	// Password specific
	Length            int    `json:"length,omitempty"`
	IncludeUppercase  *bool  `json:"include_uppercase,omitempty"`
	IncludeDigits     *bool  `json:"include_digits,omitempty"`
	IncludeSpecial    *bool  `json:"include_special,omitempty"`
	ExcludeHomoglyphs bool   `json:"exclude_homoglyphs,omitempty"`
	ExcludeChars      string `json:"exclude_chars,omitempty"`

	// Passphrase specific
	WordCount         int    `json:"word_count,omitempty"`
	Separator         string `json:"separator,omitempty"`
	Capitalize        bool   `json:"capitalize,omitempty"`
	IncludeNumber     bool   `json:"include_number,omitempty"`
	Language          string `json:"language,omitempty"` // en, fi, fr
	CustomWordListURL string `json:"custom_wordlist_url,omitempty"`
}

// GenerateResponse defines the response structure
type GenerateResponse struct {
	Passwords []string `json:"passwords"`
	Warnings  []string `json:"warnings,omitempty"`
}
