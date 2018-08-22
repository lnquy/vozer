# vozer
CLI to crawl images and URLs from VOZ (https://forums.voz.vn) thread.  

![vozer_cli](https://github.com/lnquy/vozer/blob/master/vozer_cli.gif)

### Install
You can download compiled versions of `vozer` for Linux, Windows and Mac OS X from [Github release](https://github.com/lnquy/vozer/releases).  

Or install [Go](https://golang.org) SDK and build it.
```shell
$ go get github.com/lnquy/vozer/cmd/vozer

or

$ go get github.com/lnquy/vozer
$ cd $GOPATH/github.com/lnquy/vozer
$ dep ensure
$ cd cmd/vozer
$ go build
```

### Usage
```shell
$ vozer -h
Usage of vozer:
  -ci
    	Crawls images from posts or not
  -cu
    	Crawls URLs from posts or not
  -debug
    	Print debug log
  -o string
    	Path to directory where crawled data be saved to
  -pages string
    	List of page numbers to crawl data, separated by comma (,)
  -r uint
    	Number of time to re-crawl page if failed (default 20)
  -range string
    	Page number range to crawl data, separated by hyphen (-) (default "0-0")
  -u string
    	URL to VOZ thread
  -w uint
    	Number of workers to crawl data (default 10)
```

By default, `vozer` crawls all pages from thread and stores crawled data to `data` folder at current directory.  
Uses `-o` argument to save data on another folder.  
If you want to crawl specific page(s) then you can use `-pages` or `-range` argument.


Examples:
```shell
$ vozer -u https://forums.voz.vn/showthread.php?t=7382418 -ci   // Crawls images only
$ vozer -u https://forums.voz.vn/showthread.php?t=7382418 -cu -ci   // Crawls both images and URLs
$ vozer -u https://forums.voz.vn/showthread.php?t=7382418 -cu -ci -pages 1,3,10   // Crawls page 1, 3 and 10 only
$ vozer -u https://forums.voz.vn/showthread.php?t=7382418 -cu -ci -range 5-9   // Crawls from page 5 to page 9
$ vozer -u https://forums.voz.vn/showthread.php?t=7382418 -cu -ci -w 20 -r 10 -o ~/Desktop/voz -debug   // Output to ~/Desktop/voz folder rather than current directory
```

### License

This project is under the MIT License. See the [LICENSE](https://github.com/lnquy/vozer/blob/master/LICENSE) file for the full license text.

### Roadmap

- [x] Crawl pages by range (from page x to page y).
- [x] Crawl pages by list of numbers (input from CLI args or from file).
- [x] Filter emoticons to separated folder.


