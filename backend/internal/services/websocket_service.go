package services

import (
    "sync"
    "github.com/gorilla/websocket"
    "log"
)

type WebSocketService struct {
    clients    map[string]*websocket.Conn
    clientsMux sync.RWMutex
}

var (
    wsUpgrader = websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
        CheckOrigin: func(r *http.Request) bool {
            return true // In production, implement proper origin checking
        },
    }
    wsService = &WebSocketService{
        clients: make(map[string]*websocket.Conn),
    }
)

func (ws *WebSocketService) AddClient(userID string, conn *websocket.Conn) {
    ws.clientsMux.Lock()
    defer ws.clientsMux.Unlock()
    ws.clients[userID] = conn
}

func (ws *WebSocketService) RemoveClient(userID string) {
    ws.clientsMux.Lock()
    defer ws.clientsMux.Unlock()
    delete(ws.clients, userID)
}

func (ws *WebSocketService) BroadcastUpdate(message interface{}) {
    ws.clientsMux.RLock()
    defer ws.clientsMux.RUnlock()

    for userID, client := range ws.clients {
        err := client.WriteJSON(message)
        if err != nil {
            log.Printf("Error broadcasting to client %s: %v", userID, err)
            client.Close()
            ws.RemoveClient(userID)
        }
    }
}