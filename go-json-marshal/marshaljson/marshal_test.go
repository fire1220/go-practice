package marshaljson

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Good struct {
	ID          int32     `json:"id" default:"456"`
	ValFloat    float64   `json:"val_float" default:"111"`
	Val         uint      `json:"val" default:"-111"`
	ValBool     bool      `json:"val_bool" default:"true"`
	Name        string    `json:"name" default:"123"`
	PlayTime    time.Time `json:"play_time" datetime:"2006-01-02 15:04:05"`
	ExecuteTime time.Time `json:"execute_time" datetime:"2006-01-02" default:"-"`
	CreatedAt   time.Time `json:"created_at" datetime:"2006-01-02 15:04:05" default:"0000-00-00"`
	UpdatedAt   time.Time `json:"updated_at" default:""`
}

func (t Good) MarshalJSON() ([]byte, error) {
	return MarshalFormat(t)
}

func TestMarshal(t *testing.T) {
	good := Good{ID: 0, Name: "", PlayTime: time.Now(), ExecuteTime: time.Now()}
	bytes, _ := json.Marshal(good)
	// {"id":123,"name":"jock","play_time":"2024-03-15 18:23:43","execute_time":"2024-03-15","created_at":"","updated_at":"0000-00-00 00:00:00"}
	fmt.Printf("%s\n", bytes)
}
