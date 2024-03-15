package main

import (
	"encoding/json"
	"fmt"
	"main/marshal"
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
	return marshal.Marshal(t)
}

func main() {
	good := Good{ID: 123, Name: "jock", PlayTime: time.Now(), ExecuteTime: time.Now()}
	bytes, _ := json.Marshal(good)
	fmt.Printf("%s", bytes)
}
