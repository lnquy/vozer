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

	if c.NuWorkers == 0 || c.NuWorkers > 100 {
		c.NuWorkers = 10
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

	return nil
}
