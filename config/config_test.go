package config

import "testing"

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("./../config.json")
	if err != nil {
		t.Fatalf("Error loading config: %v", err)
	}

	if config.Server.Host != "localhost" {
		t.Errorf("Expected host to be localhost, got %s", config.Server.Host)
	}

	if config.Server.Port != 8080 {
		t.Errorf("Expected port to be 8080, got %d", config.Server.Port)
	}
}
