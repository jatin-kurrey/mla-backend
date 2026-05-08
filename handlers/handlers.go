package handlers

import (
	"mla-backend/models"
	"mla-backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
	"os"
)

func Login(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type LoginRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		var user models.User
		if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"exp":     time.Now().Add(time.Hour * 72).Unix(),
		})

		t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return c.SendStatus(500)
		}

		return c.JSON(fiber.Map{"token": t, "user": user})
	}
}

// Generic CRUD Handlers
func GetItems(db *gorm.DB, model interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db.Find(model)
		return c.JSON(model)
	}
}

func CreateItem(db *gorm.DB, model interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(model); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		db.Create(model)
		return c.JSON(model)
	}
}

func UpdateItem(db *gorm.DB, model interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := c.BodyParser(model); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		db.Model(model).Where("id = ?", id).Updates(model)
		return c.JSON(model)
	}
}

func DeleteItem(db *gorm.DB, model interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		db.Where("id = ?", id).Delete(model)
		return c.JSON(fiber.Map{"message": "Deleted successfully"})
	}
}

func UploadFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "No file uploaded"})
		}

		url, err := utils.UploadToR2(file)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"url": url})
	}
}

func DeleteFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type DeleteRequest struct {
			URL string `json:"url"`
		}
		var req DeleteRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		if err := utils.DeleteFromR2(req.URL); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"message": "File deleted from R2"})
	}
}
