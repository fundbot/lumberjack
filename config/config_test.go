package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_readingConfig(t *testing.T) {
	Load()

	name := Name()
	assert.Equal(t, name, "Lumberjack_Test")

	port := Port()
	assert.Equal(t, port, 9999)

	logLevel := LogLevel()
	assert.Equal(t, logLevel, "development")

	baseURL := BaseURL()
	assert.Equal(t, baseURL, "https://leftshift.io")

	delay := Delay()
	assert.Equal(t, delay, 20)

	version := Version()
	assert.Equal(t, version, "0.1.0")

	app := Application()
	assert.Equal(t, app.name, "Lumberjack_Test")
}
