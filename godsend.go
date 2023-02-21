package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

// Get preferred outbound ip of this machine.
// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func getLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	// get the file path
	file := os.Args[1]

	// print QR code to terminal
	content := fmt.Sprintf("http://%s:8080", getLocalIP())
	qr, err := qrcode.New(content, qrcode.Highest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(qr.ToString(false))

	// serve file
	fileHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, file)
	}
	http.HandleFunc("/", fileHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
