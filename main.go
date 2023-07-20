// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	url := "https://api.xiabb.chat/chatapi/marketing/signin"

// 	req, err := http.NewRequest("POST", url, nil)
// 	if err != nil {
// 		fmt.Println("Failed to create request:", err)
// 		return
// 	}
// 	Authorization := os.Getenv("AUTHORIZATION")
	
// 	if Authorization == "" {
// 		fmt.Println("Authorization不存在，请检查是否添加")
// 		return
// 	}
// 	req.Header.Set("Authorization", Authorization)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Failed to send request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("Response Status:", resp.Status)

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Failed to read response body:", err)
// 		return
// 	}

// 	fmt.Println("Response Body:", string(body))
// }


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "https://api.xiabb.chat/chatapi/marketing/signin"

	authorizations := []string{
		os.Getenv("AUTHORIZATION_1"),
		os.Getenv("AUTHORIZATION_2"),
	}

	for i, authorization := range authorizations {
		username := fmt.Sprintf("用户%d", i+1)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			fmt.Println("Failed to create request:", err)
			return
		}
		req.Header.Add("Authorization", authorization)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Failed to send request:", err)
			return
		}
		defer resp.Body.Close()

		fmt.Printf("Response Status for %s: %s\n", username, resp.Status)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Failed to read response body:", err)
			return
		}

		fmt.Printf("Response Body for %s: %s\n", username, string(body))
	}
}
