package main

import (
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

type animal struct {
	Dog string `json:"dog" form:"dog"`
	Cat string `json:"cat" form:"cat"`
}

// 定义一个处理函数耗时的中间件
func timeMiddle(c *gin.Context) {
	fmt.Println("time in")
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	fmt.Println(cost)
	fmt.Println("time out")
}

func main() {
	animal := new(animal)
	// gin引擎实例
	r := gin.Default()
	r.Use(timeMiddle)
	// 路由方法
	// Query String
	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		// 返回json数据
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	// Form表单
	// 加载html网页
	r.LoadHTMLFiles("./login.html")
	r.GET("/login", func(c *gin.Context) {
		// 返回html网页
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "wyz" && password == "123" {
			c.JSON(http.StatusOK, gin.H{"login": "ok"})
		} else {
			c.ShouldBind(&animal)
			c.JSON(http.StatusOK, gin.H{
				"cat": animal.Cat,
				"dog": animal.Dog,
			})
		}
	})

	// 参数绑定
	r.GET("/bind", func(c *gin.Context) {
		c.ShouldBind(&animal)
		c.JSON(http.StatusOK, gin.H{
			"cat": animal.Cat,
			"dog": animal.Dog,
		})
	})
	r.POST("/form", func(c *gin.Context) {
		c.ShouldBind(&animal)
		c.JSON(http.StatusOK, gin.H{
			"cat": animal.Cat,
			"dog": animal.Dog,
		})
	})
	r.POST("/json", func(c *gin.Context) {
		c.ShouldBind(&animal)
		c.JSON(http.StatusOK, gin.H{
			"cat": animal.Cat,
			"dog": animal.Dog,
		})

	})
	r.POST("/upload", func(c *gin.Context) {
		// 读取前端文件
		f, _ := c.FormFile("filename")
		// dst := fmt.Sprintf("./%s", f.Filename)
		dst := path.Join("./", f.Filename)
		// 保存文件到本地
		c.SaveUploadedFile(f, dst)
	})

	// 中间件

	r.Run()
}
