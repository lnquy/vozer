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
  -r uint
    	Number of times to re-crawl page if failed (default 20)
  -u string
    	URL to VOZ thread
  -w uint
    	Number of workers to crawl data (default 10)
    	
Examples:
  vozer -u https://forums.voz.vn/showthread.php?t=3116194 -cu -ci
  vozer -u https://forums.voz.vn/showthread.php?t=3116194 -cu -ci -w 20 -r 10 -o ~/Desktop/voz -debug
```

### License
This project is under the MIT License. See the [LICENSE](https://github.com/lnquy/vozer/blob/master/LICENSE) file for the full license text.
