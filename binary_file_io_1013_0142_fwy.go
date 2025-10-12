// 代码生成时间: 2025-10-13 01:42:28
This tool provides a simple interface for reading and writing binary files using the Gin framework.

Usage:
- POST /write to write binary data to a file.
- GET /read to read binary data from a file.

Error Handling:
- Errors in file operations are properly handled and returned to the client.

Gin Middleware:
- Gin.Recovery() middleware is used to recover from any panics
- Gin.Logger() middleware is used for logging requests.

Best Practices:
- The code follows Go best practices for file operations and error handling.
*/

package main

import (
    "fmt"
    "net/http"
    "os"
    "io"
    "encoding/binary"
    "github.com/gin-gonic/gin"
)

// Define the maximum file size for uploads
const maxFileSize = 1024 * 1024 * 10 // 10MB

func main() {
    router := gin.Default()

    // Recovery middleware recovers from any panics and writes a 500 if there was one
    router.Use(gin.Recovery())
    // Logger middleware logs all requests, like a combined log logger middleware
    router.Use(gin.Logger())

    router.POST("/write", writeBinaryFile)
    router.GET("/read", readBinaryFile)

    router.Run(":8080") // listen and serve on 0.0.0.0:8080
}

// writeBinaryFile handles the POST request to write binary data to a file
func writeBinaryFile(c *gin.Context) {
    file, err := c.GetRawData()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Failed to get raw data",
        })
        return
    }

    filename := "output.bin"
    err = writeBinaryData(file, filename)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "File written successfully",
    })
}

// readBinaryFile handles the GET request to read binary data from a file
func readBinaryFile(c *gin.Context) {
    filename := "output.bin"
    data, err := readBinaryData(filename)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.Data(http.StatusOK, "application/octet-stream", data)
}

// writeBinaryData writes binary data to a file
func writeBinaryData(data []byte, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("failed to create file: %w", err)
    }
    defer file.Close()

    _, err = file.Write(data)
    if err != nil {
        return fmt.Errorf("failed to write data: %w", err)
    }
    return nil
}

// readBinaryData reads binary data from a file
func readBinaryData(filename string) ([]byte, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        return nil, fmt.Errorf("failed to read data: %w", err)
    }
    return data, nil
}
