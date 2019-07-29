package main

import (
    "fmt"
    "os"
    "path/filepath"
    "bufio"
)

func CountWord(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    scanner.Split(bufio.ScanWords)

    var words []string

    for scanner.Scan() {
        words = append(words,scanner.Text())
    }

    return words, nil
}

func main() {
    var files []string

    root := "/Users/vietwow/Documents/Grab/week-2/assignment/files"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        // fmt.Println(file)
        words, err := CountWord(file)
        if (err != nil) {
            panic(err)
        }

        for _, word := range words {
            fmt.Println(word)
        }
    }
}