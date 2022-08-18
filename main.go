package main

import (
	"context"
	"fmt"
	"go2/handler"
	"go2/uid"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	people = uid.New()
	// msgcontrol controller.MsgController = controller.New(msgs)
)

//	func ShowAll(ctx *gin.Context) {
//		datas := [2]string{"ok", "hi"}
//		data := gin.H{
//			"datas": datas,
//		}
//		ctx.HTML(http.StatusOK, "index.html", data)
//	}
func main() {
	server := gin.New()

	// authcred := options.Credential{
	// 	// AuthMechanism: "testdb_local",
	// 	Username: "testdbLocal",
	// 	Password: "testdbLocal",
	// }
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	opts := options.Client().ApplyURI("mongodb://testdbLocal:testdbLocal@localhost:27017/?authSource=testdb_local")

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)
	dbNames, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	fmt.Println(dbNames)

	// fmt.Printf("%T\n", client)

	testDB := client.Database("testdb_local")
	// fmt.Printf("%T\n", testDB)
	JSONData := struct {
		Path string `bson:"Path"`
	}{}
	exampleCollection := testDB.Collection("testing")
	// fmt.Println(exampleCollection.)
	// var ok []string
	type Contact struct {
		Name  string `bson:"name,omitempty"`
		Age   int64  `bson:"age,omitempty"`
		Greet string `bson:"greet,omitempty"`
	}
	var contacts []Contact
	cur, _ := exampleCollection.Find(ctx, bson.M{})
	// cur.All(ctx, &ok)
	// fmt.Print(ok)
	if err = cur.All(ctx, &contacts); err != nil {
		panic(err)
	}

	for _, c := range contacts {
		fmt.Println(c.Name)
		fmt.Printf("%v\n", c.Age)
		fmt.Println(c.Greet)
		println()
	}

	fmt.Printf("%T\n", cur.Decode(&JSONData))
	defer exampleCollection.Drop(ctx)

	// fmt.Printf("%T\n", exampleCollection)
	// example := bson.D{
	// 	{"someString", "Example String"},
	// 	{"someInteger", 12},
	// 	{"someStringSlice", []string{"Example 1", "Example 2", "Example 3"}},
	// }
	// r, err := exampleCollection.InsertOne(ctx, example)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(r.InsertedID)
	// examples := []interface{}{
	// 	bson.D{
	// 		{"someString", "Second Example String"},
	// 		{"someInteger", 253},
	// 		{"someStringSlice", []string{"Example 15", "Example 42", "Example 83", "Example 5"}},
	// 	},
	// 	bson.D{
	// 		{"someString", "Another Example String"},
	// 		{"someInteger", 54},
	// 		{"someStringSlice", []string{"Example 21", "Example 53"}},
	// 	},
	// }
	// rs, err := exampleCollection.InsertMany(ctx, examples)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(rs.InsertedIDs)

	// time.Sleep(20 * time.Second)

	os.Exit(1)
	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	// server.Use(gin.Recovery(), middlewares.BasicAuth())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/hello", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"EHEHHE": "asdhb",
			})
		})

		// apiRoutes.GET("/videos", func(context *gin.Context) {
		// 	context.JSON(200, VideoController.FindAll())
		// })

		// apiRoutes.POST("/postvid", func(ctx *gin.Context) {
		// 	err := VideoController.Save(ctx)
		// 	if err != nil {
		// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	} else {
		// 		ctx.JSON(http.StatusOK, gin.H{
		// 			"message": "Input validated",
		// 		})
		// 	}

		// })
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/getall", handler.GetPerson(people))
		viewRoutes.POST("/add", handler.AddPerson(people))
		viewRoutes.POST("/ping", handler.Pingg(people))
	}

	server.Run(":6969")
}
