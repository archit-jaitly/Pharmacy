package main

import (
	"Pharmacy/controller"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	database := client.Database("pharmacy")
	medicinesCollection := database.Collection("medicines")

	medicineController := controller.NewMedicineController(medicinesCollection)

	// Routes
	app.Post("/pharmacy/addMedicine", medicineController.AddMedicine)

	app.Put("/pharmacy/updateMedicine/:id", medicineController.UpdateMedicineDetails)

	app.Put("/pharmacy/increment/:id", medicineController.IncrementStock)

	app.Put("/pharmacy/decrement/:id", medicineController.DecrementStock)

	app.Listen(":8000")
}
