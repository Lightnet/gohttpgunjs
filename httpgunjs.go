package main
//https://gowebexamples.com/websockets/
//Packages
import (
	"fmt"
	"strconv"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"gogun"
)
//Web Socket
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}
//Server PORT
var PORT string = strconv.Itoa(8080) 
//main entry point
func main() {
	fmt.Println("init http server...")
	/* //web socket function setup
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
        for {
            // Read message from browser
            msgType, msg, err := conn.ReadMessage()
            if err != nil {
                return
            }
            // Print the message to the console
            fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
            // Write message back to browser
            if err = conn.WriteMessage(msgType, msg); err != nil {
                return
            }
        }
	})
	*/
	//Gun Web Server setup
	http.HandleFunc("/gun", gogun.WSEndpoint);
	//gogun.web(&http) //not added
	//http url handle and folder root
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
	//init server listen
	http.ListenAndServe(":"+PORT, nil)
	log.Println("Listening http://localhost:"+PORT)
	//init gun / test ?
	var Gun gogun.GunI = gogun.Gun{}
	Gun.Test()
	fmt.Println("finish server & gun init...")
}