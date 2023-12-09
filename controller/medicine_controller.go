package controller

import (
	"Pharmacy/model"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MedicineController struct {
	MedicinesCollection *mongo.Collection
	SMS                 AbstractSMSController
}

func NewMedicineController(collection *mongo.Collection) *MedicineController {
	return &MedicineController{MedicinesCollection: collection}
}

func (mc *MedicineController) AddMedicine(c *fiber.Ctx) error {

	medicine := new(model.Medicine)
	if err := c.BodyParser(medicine); err != nil {
		return err
	}

	uuid := uuid.New()
	medicine.SerialNumber = uuid.String()
	_, err := mc.MedicinesCollection.InsertOne(context.Background(), medicine)
	if err != nil {
		errMsg := fmt.Sprintf("Add medicine failed: %s", err.Error())
		return fmt.Errorf(errMsg)
	}

	mc.SMS.SendSMS("pharmacy Owner", fmt.Sprintf("New medicine added: %s  id: %s  serial number: %s", medicine.Name, medicine.ID, medicine.SerialNumber))
	status := mc.SMS.CheckSMSStatus(medicine.ID)

	return c.JSON(fiber.Map{
		"status": status,
	})
}

func (mc *MedicineController) UpdateMedicineDetails(c *fiber.Ctx) error {

	medicineID := c.Params("id")
	medicine := new(model.Medicine)
	if err := c.BodyParser(medicine); err != nil {
		return err
	}

	filter := bson.M{"_id": medicineID}
	update := bson.M{"$set": medicine}

	var result model.Medicine
	err := mc.MedicinesCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		errMsg := fmt.Sprintf("Update medicine details failed, medicine not found: %s", err.Error())
		return fmt.Errorf(errMsg)
	}

	_, err = mc.MedicinesCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		errMsg := fmt.Sprintf("Update medicine details failed: %s", err.Error())
		return fmt.Errorf(errMsg)
	}

	return c.JSON(fiber.Map{
		"status": "success",
	})
}

func (mc *MedicineController) IncrementStock(c *fiber.Ctx) error {

	medicineID := c.Params("id")
	filter := bson.M{"_id": medicineID}
	update := bson.M{"$inc": bson.M{"quantity": 1}}
	resp, err := mc.MedicinesCollection.UpdateOne(context.Background(), filter, update)
	if resp.ModifiedCount == 0 {
		return fmt.Errorf("Increment stock failed, no medicine found")
	}
	if err != nil {
		errMsg := fmt.Sprintf("Increment stock failed: %s", err.Error())
		return fmt.Errorf(errMsg)
	}

	return c.JSON(fiber.Map{
		"status": "success",
	})
}

func (mc *MedicineController) DecrementStock(c *fiber.Ctx) error {

	medicineID := c.Params("id")
	filter := bson.M{"_id": medicineID, "quantity": bson.M{"$gt": 0}}
	update := bson.M{"$inc": bson.M{"quantity": -1}}
	resp, err := mc.MedicinesCollection.UpdateOne(context.Background(), filter, update)
	if resp.ModifiedCount == 0 {
		return fmt.Errorf("Decrement stock failed, no stock found")
	}

	if err != nil {
		errMsg := fmt.Sprintf("Decrement stock failed: %s", err.Error())
		return fmt.Errorf(errMsg)
	}

	return c.JSON(fiber.Map{
		"status": "success",
	})
}
