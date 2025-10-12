// 代码生成时间: 2025-10-12 20:23:41
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Middleware that checks for admin access
func AdminAuthMiddleware(c *gin.Context) {
    // Assuming the username is passed in the query string
    username := c.Query("username")

    // You can replace this with your actual admin check logic
    // Here it is hardcoded for demonstration purposes
    adminUsername := "admin"

    if username != adminUsername {
        // If the user is not admin
        c.JSON(http.StatusForbidden, gin.H{
            "error": "Admin access required",
        })
        c.Abort() // Stop processing chain
        return
    }

    // Continue request to next middleware/handler
    c.Next()
}

func main() {
    r := gin.Default()

    // Register the AdminAuthMiddleware
    r.Use(AdminAuthMiddleware)

    // Define the route for admin only access
    r.GET("/admin", func(c *gin.Context) {
        // This handler will only be executed if AdminAuthMiddleware allows the request
        c.JSON(http.StatusOK, gin.H{
            "message": "Welcome admin!",
        })
    })

    // Start the server
    r.Run() // listen and serve on 0.0.0.0:8080
}
