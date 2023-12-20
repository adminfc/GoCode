package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type GraphQLRequest struct {
	Query string `json:"query"`
}

type GraphQLResponse struct {
	Data struct {
		Organization struct {
			Assets struct {
				Edges []struct {
					Node struct {
						Host struct {
							HostNames []struct {
								Name string `json:"name"`
							} `json:"hostNames"`
						} `json:"host"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"assets"`
		} `json:"organization"`
	} `json:"data"`
}

func main() {
	url := "https://us.api.insight.rapid7.com/graphql"

	query := `{ organization(id: "33c73f5d-e82e-4733-bd37-f4d59aace443"){ assets(first: 5000 filter: {platform:WINDOWS}){ edges{ node{ host{ hostNames{name} } } } } } }`
	reqBody := &GraphQLRequest{
		Query: query,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Api-Key", "c8009e53-35bb-40b6-bd48-e9759b97e9d5")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-version", "kratos")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var respData GraphQLResponse
	if err := json.Unmarshal(respBody, &respData); err != nil {
		panic(err)
	}

	file, err := os.OpenFile("rapid7.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, edge := range respData.Data.Organization.Assets.Edges {
		for _, hostName := range edge.Node.Host.HostNames {
			if strings.Contains(hostName.Name, ".") {
				file.WriteString(hostName.Name + "\n")
			}
		}
	}

	fmt.Println("Done, check rapid7.txt for results")
}
