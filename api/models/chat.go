package models

import "time"

type PrivateMessage struct {
	ID           int
	SenderID     int
	ReceiverID   int
	Content      string
	CreationDate time.Time
}

type Conversation struct {
    ConversationID int       
    UserID1        int       
    UserID2        int
    StartedAt      time.Time 
}