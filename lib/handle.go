package lib

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	mantela_fetcher "github.com/tkytel/mantela-adder/fetcher"
)

func fetch(path string) ([]byte, error) {
	// if path starts http:// or https://, then fetch the Mantela JSON from the internet.
	// else, read the Mantela JSON from the local file system.
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		// Create a new HTTP client and request to fetch the Mantela JSON
		client := &http.Client{}
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			log.Printf("Error creating request: %v", err)
			return nil, err
		}
		r, err := client.Do(req)
		if err != nil {
			log.Printf("Error fetching Mantela: %v", err)
			return nil, err
		}
		defer r.Body.Close()
		resp, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			return nil, err
		}
		return resp, nil
	} else {
		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error reading file: %v", err)
			return nil, err
		}
		return data, nil
	}
}

func HandleMantela(c *fiber.Ctx, source string, diff string) error {
	r1, err := fetch(source)
	if err != nil {
		log.Printf("%v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching Mantela")
	}

	m1, err := mantela_fetcher.ParseMantela(r1)
	if err != nil {
		log.Printf("Error parsing Mantela: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching Mantela")
	}

	// Get Local Mantela that the user wants to merge.
	if diff == "" {
		return c.JSON(m1)
	}
	r2, err := fetch(diff)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error reading override file")
	}
	m2, err := mantela_fetcher.ParseMantela(r2)
	if err != nil {
		log.Printf("Error parsing local Mantel (Returning is Unmerged Mantela): %v", err)
		return c.JSON(m1)
	}

	// Merge the two Mantela structures
	m := mantela_fetcher.MergeMantela(m1, m2)
	return c.JSON(m)
}
