package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

var (
	isQuiet  bool
	isDryRun bool
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("error: ")

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flag.BoolVar(&isQuiet, "quiet", false, "do not print anything")
	flag.BoolVar(&isQuiet, "q", false, "alias to \"-quiet\"")
	flag.BoolVar(&isDryRun, "dry-run", false, "do not remove files")
	flag.BoolVar(&isDryRun, "d", false, "alias to \"-dry-run\"")

	flag.Parse()

	if len(flag.Args()) < 1 {
		return nil
	}
	for _, targetPath := range flag.Args() {
		if err := rmdot(targetPath); err != nil {
			return err
		}
	}

	return nil
}

func rmdot(targetPath string) error {
	paths, err := getCandidatePaths(targetPath)

	if err != nil {
		return err
	}
	if isQuiet {
		goto REMOVE_FILES
	}
	for _, path := range paths {
		fmt.Println(path)
	}

REMOVE_FILES:

	if isDryRun {
		return nil
	}

	var wg sync.WaitGroup

	for _, path := range paths {
		wg.Add(1)

		go func(p string) {
			os.RemoveAll(p)

			wg.Done()
		}(path)
	}

	wg.Wait()

	return nil
}

func getCandidatePaths(rootPath string) ([]string, error) {
	paths := []string{}

	if err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		paths = append(paths, path)

		return err
	}); err != nil {
		return nil, err
	}

	found := map[string]struct{}{}

	// Skip first path because `paths[0] == rootPath`.
	for _, path := range paths[1:] {
		if p := normalizeCandidatePath(path); p != "" {
			found[p] = struct{}{}
		}
	}

	results := []string{}

	for k, _ := range found {
		results = append(results, k)
	}

	sort.Strings(results)

	return results, nil
}

func normalizeCandidatePath(path string) string {
	elements := strings.Split(path, string(os.PathSeparator))

	for n, element := range elements {
		if strings.HasPrefix(element, `.`) {
			return strings.Join(elements[:n+1], string(os.PathSeparator))
		}
	}

	return ""
}
