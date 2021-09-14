package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and returns the size of each found file and the number
// of the files
func walkDir(dir string) (numFiles int64, size int64) {
	time.Sleep(100 * time.Millisecond)
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			nf, fs := walkDir(subDir)
			numFiles += nf
			size += fs
		} else {
			numFiles++
			size += entry.Size()
		}
	}

	return
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dul: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	flag.Parse()

	// Determine the initial directories
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	start := time.Now()

	var nfiles, nbytes int64
	for _, root := range roots {
		nf, nb := walkDir(root)
		nfiles += nf
		nbytes += nb
	}

	fmt.Println("Total time taken: ", time.Since(start))
	printDiskUsage(nfiles, nbytes)
}
