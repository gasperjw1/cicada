package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"storj.io/uplink"

	"github.com/joho/godotenv"
)

// Compile templates on start of the application
var templates = template.Must(template.ParseFiles("public/upload.html"))

// Display the named template
func display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}

// UploadData uploads the data to objectKey in bucketName, using accessGrant.
func UploadData(ctx context.Context, data []byte) error {

	print("Now we here\n")

	//Gets access grant stored in .env
	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessGrant := envs["STORJACCESSGRANT"]

	// Parse the Access Grant.
	access, err := uplink.ParseAccess(accessGrant)
	if err != nil {
		print("d")
		return fmt.Errorf("could not parse access grant: %v", err)
	}

	print("Now we here\n")

	// Creates a project using our access
	project, err := uplink.OpenProject(ctx, access)
	if err != nil {
		print("f")
		return fmt.Errorf("could not open project: %v", err)
	}
	defer project.Close()

	// Creates the bucketName and objectKey variables to be used later
	var bucketName string = "bucket1"
	var objectKey string = time.Now().Format("2006-01-02 15:04:05")
	objectKey = strings.Replace(objectKey, ":", "-", -1)
	objectKey = strings.Replace(objectKey, " ", "_", -1)

	// Ensure the desired Bucket within the Project is created.
	_, err = project.EnsureBucket(ctx, bucketName)
	if err != nil {
		print("y")
		return fmt.Errorf("could not ensure bucket: %v", err)
	}

	print("Now we here\n")

	// Intitiate the upload of our Object to the specified bucket and key.
	upload, err := project.UploadObject(ctx, bucketName, objectKey, nil)
	if err != nil {
		print("i")
		return fmt.Errorf("could not initiate upload: %v", err)
	}

	// Copy the data to the upload.
	buf := bytes.NewBuffer(data)
	_, err = io.Copy(upload, buf)
	if err != nil {
		print("p")
		_ = upload.Abort()
		return fmt.Errorf("could not upload data: %v", err)
	}

	print("Now we here\n")

	// Commit the uploaded object.
	err = upload.Commit()
	if err != nil {
		print("u")
		return fmt.Errorf("could not commit uploaded object: %v", err)
	}

	///////////////////////////// Tests if the image uploaded is the correct image

	// Initiate a download of the same object again
	download, err := project.DownloadObject(ctx, bucketName, objectKey, nil)
	if err != nil {
		return fmt.Errorf("could not open object: %v", err)
	}
	defer download.Close()

	// Read everything from the download stream
	receivedContents, err := ioutil.ReadAll(download)
	if err != nil {
		return fmt.Errorf("could not read data: %v", err)
	}

	print(receivedContents) //Shows that the []byte contains 0 bytes (512 capacity)

	// Check that the downloaded data is the same as the uploaded data.
	if !bytes.Equal(receivedContents, data) {
		return fmt.Errorf("got different object back: %q != %q", data, receivedContents)
	}

	return nil
}

func uploadFile(w http.ResponseWriter, r *http.Request) {

	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create(handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	print(dst)

	// Converts *os.File into []byte
	b, err := ioutil.ReadAll(dst)
	if err != nil {
		log.Fatal(err)
	}

	// Calls UploadData, which takes the given file and uploads it onto Storj
	err = UploadData(context.Background(), b)
	if err != nil {
		log.Fatal(err)
		return
	}

	// This code tests if the copy works
	// f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer f.Close()

	// io.Copy(f, file)

	fmt.Fprintf(w, "Successfully Uploaded File\n")
	return
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		display(w, "upload", nil)
	case "POST":
		uploadFile(w, r)
	}
}

func main() {

	// just a test page
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Upload route
	http.HandleFunc("/upload", uploadHandler)

	fmt.Printf("Starting server at port 8080\n")
	// fmt.Printf(poof)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	// //Listen on port 8080
	// http.ListenAndServe(":8080", nil)
}
