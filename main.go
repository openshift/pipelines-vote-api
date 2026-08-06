package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

var inMemoryStore = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/vote", func(c *gin.Context) {
		payload := gin.H{}
		voteA := 0
		voteB := 0
		for _, v := range inMemoryStore {
			switch v {
			case "a":
				voteA++
			case "b":
				voteB++
			}
		}
		payload["a"] = voteA
		payload["b"] = voteB
		c.JSON(http.StatusOK, payload)
	})

	r.GET("/vote/total", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"total": len(inMemoryStore)})
	})

	r.POST("/vote", func(c *gin.Context) {
		buf := make([]byte, 1024)
		num, _ := c.Request.Body.Read(buf)
		reqBody := buf[0:num]
		temp := map[string]string{}
		json.Unmarshal(reqBody, &temp)
		c.JSON(http.StatusOK, reqBody)
		voter_id := temp["voter_id"]
		vote := temp["vote"]
		inMemoryStore[voter_id] = vote
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":9000")
}
