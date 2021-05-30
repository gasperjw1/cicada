package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	// "strconv"
	"strings"
	"time"

	"storj.io/uplink"
	// "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	//"storj.io/storj/lib/uplink"
	//"storj.io/storj/pkg/storj"
)

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
	// objects := project.ListObjects(ctx, "bucket1", nil)
	// for objects.Next() {
	// 	item := objects.Item()
	// 	k, e := strconv.Atoi(item.Key)
	// 	if e != nil {
	// 		return nil, e
	// 	}
	// 	fmt.Println(item.IsPrefix, item.Key)
	// 	sliceOfObj = append(sliceOfObj, k)
	// }
	// if err := objects.Err(); err != nil {
	// 	return nil, err
	// }

	//fmt.Println(reflect.TypeOf(sliceOfObj))

	theJSON, err := json.Marshal(sliceOfObj)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(theJSON)

	return sliceOfObj, nil
}

// UploadData uploads the data to objectKey in bucketName, using accessGrant.
func UploadData(ctx context.Context, data []byte, dataType string) error {

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
	var bucketName string = "bucket2"
	var objectKey string = time.Now().Format("2006-01-02 15:04:05")
	objectKey = strings.Replace(objectKey, ":", "-", -1)
	objectKey = strings.Replace(objectKey, " ", "_", -1)
	objectKey = objectKey + dataType

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

	return nil
}

func uploadFile(w http.ResponseWriter, r *http.Request) {

	file, header, err := r.FormFile("fileName")
	defer file.Close()
	if err != nil {
		return
	}

	print(file)
	print(header)
	fmt.Printf("Uploaded File: %+v\n", header.Filename)
	fmt.Printf("File Size: %+v\n", header.Size)
	fmt.Printf("MIME Header: %+v\n", header.Header)

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	print("yerrr" + strings.Split(header.Filename, ".")[1])

	// Calls UploadData, which takes the given file and uploads it onto Storj
	err = UploadData(context.Background(), b, "."+strings.Split(header.Filename, ".")[1])
	//err = UploadData(context.Background(), dst)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
	return
}

func serveFromStorj(objectKey string) (*os.File, error) {

	ctx := context.Background()

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
	var bucketName string = "bucket2"

	// Ensure the desired Bucket within the Project is created.
	_, err = project.EnsureBucket(ctx, bucketName)
	if err != nil {
		return nil, fmt.Errorf("could not ensure bucket: %v", err)
	}

	// Initiate a download of the same object again
	download, err := project.DownloadObject(ctx, bucketName, objectKey, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open object: %v", err)
	}
	defer download.Close()

	// Read everything from the download stream
	receivedContents, err := ioutil.ReadAll(download)
	if err != nil {
		return nil, fmt.Errorf("could not read data: %v", err)
	}

	fmt.Println(receivedContents) //Shows the []byte of the picture

	buffer := bytes.NewBuffer(receivedContents)

	fo, _ := os.Create(objectKey)

	if _, err := io.Copy(fo, buffer); err != nil {
		panic(err)
	}

	return fo, nil
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
		query := r.URL.Query()
		resp, err := serveFromStorj(query["uuid"][0]) // <----- get file based on file UID
		if err != nil {
			return
		}

		// Read everything from the download stream
		// receivedContents, err := ioutil.ReadAll(resp)
		// if err != nil {
		// 	return
		// }

		// // Write the content to the desired file
		// err = ioutil.WriteFile(query["uuid"][0], receivedContents, 0644)
		// if err != nil {
		// 	return
		// }

		theJSON, err := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(theJSON)
		// downloadFile(w, r)
	default:
		print("wrong method")
	}
}

func displayHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
		displayAll(w, r)
	default:
		print("wrong method")
	}
}

func main() {
	mux := http.NewServeMux()

	// // just a test page
	// fileServer := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileServer)

	// Upload route
	mux.HandleFunc("/upload", uploadHandler)
	mux.HandleFunc("/display", displayHandler)
	mux.HandleFunc("/download", downloadHandler)
	handler := cors.Default().Handler(mux)

	fmt.Printf("Starting server at port 8080\n")
	// fmt.Printf(poof)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
