package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var (
	re = regexp.MustCompile(`\d+`)
)

func dirfiles(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths, nil
}

func filterPrefix(prefix string, files []string) []string {
	ret := make([]string, 0)
	for _, file := range files {
		if strings.HasPrefix(file, prefix) {
			ret = append(ret, file)
		}
	}
	return ret
}

func getNextNum(prefix, suffix string, files []string) string {
	nums := make([]int, 0, len(files))
	for _, file := range files {
		stripped := strings.TrimPrefix(file, prefix)
		stripped = strings.TrimSuffix(stripped, suffix)
		num := re.FindStringSubmatch(stripped)

		if len(num) > 0 {
			n, _ := strconv.Atoi(num[0])
			nums = append(nums, n)
		}
	}

	max := 0
	for _, n := range nums {
		if max < n {
			max = n
		}
	}
	return strconv.Itoa(max + 1)
}

func main() {
	suffix := flag.String("suffix", "", "")
	flag.Parse()

	filename := "incr"
	if flag.NArg() > 0 {
		args := flag.Args()
		filename = args[0]
	}
	filename = strings.TrimSuffix(filename, *suffix)
	files, err := dirfiles(".")
	if err != nil {
		log.Fatal(err)
	}
	filtered := filterPrefix(filename, files)
	fmt.Print(filename + getNextNum(filename, *suffix, filtered) + *suffix)
}
