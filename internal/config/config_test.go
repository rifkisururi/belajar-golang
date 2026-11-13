package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Setenv("APP_ADDR", "")
	t.Setenv("APP_LOG_LEVEL", "")
	c, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if c.Addr != ":8080" || c.LogLevel != "INFO" {
		t.Fatalf("config = %+v", c)
	}
}
