package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/zhao2490/web-frame/gee"
	"github.com/zhao2490/web-frame/middleware"
)

func main() {
	stepRouter()
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func stepRouter() {
	r := gee.New()
	r.Use(middleware.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
		})
		v1.GET("/students", func(c *gee.Context) {
			type student struct {
				Name string
				Age  int8
			}
			students := []student{
				{Name: "zhao", Age: 12},
				{Name: "zzzz", Age: 23},
			}
			c.HTML(http.StatusOK, "student.html", gee.H{
				"title":  "student",
				"stuArr": students,
			})
		})
	}
	v2 := r.Group("/v2")
	{
		v2.Use(gee.Recovery())
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you are at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
		v2.GET("/panic", func(c *gee.Context) {
			mm := make([]uint32, 0)
			_ = mm[12]
			c.String(http.StatusOK, "panic")
		})
	}
	r.Run(":9999")
}
