package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var flags []string

const TargetCount = 3

func main() {
	r := gin.Default()

	for i := 0; i < TargetCount; i++ {
		flag := fmt.Sprintf("flag{%s}", uuid.New().String())
		flags = append(flags, flag)
		NewTarget(flag, 8081+i)
	}

	// platform
	r.POST("/pushflag", func(c *gin.Context) {
		flag := c.PostForm("flag")
		if flag == "" {
			c.String(400, "flag is empty")
			return
		}
		for _, f := range flags {
			if f == flag {
				c.String(200, "success")
				return
			}
		}
		c.String(400, "flag wrong")
	})

	r.Run(":8080")

}
