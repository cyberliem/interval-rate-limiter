package ratelimiter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"testing"
	"time"
)

func setUpTestServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func TestRateLimit(t *testing.T) {
	go setUpTestServer()
	time.Sleep(3 * time.Second)
	testIRL := NewIntervalRateLimiter(1000 * time.Millisecond)
	const timeout = time.Second
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/ping", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(index int) {
			fmt.Printf("doing it")
			_, err := testIRL.ForwardRequest(client, req)
			fmt.Printf("request %d done at time %s, error %v \n", index, time.Now().Format(time.StampMilli), err)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
