package sqlite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgreSql(t *testing.T) {
	assert.NotNil(t, New("sqlite"))
}
