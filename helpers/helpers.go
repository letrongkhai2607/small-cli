package heplers

import (
	"log"
	"strings"
	"net/url"
	"net/http"
	"os"
	"io"
	"fmt"
)


// HandleError is a reusable error checking function
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func IsAdult(isAdult bool) string{
	if isAdult == false{
		return " False"
	}
	return " True"
}

func BuildFileName(fullUrlFile string) string {
	dir := "./dowloader/"
	var ( fileName string )
    fileUrl, err := url.Parse(fullUrlFile)
    HandleError(err)

    path := fileUrl.Path
    segments := strings.Split(path, "/")

    fileName = segments[len(segments)-1]
	result := dir + fileName
	return result
}

func CreateFile(fileName string) *os.File {
    file, err := os.Create(fileName)

    HandleError(err)
    return file
}
func HttpClient() *http.Client {
    client := http.Client{
        CheckRedirect: func(r *http.Request, via []*http.Request) error {
            r.URL.Opaque = r.URL.Path
            return nil
        },
    }

    return &client
}

func PutFile(file *os.File, client *http.Client,  fullUrlFile string) {
    resp, err := client.Get(fullUrlFile)

    HandleError(err)

    defer resp.Body.Close()

    size, err := io.Copy(file, resp.Body)

    defer file.Close()

    HandleError(err)
	fmt.Println("Just Downloaded a file with size", size)
}