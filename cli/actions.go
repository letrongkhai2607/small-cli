package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"io"
	"os"
	"net/http"
	"log"
	"net/url"
	"net"
	"bytes"
	"mime/multipart"
	"path/filepath"
	types "example.com/m/v2/types"
	helpers "example.com/m/v2/helpers"
)



func RandowJoke() error {
	url := "https://icanhazdadjoke.com/"
	responseBytes := GetJokeData(url)
	joke := types.Joke{}

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		helpers.HandleError(err)
	}

	fmt.Println(string(joke.Joke))
	return nil
}
func GetJokeData(baseAPI  string) []byte {
	request, err := http.NewRequest(
		http.MethodGet, 
		baseAPI,        
		nil,            
	)

	if err != nil {
		helpers.HandleError(err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjokes CLI")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		helpers.HandleError(err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		helpers.HandleError(err)
	}
	return responseBytes
}

func SearchByImageFile(imagePath string) error {
	baseURL := "https://api.trace.moe/search?anilistInfo"
	/* 
		1: Checking image file is exist or not
		2: Open image file path
	*/
	
	// Step 1
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		if err != nil {
			log.Fatal("Invalid file path")
		}
	}
	// Step 2
	imageFile, err := os.Open(imagePath)
	helpers.HandleError(err)

	payload := &bytes.Buffer{}

	writer := multipart.NewWriter(payload)

	part , _ := writer.CreateFormFile("image", filepath.Base(imagePath))


	_, err = io.Copy(part, imageFile)
	helpers.HandleError(err)

	err = writer.Close()
	helpers.HandleError(err)

	// FormDataContentType returns the Content-Type 
	// Maybe work if we want to post a image to API
	response, err := http.Post(baseURL, writer.FormDataContentType(), payload)
	helpers.HandleError(err)
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	helpers.HandleError(err)

	var anime types.Response
	json.Unmarshal(responseBody, &anime)

	fmt.Println("🌸 Title Native:", anime.Result[0].Anilist.Title.Native)
	fmt.Println("🗻 Title Romaji:", anime.Result[0].Anilist.Title.Romaji)
	fmt.Println("🗽 Title English:", anime.Result[0].Anilist.Title.English)
	fmt.Print("🍓 Is Adult:" , helpers.IsAdult(anime.Result[0].Anilist.IsAdult))
	return nil
}

func SearchByImageLink(imageLink string) error {
	
	baseURL := "https://api.trace.moe/search?anilistInfo&url="
	
	// Start encode URI
	_, err := url.ParseRequestURI(imageLink)
	if err != nil {
		log.Fatal("Invalid url")
	}
	// End encode URI

	response, err := http.Get(baseURL+imageLink)
	helpers.HandleError(err)
	defer response.Body.Close()
	
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		helpers.HandleError(err)
	}

	var anime types.Response

	if err := json.Unmarshal(responseBody, &anime); err != nil {
		helpers.HandleError(err)
	}
	fmt.Println("🌸 Title Native:", anime.Result[0].Anilist.Title.Native)
	fmt.Println("🗻 Title Romaji:", anime.Result[0].Anilist.Title.Romaji)
	fmt.Println("🗽 Title English:", anime.Result[0].Anilist.Title.English)
	fmt.Print("🍓 Is Adult:" , helpers.IsAdult(anime.Result[0].Anilist.IsAdult))
	return nil
}

func getIPAdress() {
	host, _ := os.Hostname()
	address, _ := net.LookupIP(host)
	for _, addr := range address {
		if IPv4 := addr.To4(); IPv4 != nil {
			fmt.Println("IPv4: " , IPv4)
		}   
	}
}