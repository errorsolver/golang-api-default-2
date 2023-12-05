package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "http://localhost:8888/videos"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Errorf("%s", err)
		return
	}
	req.Header.Add("Authorization", "Basic YWRtaW46QWRtaW4=")

	res, err := client.Do(req)
	if err != nil {
		fmt.Errorf("%s", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("%s", err)
		return
	}
	fmt.Println(string(body))
}
