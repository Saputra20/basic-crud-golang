package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	err := Init()

	assert.Nilf(err, "Expect nil, got %v", err)
}

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	os.Setenv("ENS_PG_HOST", "127.0.0.1")
	os.Setenv("ENS_PG_DBNAME", "learning")

	cfg := Config()
	t.Log(config)

	expect := "127.0.0.1"
	assert.Equalf(expect, cfg.PGHost, "Expect %v got %v", expect, cfg.PGHost)
}
