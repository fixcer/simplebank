package api

import (
	db "github.com/fixcer/simplebank/db/sqlc"
	"github.com/fixcer/simplebank/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: 15 * time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}
