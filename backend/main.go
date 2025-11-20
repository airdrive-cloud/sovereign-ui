package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

// HeadscaleNode represents a machine/node from the Headscale API
type HeadscaleNode struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    GivenName   string `json:"givenName"`
    LastSeen    string `json:"lastSeen"`
    Online      bool   `json:"online"`
    IPAddress   string `json:"ipAddresses"`
    User        struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
    } `json:"user"`
}

// UserNodes groups nodes by user
type UserNodes struct {
    User  string          `json:"user"`
    Nodes []HeadscaleNode `json:"nodes"`
}

var headscaleAPIURL string
var headscaleAPIKey string

func init() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using environment variables")
    }

    headscaleAPIURL = os.Getenv("HEADSCALE_API_URL")
    headscaleAPIKey = os.Getenv("HEADSCALE_API_KEY")

    if headscaleAPIURL == "" || headscaleAPIKey == "" {
        log.Fatal("HEADSCALE_API_URL and HEADSCALE_API_KEY must be set")
    }
}

func getNodesFromHeadscale() ([]HeadscaleNode, error) {
    client := &http.Client{Timeout: 10 * time.Second}
    req, err := http.NewRequest("GET", headscaleAPIURL+"/api/v1/machine", nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Authorization", "Bearer "+headscaleAPIKey)
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return nil, fmt.Errorf("Headscale API returned status %d: %s", resp.StatusCode, string(body))
    }

    var nodes []HeadscaleNode
    if err := json.NewDecoder(resp.Body).Decode(&nodes); err != nil {
        return nil, err
    }

    return nodes, nil
}

func main() {
    r := gin.Default()

    // CORS middleware for frontend
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:5173", "http://localhost"} // Allow Vite dev server
    config.AllowMethods = []string{"GET"}
    r.Use(cors.New(config))

    r.GET("/api/v1/nodes", func(c *gin.Context) {
        nodes, err := getNodesFromHeadscale()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        userMap := make(map[string][]HeadscaleNode)
        for _, node := range nodes {
            userMap[node.User.Name] = append(userMap[node.User.Name], node)
        }

        var userNodesList []UserNodes
        for user, nodes := range userMap {
            userNodesList = append(userNodesList, UserNodes{User: user, Nodes: nodes})
        }

        c.JSON(http.StatusOK, userNodesList)
    })

    r.Run(":8081") // Backend server on port 8081
}