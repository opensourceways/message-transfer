package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	UserData []ForumUser `json:"users"`
}

type ForumUser struct {
	UserName string `json:"username"`
}

func GetForumUserName(userId int) string {
	url := fmt.Sprintf("https://discourse.test.osinfra.cn/user-cards.json?user_ids=%d", userId)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close() // 确保在最后关闭响应体

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", resp.Status)
	}

	var response Data
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	for _, user := range response.UserData {
		return user.UserName
	}
	return ""
}
