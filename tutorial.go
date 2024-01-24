package main
//https://docs.sqlc.dev/en/stable/tutorials/getting-started-mysql.html
import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"strconv"

	"github.com/gin-gonic/gin"
	"tutorial.sqlc.dev/app/tutorial"
)

func create(c *gin.Context) {
	ctx := context.Background()
	var author tutorial.Author
	c.BindJSON(&author)
	db, _ := sql.Open("mysql", "root:3435@/auth?parseTime=true")
	queries := tutorial.New(db)

	result, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: author.Name,
		Bio:  sql.NullString{String: author.Bio.String, Valid: true},
	})
	fmt.Println(err)
	fmt.Println(result)
	c.IndentedJSON(http.StatusCreated, author)
}

func getAuthor(c *gin.Context) {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "root:3435@/auth?parseTime=true")
	queries := tutorial.New(db)
	int1, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	fmt.Println(int1)
	fetchedAuthor, err := queries.GetAuthor(ctx, int1)
	fmt.Println(err)
	c.IndentedJSON(http.StatusOK, fetchedAuthor)
}

func delete(c *gin.Context) {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "root:3435@/auth?parseTime=true")
	queries := tutorial.New(db)
	int1, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	auth := queries.DeleteAuthor(ctx, int1)
	c.IndentedJSON(http.StatusOK, auth)
}

func authers(c *gin.Context) {
	ctx := context.Background()
	db, _ := sql.Open("mysql", "root:3435@/auth?parseTime=true")
	queries := tutorial.New(db)
	arr, _ := queries.ListAuthors(ctx)
	c.IndentedJSON(http.StatusOK, arr)
}

func main() {

	router := gin.Default()
	router.POST("/insert", create)
	router.GET("/get/:id", getAuthor)
	router.DELETE("/delete/:id", delete)
	router.GET("/get", authers)
	router.Run("localhost:8080")
}
