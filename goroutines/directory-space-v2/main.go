package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

/*
	This application uses some of the previous concepts to calculate the size of a directory or a bunch of directories
	given as input.
	This version creates a new go routine for each call to walkDir. Since we do not know the number of calls done to
	walkDir func the code uses a sync.WaitGroup to check the number of calls to walkDir that are still active and
	a closer go routine to close the fileSizes channel when the counter drops to zero.
*/

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

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
	// wait group is a way to know when the walkDir has finished its job. When the n is 0 it is the end. n is incremented
	// by Add() method and decremented by Done() method
	var n sync.WaitGroup
	start = time.Now()
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	// this go routine waits till the n counter is 0, then close the fileSizes channel
	go func() {
		n.Wait()
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
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	// at the end of walkDir n is decremented by one
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	// semaphore to prevent to open too many files (set the channel semaphore at the initial section of the code)
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
