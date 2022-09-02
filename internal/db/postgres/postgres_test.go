package postgres

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPGInit(t *testing.T) { // Integration test
	assert := assert.New(t)
	os.Setenv("ENS_PG_HOST", "localhost")
	os.Setenv("ENS_PG_PORT", "5432")
	os.Setenv("ENS_PG_DBNAME", "learning")
	os.Setenv("ENS_PG_USER", "postgres")
	os.Setenv("ENS_PG_PASSWORD", "postgres")
	PGInit()

	conn, err := PGConnection()

	assert.Nilf(err, "Expect nil got: %v", err)
	assert.NotNil(conn, "Expect not nil, got nil instead")
}
