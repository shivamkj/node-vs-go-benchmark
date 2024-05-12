package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

type User struct {
	ID         int     `json:"id"`
	Email      *string `json:"email"`
	Mobilenum  *string `json:"mobilenum"`
	Firstname  string  `json:"firstname"`
	Lastname   string  `json:"lastname"`
	Profilepic *string `json:"profilepic"`
}

func authMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	userID := claims["user_id"].(string)
	c.Locals("userID", userID)
	return c.Next()
}

func main() {
	// PostgreSQL connection
	connStr := "postgresql://shivam:pass@localhost:5432/qnify"
	pgxpool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pgxpool.Close()
	db = pgxpool

	// Create a new Fiber app
	app := fiber.New()

	// Route to get all users
	app.Get("/users", authMiddleware, func(c *fiber.Ctx) error {
		query := "SELECT id, email, mobilenum, firstname, lastname, profilepic FROM users LIMIT 10"
		rows, err := db.Query(context.Background(), query)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Email, &user.Mobilenum, &user.Firstname, &user.Lastname, &user.Profilepic)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Parsing error",
				})
			}
			users = append(users, user)
		}

		return c.JSON(users)
	})

	// Route to generate a JWT token
	app.Get("/login", func(c *fiber.Ctx) error {
		userID := "123" // Replace with actual user ID or any other payload you want to include in the token
		token, err := GenerateToken(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}

		return c.JSON(fiber.Map{
			"token": token,
		})
	})

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
