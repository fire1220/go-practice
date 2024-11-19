package marshaljson

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Good struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	PlayTime    time.Time `json:"play_time" datetime:"omitempty"`
	ExecuteTime time.Time `json:"execute_time" datetime:"2006-01-02" default:"-"`
	CreatedAt   time.Time `json:"created_at" datetime:"omitempty"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t Good) MarshalJSON() ([]byte, error) {
	return MarshalFormat(t)
}

func TestMarshal(t *testing.T) {
	good := Good{ID: 123, Name: "jock", PlayTime: time.Now(), ExecuteTime: time.Now()}
	bytes, _ := json.Marshal(good)
	// {"id":123,"name":"jock","play_time":"2024-03-15 18:23:43","execute_time":"2024-03-15","created_at":"","updated_at":"0000-00-00 00:00:00"}
	fmt.Printf("%s\n", bytes)
}
