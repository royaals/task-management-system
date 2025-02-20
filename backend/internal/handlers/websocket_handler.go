package handlers

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "task-management/internal/services"
)

func HandleWebSocket(c *gin.Context) {
    userID := c.GetString("user_id")
    conn, err := services.WsUpgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Printf("Error upgrading to websocket: %v", err)
        return
    }

    // Add client to WebSocket service
    services.WsService.AddClient(userID, conn)

    // Handle WebSocket connection
    go handleWebSocketConnection(userID, conn)
}

func handleWebSocketConnection(userID string, conn *websocket.Conn) {
    defer func() {
        conn.Close()
        services.WsService.RemoveClient(userID)
    }()

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            return
        }

        if messageType == websocket.TextMessage {
            // Handle incoming messages
            handleWebSocketMessage(userID, string(p))
        }
    }
}

func handleWebSocketMessage(userID string, message string) {
    // Implement message handling logic
    log.Printf("Received message from user %s: %s", userID, message)
}