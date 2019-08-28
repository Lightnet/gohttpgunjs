package main

import (
	"fmt"
	"gogun"
	//"log"
	"net/http"
	"io/ioutil"
)

func readfiledata(){
	// read in the contents of the localfile.data
	data, err := ioutil.ReadFile("localfile.data")
	// if our program was unable to read the file
	// print out the reason why it can't
	if err != nil {
		fmt.Println(err)
	}
 
	// if it was successful in reading the file then
	// print out the contents as a string
	fmt.Print(string(data))
}

func writefiledata(){
	mydata := []byte("all my data I want to write to a file")
	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("localfile.data", mydata, 0777)
	// handle this error
	if err != nil {
  		// print it out
  		fmt.Println(err)
	}
}

func main() {
	//var Gun = gun.Init()
	//var Gun = gun.
	//fmt.Println(gun.Add(2, 1))
	//log.Println(gun.Add(2, 1))

	fmt.Println("Testing...!")

	var Gun gogun.GunI = gogun.Gun{}
	Gun.Test()

	//readfiledata();
	writefiledata();

	//http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to my website!")
	//})

	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
	//log.Println("Listening...")
	//http.ListenAndServe(":8080", nil)

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	//http.HandleFunc("/", HelloServer)
	//http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
