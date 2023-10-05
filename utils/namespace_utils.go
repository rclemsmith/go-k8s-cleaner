package utils

import (
	"encoding/json"
	"example/user/cleaner/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"github.com/buger/jsonparser"
)

func GetAllNamespaces(apiClient models.ApiClient) []string{

	client := http.Client{}
	req, err := http.NewRequest("GET", apiClient.Url+"/api/v1/namespaces", nil)
	req.Header.Set("Authorization", "Bearer " + apiClient.Token)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var namespaces []string

	jsonparser.ArrayEach(body,func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		namespace, err := jsonparser.GetString(value,"metadata","name")
		if err != nil{
			fmt.Println(err)
		}
		if strings.Contains(namespace, "dev-") {
			namespaces = append(namespaces, namespace)
		} else if strings.Contains(namespace, "xblox-dev-") {
			namespaces = append(namespaces, namespace)
		}
	},"items")

	return namespaces
}

func GetDeploymentsInANamespace(wg *sync.WaitGroup, channel chan<- map[string]string, namespace string, apiClient models.ApiClient) {
	wg.Add(1)
	client := http.Client{}
	req, err := http.NewRequest("GET", apiClient.Url+"/apis/apps/v1/namespaces/"+namespace+"/deployments", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", "Bearer " + apiClient.Token)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var response map[string]interface{}
	json.Unmarshal([]byte(body), &response)
	deploys := response["items"].([]interface{})
	for _, deploy := range deploys {
		deploy1 := deploy.(map[string]interface{})
		metadata := deploy1["metadata"].(map[string]interface{})
		data := map[string]string{"namespace": metadata["namespace"].(string), "deploymentName": metadata["name"].(string)}
		channel <- data
	}
	// close(channel)
	wg.Done()
}
