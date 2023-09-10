package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func main() {
    // Name of the zip file to create
    zipFileName := "go-code.zip"

    // Directory containing your Golang code
    sourceDir := "stage1/main.go"

    // Create or overwrite the zip file
    zipFile, err := os.Create(zipFileName)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer zipFile.Close()

    // Create a new zip writer
    zipWriter := zip.NewWriter(zipFile)
    defer zipWriter.Close()

    // Walk through the source directory and add files to the zip archive
    err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Create a new entry for the file in the zip archive
        zipEntry, err := zipWriter.Create(filepath.ToSlash(path))
        if err != nil {
            return err
        }

        // Open the source file
        sourceFile, err := os.Open(path)
        if err != nil {
            return err
        }
        defer sourceFile.Close()

        // Copy the file's contents to the zip entry
        _, err = io.Copy(zipEntry, sourceFile)
        return err
    })

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Zip file created:", zipFileName)
}
