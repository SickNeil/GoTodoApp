// main.go
package main

import (
	"context"
	"go-todo-app/models"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var todoModel *models.TodoModel

func main() {
	// 連接 MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://todo-database:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// 選擇資料庫和集合
	todoCollection := client.Database("tododb").Collection("todos")
	todoModel = &models.TodoModel{
		Collection: todoCollection,
	}

	// 初始化 Gin
	router := gin.Default()

	// 加載模板，並添加自定義函數
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLGlob("templates/*")

	// 設定路由
	router.GET("/", showTodos)
	router.POST("/", addTodo)
	router.POST("/delete/:id", deleteTodo)

	// 啟動服務器
	router.Run(":3000")
}

// 格式化時間的函數
func formatAsDate(t time.Time) string {
	// 創建一個 UTC+8 的時區
	location := time.FixedZone("UTC+8", 8*60*60)

	// 將時間轉換為 UTC+8
	utc8Time := t.In(location)

	// 格式化為您需要的格式
	return utc8Time.Format("2006-01-02 15:04:05")
}

// 顯示待辦事項列表
func showTodos(c *gin.Context) {
	todos, err := todoModel.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "todos.html", gin.H{
		"tasks": todos,
	})
}

// 添加新的待辦事項
func addTodo(c *gin.Context) {
	task := c.PostForm("task")
	if task == "" {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	todo := models.Todo{
		Task:      task,
		CreatedAt: time.Now(),
	}
	err := todoModel.Insert(todo)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}

// 刪除待辦事項
func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	err := todoModel.Delete(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}
