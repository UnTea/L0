package main

import (
	"database/sql" // add this
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // add this

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

const connStr = "postgresql://test_user:test_password@localhost:5432/test?sslmode=disable"

type todo struct {
	Item string
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string

	rows, err := db.Query("select * from test")

	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}
	return c.Render("index", fiber.Map{
		"Todos": todos,
	})
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	newTodo := todo{}
	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("%v", newTodo)
	if newTodo.Item != "" {
		_, err := db.Exec("INSERT into test VALUES ($1)", newTodo.Item)
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
	}

	return c.Redirect("/")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	olditem := c.Query("olditem")
	newitem := c.Query("newitem")
	db.Exec("UPDATE test SET item=$1 WHERE item=$2", newitem, olditem)
	return c.Redirect("/")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	todoToDelete := c.Query("item")
	db.Exec("DELETE from test WHERE item=$1", todoToDelete)
	return c.SendString("deleted")
}

func main() {
	//clusterID := "test-cluster"
	//clientID := "420-69"
	//
	//sc, err := stan.Connect(clusterID, clientID)
	//
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//}
	//
	//// Simple Async Subscriber
	//sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
	//	fmt.Printf("Received a message: %s\n", string(m.Data))
	//})
	//
	//// Simple Synchronous Publisher
	//err = sc.Publish("foo", []byte("Hello World"))
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//} // does not return until an ack has been received from NATS Streaming
	//
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//}
	//
	//defer func() {
	//	// Unsubscribe
	//	err = sub.Unsubscribe()
	//	if err != nil {
	//		fmt.Printf("%v\n", err)
	//	}
	//
	//	// Close connection
	//	err = sc.Close()
	//	if err != nil {
	//		fmt.Printf("%v\n", err)
	//	}
	//}()
	//
	//r := gin.Default()
	//
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Static("/", "./public")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))

}
