package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var sortedList []string

func main() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
	}

	pathTranco, pathURLs, threshold, threads := parseFlags()

	sortedList = make([]string, threshold+1)

	if pathTranco == "" || pathURLs == "" {
		fmt.Printf("%sError: -tranco (-t) or -urls (-u) wasn't specified%s\n", Red, Reset)
		os.Exit(1)
	}

	sliceTranco := readLocalFile(pathTranco)
	sliceURLs := readLocalFile(pathURLs)

	fmt.Printf("Checking %d URLs for their ranking\n", len(sliceURLs))

	sem := make(chan int, threads)
	var wg sync.WaitGroup
	wg.Add(len(sliceURLs))
	var m sync.Mutex

	count := 0

	for _, url := range sliceURLs {
		if url == "" {
			wg.Done()
			continue
		}
		go func(url string) {
			defer wg.Done()
			sem <- 1
			for _, tranco := range sliceTranco {
				if tranco == "" {
					continue
				}

				trancoSplitted := strings.Split(tranco, ",")
				if len(trancoSplitted) != 2 {
					fmt.Printf("%sError: length != 2: %s%s\n", Red, trancoSplitted, Reset)
				} else {
					trancoURL := strings.TrimSpace(trancoSplitted[1])
					url := strings.TrimSpace(url)
					if url == trancoURL {
						rank, err := strconv.Atoi(trancoSplitted[0])
						if err != nil {
							fmt.Println(Red + tranco + ": " + err.Error() + Reset)
							<-sem
							return
						}
						if rank <= threshold {
							sortedList[rank] = tranco
							fmt.Println(tranco)
							m.Lock()
							count++
							m.Unlock()
						}
					}
				}
			}
			<-sem
		}(url)
	}
	wg.Wait()

	fmt.Printf("\n%sFound %d URLs for the top %d of the tranco-list%s\n", Green, count, threshold, Reset)

	for _, x := range sortedList {
		if x != "" {
			fmt.Println(x)
		}
	}
}

func readLocalFile(path string) []string {

	w, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s%s: %s%s\n", Red, path, err.Error(), Reset)
		os.Exit(2)
	}

	return strings.Split(string(w), "\n")
}

func parseFlags() (string, string, int, int) {
	var pathTranco string
	var pathURLs string
	var threshold int
	var threads int

	flag.StringVar(&pathTranco, "tranco", "", "path to the tranco-list")
	flag.StringVar(&pathTranco, "t", "", "short for tranco")
	flag.StringVar(&pathURLs, "urls", "", "path to the list with URLs")
	flag.StringVar(&pathURLs, "u", "", "short for urls")
	flag.IntVar(&threshold, "threshold", 1000, "Ranking threshold for which URLs will be printed. Default is 1000")
	flag.IntVar(&threads, "threads", 50, "Threads to use. Default is 50")

	flag.Parse()

	return pathTranco, pathURLs, threshold, threads
}
