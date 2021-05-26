package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	githubactions "github.com/sethvargo/go-githubactions"
)

type args struct {
	name  string
	value string
}

const (
	serverIdx    = iota
	variableIdx  = iota
	namespaceIdx = iota
	workflowIdx  = iota
	tokenIdx     = iota
	dataIdx      = iota
	protocolIdx  = iota
)

func main() {

	in := []args{
		args{
			name: "server",
		},
		args{
			name: "namespace",
		},
		args{
			name: "workflow",
		},
		args{
			name: "token",
		},
		args{
			name: "variable",
		},
		args{
			name: "data",
		},
		args{
			name: "protocol",
		},
	}

	for i := range in {
		getValue(&in[i].value, in[i].name)
	}

	githubactions.Infof("using server: %v\n", in[serverIdx].value)

	githubactions.Infof("IN: %+v", in)
	if in[serverIdx].value == "" || in[namespaceIdx].value == "" {
		githubactions.Fatalf("server and workflow values are required\n")
	}

	doRequest(in)
}

func doRequest(in []args) {
	namespace := in[namespaceIdx].value
	workflow := in[workflowIdx].value
	variable := in[variableIdx].value

	urlPath := "/api/namespaces/%s"

	if workflow != "" {
		urlPath += fmt.Sprintf("/workflows/%s/variables/%s", workflow, variable)
	} else {
		urlPath += fmt.Sprintf("/variables/%s", variable)
	}

	githubactions.Infof("uploading data to %s/%s", namespace, workflow)

	u := &url.URL{}
	u.Scheme = in[protocolIdx].value
	u.Host = in[serverIdx].value
	u.Path = fmt.Sprintf("%s", urlPath)
	var err error
	var dataGettingSent []byte
	fi, _ := os.Stat(in[dataIdx].value)
	if fi == nil {
		// pass the data as string input instead
		dataGettingSent = []byte(in[dataIdx].value)
	} else {
		dataGettingSent, err = ioutil.ReadFile(in[dataIdx].value)
		if err != nil {
			githubactions.Fatalf("unable to read file to parse data: %v", err)
		}
	}

	// TODO apply data here
	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(dataGettingSent))
	if err != nil {
		githubactions.Fatalf("can not create request: %v", err)
	}

	if len(in[tokenIdx].value) > 0 {
		githubactions.Infof("using token authentication\n")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", in[tokenIdx].value))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		githubactions.Fatalf("can not post request: %v", err)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		githubactions.Fatalf("unable to read response: %v", err)
	}

	githubactions.SetOutput("body", fmt.Sprintf("%s", data))
}

func getValue(val *string, key string) {
	*val = githubactions.GetInput(key)
}
