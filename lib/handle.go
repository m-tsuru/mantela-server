package lib

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	mantela_fetcher "github.com/tkytel/mantela-adder/fetcher"
)

func HandleMantela(c *fiber.Ctx, source string, diff string) error {
	// Create a new HTTP client and request to fetch the Mantela JSON
	client := &http.Client{}
	req, err := http.NewRequest("GET", source, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching Mantela")
	}
	r, err := client.Do(req)
	if err != nil {
		log.Printf("Error fetching Mantela: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching Mantela")
	}
	defer r.Body.Close()
	resp, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching Mantela")
	}

	m1, err := mantela_fetcher.ParseMantela(resp)
	if err != nil {
		log.Printf("Error parsing Mantela: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching Mantela")
	}

	// Get Local Mantela that the user wants to merge.
	if diff == "" {
		return c.JSON(m1)
	}
	data, err := os.ReadFile(diff)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error reading override file")
	}
	m2, err := mantela_fetcher.ParseMantela([]byte(data))
	if err != nil {
		log.Printf("Error parsing local Mantel (Returning is Unmerged Mantela): %v", err)
		return c.JSON(m1)
	}

	// Merge the two Mantela structures
	m := mantela_fetcher.MergeMantela(m1, m2)
	return c.JSON(m)
}
