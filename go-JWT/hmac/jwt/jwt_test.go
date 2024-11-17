package jwt

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	j := GetJWT()
	str, err := j.Encode(CustomClaimsUser{
		UserId:   1,
		Username: "jock",
	})
	fmt.Printf("%v,%v\n", str, err)
}

func TestDecode(t *testing.T) {
	str := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDdXN0b21DbGFpbXMiOnsidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJqb2NrIn0sImlzcyI6InByb2plY3QiLCJzdWIiOiJqb2NrIiwiYXVkIjpbImFsbCJdLCJleHAiOjE3MzE4NDc2NTgsImlhdCI6MTczMTg0NzYzOCwianRpIjoiNjJhNWNkMzUtN2JiMy00YjllLWI1MGYtNzQ1ZmRhZDIzYzEwIn0.3Bw_6jOd1oNsBguY8XNRmflWfpgrn0d4StRsh7iyIo8`
	j, err := GetJWT().Decode(str)
	fmt.Printf("%#v,%v\n", j, err)
}
