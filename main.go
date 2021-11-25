package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
)

func main() {
	var files []string

	replacers_regex := make(map[string]regexp.Regexp)

	if x, err := regexp.Compile("[aA]met"); err != nil {
		panic(err)
	} else {
		replacers_regex["AMET_REPLACE"] = *x
	}

	if x, err := regexp.Compile("[eE]nim"); err != nil {
		panic(err)
	} else {
		replacers_regex["ENIM_REPLACE"] = *x
	}

	srcPath := os.Args[1]
	targetPath := os.Args[2]

	os.RemoveAll(targetPath)

	err := filepath.Walk(srcPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		info, err = os.Stat(path)

		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	logOnErr(err)
	status := make(chan int)

	start := time.Now()
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)

		go func(file string, targetPath string, replacers_regex map[string]regexp.Regexp, wg *sync.WaitGroup, status chan int) {
			anon_file(file, targetPath, replacers_regex, wg)
			status <- 1
		}(file, targetPath, replacers_regex, &wg, status)
	}

	doneCounter := 0

	go func() {
		wg.Wait()
		close(status)
	}()

	bar := progressbar.NewOptions(len(files),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan][Working...][reset] Anonymizing files..."),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionShowCount(),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[red]>[reset]",
			SaucerPadding: "[blue]-[reset]",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	for i := range status {
		doneCounter += i
		// avg := float32(doneCounter) / float32(len(files)) * 100
		// fmt.Printf("\rDone %.2f%% (%d/%d)", avg, doneCounter, len(files))
		bar.Add(1)
	}

	fmt.Printf("\nTook %s \n", time.Since(start))

}

func anon_file(srcPath string, targetPath string, replacers map[string]regexp.Regexp, wg *sync.WaitGroup) {

	file, err := os.Open(srcPath)
	logOnErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		content := scanner.Text()

		for replacer, regex := range replacers {

			modify(&content, replacer, regex)

		}
		appendToFile(content, srcPath, targetPath)
	}
	wg.Done()
}

func modify(content *string, replacer string, regex regexp.Regexp) {
	contentBytes := []byte(*content)
	replacerBytes := []byte(replacer)
	result := string(regex.ReplaceAll(contentBytes, replacerBytes))

	*content = result
}

func appendToFile(content string, srcPath string, targetPath string) {

	srcSplit := strings.Split(srcPath, string(os.PathSeparator))
	newPathSplit := []string{targetPath}
	newPathSplit = append(newPathSplit, srcSplit[1:]...)
	newPathFolder := strings.Join(newPathSplit[:len(newPathSplit)-1], string(os.PathSeparator))

	newPath := strings.Join(newPathSplit, string(os.PathSeparator))

	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		os.MkdirAll(newPathFolder, 0755) // Create your file
	}

	info, err := os.Stat(srcPath)
	logOnErr(err)

	mode := info.Mode()

	f, err := os.OpenFile(newPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, mode)
	logOnErr(err)

	defer f.Close()

	if _, err = f.WriteString(content + "\n"); err != nil {
		panic(err)
	}

}

func logOnErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
