package main

import (
	"fmt"
	"html/template"
	"net"
	"os"
	"path/filepath"
	"sync"

	_ "embed"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//go:embed index.html
var indexPage string
var indexPageTpl = template.Must(template.New("index").Parse(indexPage))

type Target struct {
	Engine *gin.Engine
}

func NewTarget(flag string, port int) *Target {
	var csrfMap sync.Map
	uu := uuid.New().String()
	os.MkdirAll(fmt.Sprintf("%s/test", uu), 0666)
	os.WriteFile(fmt.Sprintf("%s/flag", uu), []byte(flag), 0644)
	os.WriteFile(fmt.Sprintf("%s/test/test", uu), []byte("hello"), 0644)

	r := gin.Default()
	r.SetHTMLTemplate(indexPageTpl)

	r.GET("/", func(c *gin.Context) {
		csrfToken := uuid.New().String()
		from, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			c.String(400, "bad request")
			return
		}
		csrfMap.Store(from, csrfToken)
		c.HTML(200, "index", gin.H{"csrf": csrfToken})
	})
	r.POST("/", func(c *gin.Context) {
		from, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
		csrfToken, ok := csrfMap.Load(from)
		if !ok {
			c.String(400, "csrf")
			return
		}
		if c.PostForm("csrf") != csrfToken {
			c.String(400, "csrf")
			return
		}
		if c.PostForm("file") == "" {
			c.String(400, "filename is empty")
			return
		}
		fp := c.PostForm("file")
		fp = filepath.Join(uu+"/test", fp)
		fp = filepath.Join("/", fp)
		fp, err := filepath.Rel("/", fp)
		if err != nil {
			c.String(400, "filename is wrong")
			return
		}
		c.File(fp)
	})
	go r.Run(fmt.Sprintf(":%d", port))
	return &Target{r}
}
