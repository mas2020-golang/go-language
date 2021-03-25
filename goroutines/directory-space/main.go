package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

/*
	This application uses some of the previous concepts to calculate the size of a directory or a bunch of directories
	given as input.
*/
func main() {
	var verbose = flag.Bool("v", false, "show verbose progress messages")
	var start, end time.Time
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	/*
		Traverse the file tree with the walkDir func in a separate go routine. fileSizes is an
		unbuffered channel that is incremented on the way. When the recursion is done the channel
		is closed.
	*/
	fileSizes := make(chan int64)
	start = time.Now()
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(1000 * time.Millisecond)
	}

	// Print the results.
	var nfiles, nbytes int64
loop:
	for {
		select {
		// the fileSizes is read by the main go routine until it is closed. It increments the total of
		// bytes and the number of file. The for ends (thanks to the range statement) when the channel is closed.
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			// print the number of files and the space along the way
			printDiskUsage(nfiles, nbytes)
		}
	}
	end = time.Now()
	// print the number of final files and final space
	printDiskUsage(nfiles, nbytes)
	execTime := end.Sub(start)
	fmt.Printf("execution time %d ms\n", execTime.Milliseconds())
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB (%d bytes)\n", nfiles, float64(nbytes)/1e9, nbytes)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
