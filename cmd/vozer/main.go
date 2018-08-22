package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/lnquy/vozer"
	"github.com/sirupsen/logrus"
)

const VOZER_VERSION = "0.0.3"

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
	fVersion     = flag.Bool("v", false, "Print vozer version and exit")
)

func main() {
	flag.Parse()

	if *fVersion {
		fmt.Printf("vozer-%s-%s_v%s\n", runtime.GOOS, runtime.GOARCH, VOZER_VERSION)
		os.Exit(0)
	}

	if *fVerbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	crawlRange := strings.Split(*fCrawlRange, "-")
	if len(crawlRange) != 2 {
		logrus.Errorf("Invalid page range: %s", *fCrawlRange)
		os.Exit(1)
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
		os.Exit(1)
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
		os.Exit(1)
	}

	if ctx.Err() != nil {
		logrus.Infof("operation cancelled by user")
		os.Exit(0)
	}
	logrus.Infof("crawled thread \"%s\" successfully in %v", cfg.ThreadURL, time.Since(start))
}

func parseUint(s string) uint {
	u, _ := strconv.ParseUint(strings.TrimSpace(s), 10, 64)
	return uint(u)
}
