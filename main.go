package main

import (
	"example/user/cleaner/models"
	"example/user/cleaner/utils"
	"sync"
)



func main() {

	channel := make(chan map[string]string)
	wg := &sync.WaitGroup{}
	var apiClient models.ApiClient

	apiClient.Url = "http://localhost:8080"
	apiClient.Token = "eyJhbGciOiJSUzI1NiIsImtpZCI6Imd1Q3FoRXZvUEhRcHp6ZTRzU2JlLWF6Y085Mzd3Zl9oRmlPTEVKYU5KQ3MifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJjbGVhbmVyIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImNsZWFuZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiY2xlYW5lciIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImM0N2I2YjcwLTJjYWYtNGNhMS1iMGY4LTRlNTQ0NzJkN2UxZCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpjbGVhbmVyOmNsZWFuZXIifQ.ESeIMkUMq_hokC41iNbW0xBw1m_pU-5fAkSAvsNUpk35W_sKxa9oOgjMrLpOGZ9rOWvgxBPy0swyeWwGq7qkhuOkLBfznpsFeAb8as-oCP0q_FxPC_g3N1TIkxX56nObBAJZgh75YTIkhP_uMiTbi69jfOwizA1wnYs8-QWYY07nqD7dxsqkQ-MubdPbypQ-9BCydP-Z5wRepq3PS-fOMQvVXZ3XarWxAYBeK_UWZlisRcJksz2GXWtEU8saI-Noa0YaLh_yK0Xx3MoyqRZooxyh_YiG72HoPkP09XLEEt1kHhDm7y8lXKE-OMZ2o7NEVDnOvP4x-P4ZMb_-VbkH3Q"

	go utils.DeleteDeployments(wg,channel,apiClient)
	go utils.DeleteDeployments(wg,channel,apiClient)
	go utils.DeleteDeployments(wg,channel,apiClient)
	go utils.DeleteDeployments(wg,channel,apiClient)


	namespaces := utils.GetAllNamespaces(apiClient)
	for _, namespace := range namespaces {
		go utils.GetDeploymentsInANamespace(wg,channel,namespace,apiClient)
	}

	wg.Wait()

}

