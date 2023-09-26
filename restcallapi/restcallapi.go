package restcallapi

import (
	"github.com/SRdev04/api-native-store-online/execption"
	"github.com/SRdev04/api-native-store-online/helper"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RestCallApiSet(url, userId, accessToken string) error {

	requestBody := []byte(fmt.Sprintf(`{"access_token": "%v", "id_user": "%s"}`, accessToken, userId))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(execption.NewNotFound(err.Error()))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Gagal melakukan permintaan:", err)
		panic(execption.NewNotFound(err.Error()))

	}
	req.Header.Set("Content-Type", "application/json")
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
	fmt.Println(resp.Body)
	return nil
}

func RestCallApiChechk(url, accessToken string) (bool, error) {

	requestBody := []byte(fmt.Sprintf(`{"access_token": "%s"}`, accessToken))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	helper.IfError(err)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Gagal melakukan permintaan:", err)
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	defer resp.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	data, _ := result["Data"].(bool)

	if !data {

		fmt.Println("not found result")

		return data, err

	}

	return data, nil
}
