package wget

import (
	"fmt"
	"log"
	"os"
	"github.com/gocolly/colly"
)

func SaveFile(path string, b []byte) error {
	if err := os.WriteFile(path, b, 0644); err != nil {
		return err
	}
	return nil
}

func Wget(cfg *Config, logger *log.Logger) {
	c := colly.NewCollector(
		colly.AllowedDomains(cfg.Url.Host),
		colly.Async(true),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {

		if err := SaveFile("data", r.Body); err != nil {
			logger.Println(err)
		}
	})
	if err := c.Visit(cfg.Url.String()); err != nil {
		logger.Println(err)
	}
	c.Wait()
}