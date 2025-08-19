package passcheck

import (
	//"encoding/json"
	//"log"
	"net/http"

	//"github.com/NIA-VEnter/my-pgx---echo-test/tree/main/unitedbackendprgproject/connect/conn"
	"github.com/jackc/pgx/v5"
	//conn "main.go/connect"
)

type UserData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type ResponseData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func checkPassword(conn *pgx.Conn, w http.ResponseWriter, r *http.Request) {

}
