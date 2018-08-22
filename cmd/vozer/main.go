package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/lnquy/vozer"
	"github.com/sirupsen/logrus"
)

var (
	fThreadURL   = flag.String("u", "", "URL to VOZ thread")
	fNuWorkers   = flag.Uint("w", 10, "Number of workers to crawl data")
	fCrawlURLs   = flag.Bool("cu", false, "Crawls URLs from posts or not")
	fCrawlImages = flag.Bool("ci", false, "Crawls images from posts or not")
	fDestPath    = flag.String("o", "", "Path to directory where crawled data be saved to")
	fRetries     = flag.Uint("r", 20, "Number of time to re-crawl page if failed")
	fVerbose     = flag.Bool("debug", false, "Print debug log")
	fCrawlRange  = flag.String("range", "0-0", "Page range to crawl data, separated by hyphen (-)")
	fCrawlPages  = flag.String("pages", "", "List of page numbers to crawl data, separated by comma (,)")
)

func main() {
	flag.Parse()
	if *fVerbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	crawlRange := strings.Split(*fCrawlRange, "-")
	if len(crawlRange) != 2 {
		logrus.Errorf("Invalid page range: %s", *fCrawlRange)
		return
	}

	var pages []uint
	if *fCrawlPages != "" {
		s := strings.Split(*fCrawlPages, ",")
		for _, p := range s {
			pages = append(pages, parseUint(p))
		}
	}

	cfg := vozer.VozerConfig{
		ThreadURL:     *fThreadURL,
		NuWorkers:     *fNuWorkers,
		IsCrawlImages: *fCrawlImages,
		IsCrawlURLs:   *fCrawlURLs,
		DestPath:      *fDestPath,
		Retries:       *fRetries,
		CrawlPages:    pages,
		CrawlFromPage: parseUint(crawlRange[0]),
		CrawlToPage:   parseUint(crawlRange[1]),
	}
	if err := cfg.Validate(); err != nil {
		logrus.Error(err)
		flag.PrintDefaults()
		return
	}

	start := time.Now()
	ctx, ctxCancel := context.WithCancel(context.Background())

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
		<-sig
		ctxCancel()
	}()

	if err := vozer.Crawl(ctx, cfg); err != nil {
		logrus.Errorf("failed to crawl \"%s\": %s", cfg.ThreadURL, err)
		return
	}

	if ctx.Err() != nil {
		logrus.Infof("operation cancelled by user")
		return
	}
	logrus.Infof("crawled thread \"%s\" successfully in %v", cfg.ThreadURL, time.Since(start))
}

func parseUint(s string) uint {
	u, _ := strconv.ParseUint(strings.TrimSpace(s), 10, 64)
	return uint(u)
}
