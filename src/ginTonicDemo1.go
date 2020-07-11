package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func RepArmandXU() {
	fmt.Println("一个网站上的教程，简单")

	r := gin.Default()                        //我们使用 gin.Default() 生成了一个实例，这个实例即 WSGI 应用程序
	r.GET("/ArmandXU", func(c *gin.Context) { // 声明了一个路由，告诉Gin什么样的URL能触发传入的函数
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello ArmandXU !",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to my website !",
		})
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务 r.Run(":9999")可以运行在_9999_端口
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}

func BasicService() { // "只需要两行代码就可以让服务跑起来"	我们只要点击运行，项目便会启动一个 8080 端口，打开浏览器 localhost:8080
	// 我们便可以看到页面上提示出 404 page not found ，这是因为我们的根路由上并没有返回任何结果。同时我们可以在控制台上看到一些打印信息，
	//其中就包括我们刚刚访问根路由的端口。
	// https://studygolang.com/articles/21883
	router := gin.Default()
	router.Run()
}

func BasicServiceWithOneGET() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Gin !")
	})
	// POST,PUT,DELETE 还可以实现这几个

	router.POST("/post", func(c *gin.Context) {
		c.String(200, "Bienvenue POST !")
	})

	router.Run()
}

func XMLJSONYAMLPROTOBUFRender() { // https://learnku.com/docs/gin-gonic/2019/examples-rendering/6168 XML/JSON/YAML/ProtoBuf 渲染
	r := gin.Default()

	// gin.H 是 map[string]interface{} 的一种快捷方式
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// 也可以使用一个结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}

		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意 msg.Name 在 JSON 中变成了 "user"
		// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}

		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})

	r.Run()
}

func HTMLRender() { // 这段代码的作用是，返回一个本地（服务器端）的html文件并在客户端渲染
	// 有模版功能！（template）
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/tutorial", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tutorial.html", gin.H{})
	})

	router.GET("/tutorial/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
		//c.HTML(http.StatusOK, "tutorial.html", gin.H{})

	})

	router.Run(":8080")
}
