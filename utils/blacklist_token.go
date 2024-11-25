package utils

import (
	"encoding/json"
	"os"
)

const blacklistFile = "./data/blacklist.json"

type Blacklist struct {
	BlacklistedTokens []string `json:"blacklisted_tokens"`
}

func ReadBlacklist() (Blacklist, error) {
	var blacklist Blacklist

	file, err := os.ReadFile(blacklistFile)
	if err != nil {
		if os.IsNotExist(err) {

			return blacklist, nil
		}
		return blacklist, err
	}

	err = json.Unmarshal(file, &blacklist)
	if err != nil {
		return blacklist, err
	}

	return blacklist, nil
}

func AddToBlacklist(tokenString string) error {
	blacklist, err := ReadBlacklist()
	if err != nil {
		return err
	}

	blacklist.BlacklistedTokens = append(blacklist.BlacklistedTokens, tokenString)

	data, err := json.MarshalIndent(blacklist, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(blacklistFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func IsTokenBlacklisted(tokenString string) bool {
	blacklist, err := ReadBlacklist()
	if err != nil {
		return false
	}

	for _, token := range blacklist.BlacklistedTokens {
		if token == tokenString {
			return true
		}
	}
	return false
}
