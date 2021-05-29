package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/joho/godotenv"
	// "storj.io/uplink"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {

	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", reqBody)

	// file, handler, err := r.FormFile("myFile")
	// if err != nil {
	// 	fmt.Println("Error Retrieving the File")
	// 	fmt.Println(err)
	// 	return
	// }

	// defer file.Close()
	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	// switch r.Method {
	// case "POST":
	// 	r.ParseMultipartForm(10 << 20)

	// 	// Get handler for filename, size and headers
	// 	file, handler, err := r.FormFile("file") //<- file from form
	// 	if err != nil {
	// 		fmt.Println("Error Retrieving the File")
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	defer file.Close()
	// 	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// 	fmt.Printf("File Size: %+v\n", handler.Size)
	// 	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// 	// Create file
	// 	dst, err := os.Create(handler.Filename)
	// 	defer dst.Close()
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	// Copy the uploaded file to the created file on the filesystem
	// 	if _, err := io.Copy(dst, file); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	fmt.Fprintf(w, "Successfully Uploaded File\n")

	// default:
	// 	fmt.Fprintf(w, "Sorry, only POST metherd is supported.")
	// }

}

// Main
func main() {
	//Creates a project
	// var envs map[string]string
	// envs, err := godotenv.Read(".env")

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// name := envs["STORJACCESSGRANT"]
	// project, err := uplink.OpenProject(ctx, access)
	// if err != nil {
	// 	return err
	// }
	// defer project.Close()

	// just a test page
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// go router
	router := mux.NewRouter()
	router.HandleFunc("/upload-file", UploadFile).Methods("POST")
	//router.Handle("/upload-file", UploadFile)
	// http.HandleFunc("/upload-file", UploadFile) //enctype="multipart/form-data"

	fmt.Printf("Starting server at port 8080\n")
	// fmt.Printf(poof)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
