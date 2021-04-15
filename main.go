package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func dirfiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
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

func getNextNum(prefix string, files []string) string {
	re := regexp.MustCompile(`\d+`)

	nums := make([]int, 0, len(files))
	for _, file := range files {
		stripped := strings.TrimPrefix(file, prefix)
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
	n := len(os.Args)
	filename := "incr"
	if n > 1 {
		filename = os.Args[1]
	}
	files, err := dirfiles(".")
	if err != nil {
		log.Fatal(err)
	}
	filtered := filterPrefix(filename, files)
	fmt.Print(filename + getNextNum(filename, filtered))
}
