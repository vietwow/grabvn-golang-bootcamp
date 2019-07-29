package main

import (
    "fmt"
    "os"
    "path/filepath"
    "bufio"
    "sync"
)

func CountWord(path string, wg *sync.WaitGroup) ([]string, error) {
    // Call Done() using defer as it's be easiest way to guarantee it's called at every exit
    defer wg.Done()

    fmt.Println("=> Inside, processing file : ", path)
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

    fmt.Println("=> Finished file : ", path)
    return words, nil
}

func main() {
    var files []string

    root := "./files"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if info.Mode().IsRegular() {
            // fmt.Println("Listed path : ", path)
            files = append(files, path)
        }
        return nil
    })
    if err != nil {
        panic(err)
    }

    // Create the waitgroup and add the total number of goroutines we're going to use
    total := 2
    var wg sync.WaitGroup
    wg.Add(total)

    for _, file := range files {
        fmt.Println("Processing file : ", file)
        go func() {
            words, err := CountWord(file, &wg)
            if (err != nil) {
                panic(err)
            }

            for _, word := range words {
                fmt.Println(word)
            }
        }()
    }

    // Wait for all goroutines to be finished
    wg.Wait()
    fmt.Println("Finished All")
}
