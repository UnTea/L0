package main

import (
	"database/sql" // add this
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"time"

	_ "github.com/lib/pq" // add this

	"github.com/gofiber/fiber/v2"
)

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
		_, err := db.Exec("insert into test values ($1)", newTodo.Item)
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

type Model struct {
	OrderUid          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Items   `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Items struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func (model *Model) print() {
	s := reflect.ValueOf(model).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		fmt.Printf("%s: %v\n", typeOfT.Field(i).Name, f.Interface())
	}
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

	connString := "postgresql://test_user:test_password@localhost:5432/test?sslmode=disable"

	// Connect to database
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

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

	//app.Static("/", "./public")
	//log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))

	jsonFilePath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	jsonFilePath = jsonFilePath + "/info/model.json"

	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		err = jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	modelPure, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	model := Model{}

	err = json.Unmarshal(modelPure, &model)
	if err != nil {
		fmt.Println(err)
	}
}
