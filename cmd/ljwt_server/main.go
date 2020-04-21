package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	helper "github.com/clsung/line-oauth2-helper"
	"github.com/gin-gonic/gin"
)

func main() {
	re := regexp.MustCompile(`^\d+$`)
	router := gin.Default()
	router.MaxMultipartMemory = 2 << 20 // 2 MB
	router.Static("/", "./static")
	router.POST("/gen", func(c *gin.Context) {
		channelID := c.PostForm("channelId")
		if !re.MatchString(channelID) {
			c.String(http.StatusBadRequest, fmt.Sprintf("Channel ID need digits, got %s", channelID))
			return
		}
		expiry, err := strconv.Atoi(c.PostForm("expiry"))
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("expiry error: %s", err.Error()))
			return
		}

		tokenExp, err := strconv.Atoi(c.PostForm("tokenExp"))
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("tokenExp error: %s", err.Error()))
			return
		}
		jsonContent := c.PostForm("content")

		h := helper.New(channelID).WithExpiry(time.Duration(expiry) * time.Minute).WithTokenExpire(tokenExp)

		var jwt string
		file, err := c.FormFile("file")
		if err != nil {
			if err != http.ErrMissingFile {
				c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
				return
			}
			jwt, err = h.GetLineJWT(strings.NewReader(jsonContent))
		} else {
			jwt, err = h.GetLineJWTFromFile(file.Filename)
		}
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("generate JWT err: %s", err.Error()))
			return
		}
		c.String(http.StatusOK, jwt)
	})
	router.Run(":8080")
}
