package credentials

import (
	"errors"
	"os"
	"strings"
)

const EnvAPIKey = "WALLBIT_API_KEY"

type Source string

const (
	SourceFlag Source = "flag"
	SourceEnv  Source = "env"
	SourceFile Source = "file"
)

var (
	ErrNotConfigured = errors.New("api key not configured: set " + EnvAPIKey + " or run wallbit auth login")
	errFileNotFound  = errors.New("credentials file not found or empty")
)

func trimKey(s string) string {
	return strings.TrimSpace(s)
}

func Load(apiKeyFromFlag string) (apiKey string, source Source, err error) {
	if k := trimKey(apiKeyFromFlag); k != "" {
		return k, SourceFlag, nil
	}

	if k := trimKey(os.Getenv(EnvAPIKey)); k != "" {
		return k, SourceEnv, nil
	}

	k, err := loadFromFile()
	if err != nil {
		if errors.Is(err, errFileNotFound) {
			return "", "", ErrNotConfigured
		}
		return "", "", err
	}
	return k, SourceFile, nil
}

func Save(apiKey string) error {
	k := trimKey(apiKey)
	if k == "" {
		return errors.New("api key is empty")
	}
	return saveToFile(k)
}

func Delete() error {
	return removeFromFile()
}

func EnvConfigured() bool {
	return trimKey(os.Getenv(EnvAPIKey)) != ""
}

func FileStoreConfigured() bool {
	return fileStoreExists()
}
