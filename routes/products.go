package routes

import (
	"errors"

	"github.com/Adlemas/fiber-api/database"
	"github.com/Adlemas/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type ProductSerializer struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) ProductSerializer {
	return ProductSerializer{
		ID:           productModel.ID,
		Name:         productModel.Name,
		SerialNumber: productModel.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(map[string]string{
			"message": err.Error(),
		})
	}

	database.Database.Db.Create(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	database.Database.Db.Find(&products)

	responseProducts := []ProductSerializer{}

	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}

	return c.Status(200).JSON(responseProducts)
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("product does not exist")
	}

	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(map[string]string{
			"message": "invalid id param",
		})
	}

	var product models.Product

	if err := findProduct(id, &product); err != nil {
		return c.Status(404).JSON(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(map[string]string{
			"message": "invalid id param",
		})
	}

	var product models.Product

	if err := findProduct(id, &product); err != nil {
		return c.Status(404).JSON(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	type UpdateProduct struct {
		Name         *string `json:"name"`
		SerialNumber *string `json:"serial_number"`
	}

	var updateData UpdateProduct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(map[string]string{
			"message": err.Error(),
		})
	}

	if updateData.Name != nil {
		product.Name = *updateData.Name
	}

	if updateData.SerialNumber != nil {
		product.SerialNumber = *updateData.SerialNumber
	}

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(map[string]string{
			"message": "invalid id param",
		})
	}

	var product models.Product

	if err := findProduct(id, &product); err != nil {
		return c.Status(404).JSON(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(map[string]interface{}{
		"success": true,
	})
}
