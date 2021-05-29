package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"storj.io/uplink"
	// "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// // Compile templates on start of the application
// var templates = template.Must(template.ParseFiles("public/upload.html"))

// // Display the named template
// func display(w http.ResponseWriter, page string, data interface{}) {
// 	templates.ExecuteTemplate(w, page+".html", data)
// }

func displayAll(w http.ResponseWriter, r *http.Request) ([]int, error) { // (int[] , error)

	ctx := context.Background()

	var sliceOfObj []int = make([]int, 0)

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
		return nil, fmt.Errorf("could not parse access grant: %v", err)
	}

	// Creates a project using our access
	project, err := uplink.OpenProject(ctx, access)
	if err != nil {
		return nil, fmt.Errorf("could not open project: %v", err)
	}
	defer project.Close()

	// Creates the bucketName and objectKey variables to be used later
	var bucketName string = "bucket1"

	// Ensure the desired Bucket within the Project is created.
	_, err = project.EnsureBucket(ctx, bucketName)
	if err != nil {
		return nil, fmt.Errorf("could not ensure bucket: %v", err)
	}

	// Gets the list of objects currently in Storj and puts them into a list
	objects := project.ListObjects(ctx, "bucket1", nil)
	for objects.Next() {
		item := objects.Item()
		k, e := strconv.Atoi(item.Key)
		if e != nil {
			return nil, e
		}
		fmt.Println(item.IsPrefix, item.Key)
		sliceOfObj = append(sliceOfObj, k)
	}
	if err := objects.Err(); err != nil {
		return nil, err
	}

	fmt.Println(reflect.TypeOf(sliceOfObj))

	return sliceOfObj, nil
}

// UploadData uploads the data to objectKey in bucketName, using accessGrant.
func UploadData(ctx context.Context, data []byte) error {
	//func UploadData(ctx context.Context, data *os.File) error {

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
		return fmt.Errorf("could not parse access grant: %v", err)
	}

	// Creates a project using our access
	project, err := uplink.OpenProject(ctx, access)
	if err != nil {
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
		return fmt.Errorf("could not ensure bucket: %v", err)
	}

	// Intitiate the upload of our Object to the specified bucket and key.
	upload, err := project.UploadObject(ctx, bucketName, objectKey, nil)
	if err != nil {
		return fmt.Errorf("could not initiate upload: %v", err)
	}

	// Copy the data to the upload.
	buf := bytes.NewBuffer(data)
	_, err = io.Copy(upload, buf)
	if err != nil {
		_ = upload.Abort()
		return fmt.Errorf("could not upload data: %v", err)
	}

	// Commit the uploaded object.
	err = upload.Commit()
	if err != nil {
		return fmt.Errorf("could not commit uploaded object: %v", err)
	}

	// Test to see what is uploaded in our bucket
	// objects := project.ListObjects(ctx, "bucket1", nil)
	// for objects.Next() {
	// 	item := objects.Item()
	// 	fmt.Println(item.IsPrefix, item.Key)
	// }
	// if err := objects.Err(); err != nil {
	// 	return err
	// }

	// /////////////////////////// Tests if the image uploaded is the correct image

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

	fmt.Println(receivedContents) //Shows the []byte of the picture

	// Check that the downloaded data is the same as the uploaded data.
	if !bytes.Equal(receivedContents, data) {
		return fmt.Errorf("got different object back: %q != %q", data, receivedContents)
	}

	// // Tests to see if the file is legit
	// f, err := os.OpenFile("./omgItWorks.png", os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer f.Close()

	// io.Copy(f, download)

	// c, err := os.Create("omgItWorks.png")
	// defer c.Close()
	// if err != nil {
	// 	// http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	// return
	// }

	// // Copy the uploaded file to the created file on the filesystem
	// if _, err := io.Copy(c, data); err != nil {
	// 	// http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	// return
	// }

	return nil
}

func uploadFile(w http.ResponseWriter, r *http.Request) {

	// Maximum upload of 10 MB files
	// r.ParseMultipartForm(10 << 20)

	// // Get handler for filename, size and headers
	// file, handler, err := r.FormFile("myFile") //r.Body
	// if err != nil {
	// 	fmt.Println("Error Retrieving the File")
	// 	fmt.Println(err)
	// 	return
	// }

	// defer file.Close()
	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	// // Create file
	// dst, err := os.Create(handler.Filename)
	// defer dst.Close()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// // Copy the uploaded file to the created file on the filesystem
	// if _, err := io.Copy(dst, file); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// Converts *os.File into []byte
	dst := r.Body
	defer dst.Close()
	// _, _ = dst.Seek(0, 0)
	b, err := ioutil.ReadAll(dst)
	if err != nil {
		log.Fatal(err)
	} //000000000

	// Calls UploadData, which takes the given file and uploads it onto Storj
	err = UploadData(context.Background(), b)
	//err = UploadData(context.Background(), dst)
	if err != nil {
		log.Fatal(err)
		return
	}

	// // This code tests if the copy works
	// f, err := os.Create("omg2.png")
	// defer f.Close()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// if _, err := io.Copy(f, file); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	fmt.Fprintf(w, "Successfully Uploaded File\n")
	return
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
		// display(w, "upload", nil)
	case "POST":
		uploadFile(w, r)
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
		// display(w, "upload", nil)
		displayAll(w, r)
	case "POST":
		// downloadFile(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//     w.Header().Set("Content-Type", "application/json")
	//     w.Write([]byte("{\"hello\": \"world\"}"))
	// })

	// just a test page
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Upload route
	mux.HandleFunc("/upload", uploadHandler)
	mux.HandleFunc("/download", downloadHandler)
	handler := cors.Default().Handler(mux)

	fmt.Printf("Starting server at port 8080\n")
	// fmt.Printf(poof)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}

	// //Listen on port 8080
	// http.ListenAndServe(":8080", nil)
}
