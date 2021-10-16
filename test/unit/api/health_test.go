package api

import (
	_ "postgre_mate_go/test/unit/suite"
)

import (
	"github.com/gogf/gf/frame/g"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestCreateHealth(t *testing.T) {
	server := g.Server()
	server.SetPort(10001)
	err := server.Start()
	assert.Nil(t, err)
	time.Sleep(100 * time.Millisecond)
	client := g.Client()
	postResp, err := client.Post("http://localhost:10001/v1/postgre/health")
	assert.Nil(t, err)
	defer postResp.Close()
	assert.Equal(t, http.StatusCreated, postResp.StatusCode)
	deleteResp, err := client.Delete("http://localhost:10001/v1/postgre/health")
	assert.Nil(t, err)
	defer deleteResp.Close()
	assert.Equal(t, http.StatusNoContent, deleteResp.StatusCode)
}
