package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "https://api.xiabb.chat/chatapi/marketing/signin"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}
	Authorization := os.Getenv("Authorization")
	if Authorization == "" {
		fmt.Println("Authorization不存在，请检查是否添加")
		return
	}
	req.Header.Set("Authorization", Authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	fmt.Println("Response Body:", string(body))
}
