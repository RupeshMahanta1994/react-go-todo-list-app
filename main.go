package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello world")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error in loading env file", err)
	}
	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error in mongo connection", err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error in pinging the Mongo uri", err)
	}
	fmt.Println("Connected to Mongo db atlas")

	collection = client.Database("golang_db").Collection("todos")

	app := fiber.New()

	//app.Use(cors.New(cors))

	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodos)
	app.Patch("/api/todos/:id", updateTodos)
	app.Delete("/api/todos/:id", deleteTodos)

	port := os.Getenv("PORT")

	log.Fatal(app.Listen(":" + port))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal("Error in finding collection in db", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			log.Fatal("Error in getting individual todo", err)
		}
		todos = append(todos, todo)
	}
	return c.JSON(todos)
}

func createTodos(c *fiber.Ctx) error {
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		log.Fatal("Error in parsing the body in create todo", err)
	}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Fatal("Error in inserting the todo", err)
	}
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.Status(201).JSON(todo)

}

func updateTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid tot ID"})
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}
	updatedTodo, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal("Error in updating todo", err)
	}

	return c.Status(201).JSON(updatedTodo)

}

func deleteTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}
	filter := bson.M{"_id": objectID}
	deleteTodo, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal("Error in deleting todo", err)
	}

	return c.Status(201).JSON(deleteTodo)
}

// type Todo struct {
// 	ID        int    `json:"id"`
// 	Completed bool   `json:"completed"`
// 	Body      string `json:"body"`
// }

// func main() {
// 	fmt.Println("Hello, Banglore!")
// 	app := fiber.New()

// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Error in loading .env file")
// 	}
// 	PORT := os.Getenv("PORT")

// 	todos := []Todo{}

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.Status(200).JSON(todos)
// 	})

// 	//Create a Todo
// 	app.Post("/api/todos", func(c *fiber.Ctx) error {
// 		todo := &Todo{}
// 		if err := c.BodyParser(todo); err != nil {
// 			return err
// 		}
// 		if todo.Body == "" {
// 			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
// 		}

// 		todo.ID = len(todos) + 1
// 		todos = append(todos, *todo)

// 		return c.Status(201).JSON(todo)
// 	})

// 	// Update a TOdo
// 	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")

// 		for i, todo := range todos {
// 			if fmt.Sprint(todo.ID) == id {
// 				todos[i].Completed = true
// 				return c.Status(200).JSON(todos[i])
// 			}

// 		}
// 		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})

// 	})

// 	//Delete a Todo
// 	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		for i, todo := range todos {
// 			if fmt.Sprint(todo.ID) == id {
// 				todos = append(todos[:i], todos[i+1:]...)
// 				return c.Status(200).JSON(fiber.Map{"Success": "true"})
// 			}
// 		}
// 		return c.Status(404).JSON(fiber.Map{"msg": "Todo not found"})
// 	})
// 	log.Fatal(app.Listen(":" + PORT))
// }
