package credentials

import (
	"errors"
	"path/filepath"
	"runtime"
	"testing"
)

// redirectConfigDir sends the credentials file under t.TempDir() so tests never touch the real user profile.
func redirectConfigDir(t *testing.T) {
	t.Helper()
	root := t.TempDir()
	switch runtime.GOOS {
	case "windows":
		t.Setenv("AppData", root)
	case "darwin":
		t.Setenv("HOME", root)
	default:
		t.Setenv("XDG_CONFIG_HOME", root)
	}
}

func TestLoad_flagOverridesEnv(t *testing.T) {
	t.Setenv(EnvAPIKey, "from-env")
	k, src, err := Load("  from-flag  ")
	if err != nil {
		t.Fatal(err)
	}
	if src != SourceFlag || k != "from-flag" {
		t.Fatalf("got key=%q src=%v", k, src)
	}
}

func TestLoad_fromEnv(t *testing.T) {
	redirectConfigDir(t)
	t.Setenv(EnvAPIKey, "secret-env")
	k, src, err := Load("")
	if err != nil {
		t.Fatal(err)
	}
	if src != SourceEnv || k != "secret-env" {
		t.Fatalf("got key=%q src=%v", k, src)
	}
}

func TestLoad_errNotConfigured(t *testing.T) {
	redirectConfigDir(t)
	t.Setenv(EnvAPIKey, "")
	_, _, err := Load("")
	if !errors.Is(err, ErrNotConfigured) {
		t.Fatalf("want ErrNotConfigured, got %v", err)
	}
}

func TestSave_emptyRejected(t *testing.T) {
	if err := Save(""); err == nil {
		t.Fatal("expected error for empty key")
	}
	if err := Save("   "); err == nil {
		t.Fatal("expected error for whitespace-only key")
	}
}

func TestSaveLoadDelete_roundtrip(t *testing.T) {
	redirectConfigDir(t)
	t.Setenv(EnvAPIKey, "")

	const secret = "roundtrip-test-key"
	if err := Save(secret); err != nil {
		t.Fatal(err)
	}

	k, src, err := Load("")
	if err != nil {
		t.Fatal(err)
	}
	if src != SourceFile || k != secret {
		t.Fatalf("got key=%q src=%v", k, src)
	}

	path, err := storedCredentialPath()
	if err != nil {
		t.Fatal(err)
	}
	if filepath.Base(path) != credentialsFile {
		t.Fatalf("unexpected credentials filename: %s", path)
	}

	if err := Delete(); err != nil {
		t.Fatal(err)
	}
	_, _, err = Load("")
	if !errors.Is(err, ErrNotConfigured) {
		t.Fatalf("after delete want ErrNotConfigured, got %v", err)
	}
}

func TestEnvConfigured(t *testing.T) {
	t.Setenv(EnvAPIKey, "x")
	if !EnvConfigured() {
		t.Fatal("expected true")
	}
	t.Setenv(EnvAPIKey, "")
	if EnvConfigured() {
		t.Fatal("expected false")
	}
}

func TestFileStoreConfigured_smoke(t *testing.T) {
	redirectConfigDir(t)
	t.Setenv(EnvAPIKey, "")
	if FileStoreConfigured() {
		t.Fatal("unexpected file store before save")
	}
	if err := Save("k"); err != nil {
		t.Fatal(err)
	}
	if !FileStoreConfigured() {
		t.Fatal("expected file store after save")
	}
}
