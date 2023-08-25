package api

import (
	"os"
	"testing"
	"time"

	db "github.com/bagus-k/simple-bank/db/sqlc"
	"github.com/bagus-k/simple-bank/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server{
	config := utils.Config{
		TokenSymetricKey: utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
