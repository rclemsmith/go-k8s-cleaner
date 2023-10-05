package utils

import (
	"example/user/cleaner/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"github.com/buger/jsonparser"
)

func DeleteDeployments(wg *sync.WaitGroup, channel <-chan map[string]string, apiClient models.ApiClient) {
	for deploys := range channel {
		fmt.Println(deploys)
		go GetDeploymentInfo(wg, deploys["namespace"], deploys["deploymentName"], apiClient)
	}
	wg.Done()
}

func GetDeploymentInfo(wg *sync.WaitGroup, namespace string, deployment string, apiClient models.ApiClient) {
	wg.Add(1)
	client := http.Client{}
	url := fmt.Sprintf("%s/apis/apps/v1/namespaces/%s/deployments/%s", apiClient.Url, namespace, deployment)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+apiClient.Token)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	kind, err := jsonparser.GetString(body, "metadata","namespace")
	fmt.Println(kind)
	wg.Done()
}
