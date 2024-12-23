package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Good struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	PlayTime    time.Time `json:"play_time" datetime:"omitempty"`
	ExecuteTime time.Time `json:"execute_time" datetime:"2006-01-02"`
	CreatedAt   time.Time `json:"created_at" datetime:"omitempty"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t Good) MarshalJSON() ([]byte, error) {
	return marshaljson.MarshalFormat(t)
}

func main() {
	good := Good{ID: 123, Name: "jock", PlayTime: time.Now(), ExecuteTime: time.Now()}
	bytes, _ := json.Marshal(good)
	// {"id":123,"name":"jock","play_time":"2024-03-15 18:23:43","execute_time":"2024-03-15","created_at":"","updated_at":"0000-00-00 00:00:00"}
	fmt.Printf("%s\n", bytes)
}
