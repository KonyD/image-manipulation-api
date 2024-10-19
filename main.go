package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/nfnt/resize"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func main() {
	// Initialize a new Fiber app
	app := fiber.New()
	// localhost:3000/?width=100&height=50
	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		queries := c.Queries()

		mode := queries["mode"]

		width, err := strconv.ParseUint(queries["width"], 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid 'width' parameter.")
		}
		height, err := strconv.ParseUint(queries["height"], 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid 'height' parameter.")
		}

		image_url := queries["image"]
		response, err := http.Get(image_url)

		if err != nil || response.StatusCode != http.StatusOK {
			return c.Status(fiber.StatusBadGateway).SendString("Failed to fetch image")
		}
		defer response.Body.Close()

		img, format, err := image.Decode(response.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to decode image")
		}

		if mode == "resize" || mode == "" {
			img = resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
		} else if mode == "crop" {
			x, err := strconv.Atoi(queries["x"])
			if err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("Invalid 'x' parameter.")
			}
			y, err := strconv.Atoi(queries["y"])
			if err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("Invalid 'y' parameter.")
			}

			cropSize := image.Rect(0, 0, int(width), int(height))
			cropSize = cropSize.Add(image.Point{x, y})
			img = img.(SubImager).SubImage(cropSize)
		}

		buff := new(bytes.Buffer)

		switch strings.ToLower(format) {
		case "jpeg", "jpg":
			c.Set("Content-Type", "image/jpeg")
			err = jpeg.Encode(buff, img, nil) // Use jpeg.Encode for JPEG images
		case "png":
			c.Set("Content-Type", "image/png")
			err = png.Encode(buff, img) // Use png.Encode for PNG images
		default:
			return c.Status(fiber.StatusUnsupportedMediaType).SendString("Unsupported media type")
		}

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to encode image")
		}

		return c.SendStream(buff)
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
