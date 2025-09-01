package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var accessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IkNOdjBPSTNSd3FsSEZFVm5hb01Bc2hDSDJYRSIsImtpZCI6IkNOdjBPSTNSd3FsSEZFVm5hb01Bc2hDSDJYRSJ9.eyJhdWQiOiJodHRwczovL2FuYWx5c2lzLndpbmRvd3MubmV0L3Bvd2VyYmkvYXBpIiwiaXNzIjoiaHR0cHM6Ly9zdHMud2luZG93cy5uZXQvN2Y3YjkzNTctOWM0NC00NDEwLTk1ZGYtMmM1OWI3YzE4NzJiLyIsImlhdCI6MTc0NjQzODMxNiwibmJmIjoxNzQ2NDM4MzE2LCJleHAiOjE3NDY0NDIyMTYsImFpbyI6ImsyUmdZSkRvZWZ6SlhhL0o2bTNEdHhrdTNSWTNBUT09IiwiYXBwaWQiOiI3MzQ5ZGZiMS1kZTM0LTQyYmMtOWU1NS01NTQ2OWU3ODU1ZWUiLCJhcHBpZGFjciI6IjEiLCJpZHAiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC83ZjdiOTM1Ny05YzQ0LTQ0MTAtOTVkZi0yYzU5YjdjMTg3MmIvIiwiaWR0eXAiOiJhcHAiLCJvaWQiOiI2NWU4M2RkYS0wMjdlLTRiYTgtYWE0MS1iYTgwMTUzZDNkMjIiLCJyaCI6IjEuQVI4QVY1TjdmMFNjRUVTVjN5eFp0OEdIS3drQUFBQUFBQUFBd0FBQUFBQUFBQUJNQVFBZkFBLiIsInN1YiI6IjY1ZTgzZGRhLTAyN2UtNGJhOC1hYTQxLWJhODAxNTNkM2QyMiIsInRpZCI6IjdmN2I5MzU3LTljNDQtNDQxMC05NWRmLTJjNTliN2MxODcyYiIsInV0aSI6InNkRUQ0SGoybmtXajlFVk9SdWdkQUEiLCJ2ZXIiOiIxLjAiLCJ4bXNfaWRyZWwiOiI3IDgifQ.Rt-ueO24kZIem6Dv2ipL83SYH_i4onx0jHPwjB0hMLzvwcCEfo6qiQq1Du5jOktBZXOjGrQpWvSTEz_qsmbn9jO6wHqFdfZS5W2h3wg1JCalg50QC0DuIJ0ENg0CLfdxjppOPVf91_Gn935v1N8PDQxnqNTpXi6uYTlNkVffvc9w2OTuVe_PgpXDwGHmTxx4hLIMDtNAKu9Y2uMPg9DE9NS5WqAywg2KRGio0YcXb7kubgcRsavrpZPGOhX070lzWavDJOLgcB-tHhiRhDa7vJ7jZGooU8GtIlLlKciSYu7vbcMYos2fGaoNUaCPi1HFE0HwiQGcjECmY6Lnq6A36w"

func getDataset() {
	req, _ := http.NewRequest("GET", "https://api.powerbi.com/v1.0/myorg/groups", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Power BI response:", resp.StatusCode, resp.Status)
	fmt.Println(string(body))
}

func main() {
	//tenantID := "ВАШ_TENANT_ID"
	//clientID := "ВАШ_CLIENT_ID"
	//clientSecret := "ВАШ_CLIENT_SECRET"
	//
	//authURL := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tenantID)
	//
	//data := url.Values{}
	//data.Set("grant_type", "client_credentials")
	//data.Set("client_id", clientID)
	//data.Set("client_secret", clientSecret)
	//data.Set("scope", "https://analysis.windows.net/powerbi/api/.default")
	//
	//resp, err := http.PostForm(authURL, data)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//body, _ := ioutil.ReadAll(resp.Body)
	//
	//var tokenResp struct {
	//	AccessToken string `json:"access_token"`
	//}
	//if err := json.Unmarshal(body, &tokenResp); err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Access token:", tokenResp.AccessToken)

	// Теперь можно использовать токен для вызова Power BI API
	//callPowerBI(tokenResp.AccessToken)
	getDataset()
	//callPowerBI(accessToken)
}

func callPowerBI(accessToken string) {
	req, _ := http.NewRequest("GET", "https://api.powerbi.com/v1.0/myorg/groups", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Power BI response:", resp.StatusCode, resp.Status)
	fmt.Println("Power BI response:", string(body))
}
