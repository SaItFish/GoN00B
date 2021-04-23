// Package view
// @file: view.go
// @description:
// @author: SaltFish
// @date: 2020/09/20
package view

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"network-file/api/controller"

	"github.com/gin-gonic/gin"
)

type File struct {
	Name string `uri:"name" binding:"required"`
}

func StartServer() {
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	files := v1.Group("/files")
	{
		files.POST("/", func(c *gin.Context) {
			file, err := c.FormFile("file")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			n, err := controller.Upload(file)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"success": "Uploaded successfully",
				"name":    fmt.Sprintf("%s", n),
			})
		})
		files.GET("/:name", func(c *gin.Context) {
			var f File
			if err := c.ShouldBindUri(&f); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}
			m, cn, err := controller.Download(f.Name)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err})
				return
			}
			c.Header("Content-Disposition", "attachment; filename="+f.Name)
			c.Data(http.StatusOK, m, cn)
		})
		files.GET("/", func(c *gin.Context) {
			var filenames []string
			fs, _ := ioutil.ReadDir("/home/saltfish/go/src/github.com/SaItFish/GoN00B/sundries/file_rest_api_example/files")
			for _, f := range fs {
				filenames = append(filenames, f.Name())
			}
			if len(filenames) > 0 {
				c.JSON(http.StatusOK, filenames)
			} else {
				c.AbortWithStatus(http.StatusNotFound)
			}
		})
	}

	_ = router.Run("0.0.0.0:8080")
}
