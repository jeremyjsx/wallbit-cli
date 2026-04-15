package credentials

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	appConfigDirName = "wallbit"
	credentialsFile  = "credentials.json"
)

type filePayload struct {
	APIKey string `json:"api_key"`
}

func configDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("user config dir: %w", err)
	}
	return filepath.Join(base, appConfigDirName), nil
}

func storedCredentialPath() (string, error) {
	dir, err := configDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, credentialsFile), nil
}

func loadFromFile() (string, error) {
	path, err := storedCredentialPath()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", errFileNotFound
		}
		return "", fmt.Errorf("read credentials file: %w", err)
	}

	var p filePayload
	if err := json.Unmarshal(data, &p); err != nil {
		return "", fmt.Errorf("parse credentials file: %w", err)
	}
	key := trimKey(p.APIKey)
	if key == "" {
		return "", errFileNotFound
	}
	return key, nil
}

func saveToFile(apiKey string) error {
	path, err := storedCredentialPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}
	payload, err := json.Marshal(filePayload{APIKey: apiKey})
	if err != nil {
		return fmt.Errorf("encode credentials: %w", err)
	}
	if err := atomicWriteFile(path, payload, 0o600); err != nil {
		return fmt.Errorf("write credentisl file: %w", err)
	}
	return nil
}

func removeFromFile() error {
	path, err := storedCredentialPath()
	if err != nil {
		return err
	}
	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("remove credentials file: %w", err)
	}
	return nil
}

func fileStoreExists() bool {
	path, err := storedCredentialPath()
	if err != nil {
		return false
	}
	st, err := os.Stat(path)
	return err == nil && st.Size() > 0
}

func atomicWriteFile(path string, data []byte, perm os.FileMode) error {
	dir := filepath.Dir(path)
	f, err := os.CreateTemp(dir, ".wallbit-credentials-*")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	tmpName := f.Name()
	cleanup := true
	defer func() {
		if cleanup {
			_ = os.Remove(tmpName)
		}
	}()

	if _, err := f.Write(data); err != nil {
		_ = f.Close()
		return fmt.Errorf("write temp: %w", err)
	}
	if err := f.Sync(); err != nil {
		_ = f.Close()
		return fmt.Errorf("sync temp: %w", err)
	}
	if err := f.Chmod(perm); err != nil {
		_ = f.Close()
		return fmt.Errorf("chmod temp: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("close temp: %w", err)
	}
	if err := os.Rename(tmpName, path); err != nil {
		return fmt.Errorf("rename temp to credentials file: %w", err)
	}
	cleanup = false
	return nil
}
