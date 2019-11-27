package main

import (
	"fmt"//is a core library package that contains functionalities related to formatting and printing
    "html/template" // this implements data-driven templates for generating HTML output safe against code injection.
	"io"// this  provides basic interfaces to I/O primitives
	"net/http"//this provides HTTP client and server implementations.
	"os"//this provides a platform-independent interface to operating system functionality
)
//  uplaod handling part and this  perform all the processing and simply return a response to the client.
func upload(w http.ResponseWriter, r *http.Request) {
    // GET displays the upload form.
	if r.Method == "GET" {

		t, _ := template.ParseFiles("upload.html")

		t.Execute(w, nil)
		
		// this 'POST' take the uploaded file(s) and saves it to disk.
	} else if r.Method == "POST" {
		//Get a file handle, store the file
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		// close the file variable
		defer file.Close()
		
		//Create uploaded destination file
		f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		//define error message if file path not found
		if err != nil {
			fmt.Println(err)
			return
		}
		//close f variable
		defer f.Close()
		
		//copy the uploaded file to the destination file
		io.Copy(f, file)
		
		fmt.Fprintf(w, "Successfully Uploaded File\n")

	} else {
		fmt.Println("Unknown HTTP " + r.Method + "  Method")
	}
}
//creating maing fuction
func main() {
	//static file handler.
	http.HandleFunc("/upload", upload)
	//Listen on port 9090
	http.ListenAndServe(":9090", nil) 
}
