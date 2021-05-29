package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
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

	return sliceOfObj, nil
}

// UploadData uploads the data to objectKey in bucketName, using accessGrant.
func UploadData(ctx context.Context, data []byte, dataType string) error {
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

	// print("cheese rice")

	// // Tests to see if the file is legit
	// f, err := os.OpenFile("./omgItWorks.png", os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer f.Close()

	// io.Copy(f, download)

	if dataType == ".txt" {
		buffer := bytes.NewBuffer(receivedContents)

		fo, _ := os.Create("output" + dataType)
		// if _, err := fo.Write(buffer); err != nil {
		// 	panic(err)
		// }
		if _, err := io.Copy(fo, buffer); err != nil {
			panic(err)
		}
	}

	if dataType == ".jpeg" {
		img, _, err := image.Decode(bytes.NewReader(receivedContents))
		if err != nil {
			log.Fatalln(err)
		}

		out, _ := os.Create("./img.jpeg")
		defer out.Close()

		var opts jpeg.Options
		opts.Quality = 1

		err = jpeg.Encode(out, img, &opts)

		if err != nil {
			log.Println(err)
		}
	}

	//permissions := 0644
	// err = ioutil.WriteFile("yer.txt", receivedContents[40:], 0644)
	// if err != nil {
	// 	return fmt.Errorf("could not write data: %v", err)
	// }

	// f, err := os.Open("inputimage.jpg")
	// if err != nil {
	// 	// Handle error
	// }
	// defer f.Close()

	// img, fmtName, err := image.Decode(download)
	// if err != nil {
	// 	// Handle error
	// 	print("oof 1")
	// 	return err
	// }

	// print(fmtName)

	// f, err := os.Create("outimage.png")
	// if err != nil {
	// 	// Handle error
	// 	print("oof 2")
	// 	return err
	// }
	// defer f.Close()

	// // Encode to `PNG` with `DefaultCompression` level
	// // then save to file
	// err = png.Encode(f, img)
	// if err != nil {
	// 	// Handle error
	// 	print("oof 3")
	// 	return err
	// }

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

	// bDst := r.Body

	// defer bDst.Close()
	// _, _ = bDst.Seek(0, 0)
	// b, err := ioutil.ReadAll(bDst)
	// if err != nil {
	// 	log.Fatal(err)
	// } //000000000

	// buf := bytes.NewBuffer(b)
	// print(buf)
	print("hi")

	reader, err := r.MultipartReader()
	if err != nil {
		print("yer 1")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// parse text field
	text := make([]byte, 2048)

	p, err := reader.NextPart()
	// one more field to parse, EOF is considered as failure here
	if err != nil {
		print("yer 2")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// if p.FormName() != "text_field" {
	// 	http.Error(w, "text_field is expected", http.StatusBadRequest)
	// 	return
	// }

	_, err = p.Read(text)
	if err != nil && err != io.EOF {
		print("yer 3")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	print(text)

	//dst := r.Body
	var fileType string = ".txt" //////////////Should be variable based on the type of file passed by user
	dst := text
	// defer dst.Close()
	// // _, _ = dst.Seek(0, 0)
	// b, err := ioutil.ReadAll(dst)
	// if err != nil {
	// 	log.Fatal(err)
	// } //000000000

	// Calls UploadData, which takes the given file and uploads it onto Storj
	err = UploadData(context.Background(), dst, fileType)
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

	// if _, err := io.Copy(f, dst); err != nil {
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
