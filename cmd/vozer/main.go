package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
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
)

func main() {
	flag.Parse()
	if *fVerbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	cfg := vozer.VozerConfig{
		ThreadURL:     *fThreadURL,
		NuWorkers:     *fNuWorkers,
		IsCrawlImages: *fCrawlImages,
		IsCrawlURLs:   *fCrawlURLs,
		DestPath:      *fDestPath,
		Retries:       *fRetries,
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

	select {
	case _, ok := <-ctx.Done():
		if ok {
			logrus.Infof("processing has been canceled")
			return
		}
	default:
	}

	logrus.Infof("crawled thread \"%s\" successfully in %v", cfg.ThreadURL, time.Since(start))
}
