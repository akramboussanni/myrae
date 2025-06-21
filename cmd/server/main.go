package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/akramboussanni/myrae/internal/api/routes"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	r := routes.SetupRouter()

	log.Println("Port to use?")
	portNum, _ := reader.ReadString('\n')
	port := fmt.Sprint(":", strings.TrimSpace(portNum))

	log.Println(fmt.Sprint("Server will run on", port))
	http.ListenAndServe(port, r)
}
