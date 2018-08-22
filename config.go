package vozer

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
)

type VozerConfig struct {
	ThreadURL     string `json:"thread_url"`
	NuWorkers     uint   `json:"workers"`
	IsCrawlURLs   bool   `json:"is_crawl_urls"`
	IsCrawlImages bool   `json:"is_crawl_images"`
	DestPath      string `json:"destination_path"`
	Retries       uint   `json:"retries"`
	CrawlPages    []uint `json:"crawl_pages"`
	CrawlFromPage uint   `json:"crawl_from_page"`
	CrawlToPage   uint   `json:"crawl_to_page"`
}

func (c *VozerConfig) Validate() error {
	if c.ThreadURL == "" {
		return errors.New("URL to VOZ thread must be specified")
	}
	u, err := url.Parse(c.ThreadURL)
	if err != nil {
		return fmt.Errorf("Invalid URL: %s", err)
	}
	if u.Host != "forums.voz.vn" {
		return errors.New("Invalid URL, must point to a VOZ thread")
	}

	if c.NuWorkers == 0 {
		c.NuWorkers = 10
	}
	if c.NuWorkers > 100 {
		c.NuWorkers = 100
	}

	if !c.IsCrawlURLs && c.IsCrawlImages {
		return errors.New("Must specify which data you want to crawl? (images, URLs or both)")
	}

	if c.DestPath == "" {
		dp, _ := os.Getwd()
		c.DestPath = path.Join(dp, "data")
	}

	if c.Retries > 50 {
		c.Retries = 50
	}

	for i := range c.CrawlPages {
		if c.CrawlPages[i] == 0 {
			c.CrawlPages = append(c.CrawlPages[:i], c.CrawlPages[i+1:]...)
		}
	}

	if c.CrawlFromPage > c.CrawlToPage {
		return fmt.Errorf("Invalid page range: %d-%d", c.CrawlFromPage, c.CrawlToPage)
	}

	return nil
}
