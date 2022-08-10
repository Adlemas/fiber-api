package routes

import (
	"errors"
	"time"

	"github.com/Adlemas/fiber-api/database"
	"github.com/Adlemas/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type OrderSerializer struct {
	ID        uint              `json:"id"`
	User      UserSerializer    `json:"user"`
	Product   ProductSerializer `json:"product"`
	CreatedAt time.Time         `json:"order_date"`
}

func CreateResponseOrder(orderModel models.Order, user UserSerializer, product ProductSerializer) OrderSerializer {
	return OrderSerializer{
		ID:        orderModel.ID,
		User:      user,
		Product:   product,
		CreatedAt: orderModel.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(404).JSON(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	var user models.User

	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(404).JSON(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	var product models.Product

	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(404).JSON(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}

	database.Database.Db.Find(&orders)

	responseOrders := []OrderSerializer{}

	for _, order := range orders {
		var user models.User
		var product models.Product

		findUser(order.UserRefer, &user)
		findProduct(order.ProductRefer, &product)

		responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(200).JSON(responseOrders)
}

func findOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}

	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(map[string]string{
			"message": "invalid id param",
		})
	}

	var order models.Order

	if err := findOrder(id, &order); err != nil {
		return c.Status(404).JSON(map[string]interface{}{
			"success": true,
			"message": err.Error(),
		})
	}

	var user models.User
	var product models.Product

	findUser(order.UserRefer, &user)
	findProduct(order.ProductRefer, &product)

	responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))

	return c.Status(200).JSON(responseOrder)
}
