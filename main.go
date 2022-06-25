package main
import (
	"github.com/KibetBrian/backend/api"
)

func main() {
	server := api.NewServer()
	server.Start("http://localhost:8080/")
}
