package smtp

import (
	"errors"
	"regexp"
	"strings"
)

type Email string

func extractNameAndEmailFromString(raw string) (string, string, error) {
	haveName, err := regexp.MatchString(`<.+>.+`, raw)
	if err != nil {
		return "", "", err
	}

	if haveName {
		parts := strings.Split(raw, ">")
		if len(parts) < 2 {
			return "", "", errors.New("invalid string form from address")
		}
		name := parts[0]
		email := parts[1]

		name = strings.Replace(name, "<", "", -1)
		name = strings.TrimSpace(name)

		email = strings.TrimSpace(email)

		return name, email, nil
	}

	email := strings.TrimSpace(raw)

	return "", email, nil
}