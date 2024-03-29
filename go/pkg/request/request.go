package request

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Request(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return url, errors.New("Get() response error")
	}

	// close once body is returned
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("couldn't read body")
	}

	data := string(body)

	return data, nil
}


func APIRequest(url string) (string, error) {
	req, err := http.NewRequest("GET", os.Getenv("MORALIS_API_URL") + url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set Headers
	req.Header = http.Header{
		"accept":    []string{"application/json"},
		"x-api-key": []string{os.Getenv("MORALIS_API_KEY")},
	}

	// Make Request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Format Body
	defer resp.Body.Close()
	resBody, _ := ioutil.ReadAll(resp.Body)
	response := string(resBody)

	return response, nil
}