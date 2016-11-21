package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/dghubble/sling"
	"net/http"
	"net/http/httputil"
	//"reflect"
	"encoding/json"
)

var ServerBase *sling.Sling

func createOrEditPipeline(payload map[string]interface{}) {
	debug, _ := json.MarshalIndent(payload, "", "    ")
	log.Debug(string(debug))
	pipelineName := payload["pipeline"].(map[string]interface{})["name"].(string)
	log.Debug("Check if " + pipelineName + " exists.")
	etag, _ := getEtag(pipelineName)
	var response *http.Response
	var err error
	if etag != "" {
		log.Debug(pipelineName + " already exists. Making patch request")
		response, err = edit(payload, etag)
	} else {
		log.Debug("Creating new pipeline: " + pipelineName)
		response, err = create(payload)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Could not create or edit pipeline: " + pipelineName)
	}
	if response.StatusCode != 200 {
		log.WithFields(log.Fields{"message": response.Body}).Fatal("Could not create or edit " + pipelineName + " reason: ")

	} else {
		log.Info("Updated: " + pipelineName)
	}
}

func edit(payload map[string]interface{}, etag string) (*http.Response, error) {
	path := "/go/api/admin/pipelines/" + payload["pipeline"].(map[string]interface{})["name"].(string)
	return patchRequest(path, payload, etag)
}

func create(payload map[string]interface{}) (*http.Response, error) {
	path := "/go/api/admin/pipelines"
	return putRequest(path, payload)

}

func getEtag(name string) (string, error) {
	path := "/go/api/admin/pipelines/" + name
	resp, err := getRequest(path)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Could not get etag.")
		return "", err
	}
	etag := resp.Header.Get("Etag")

	return etag, nil
}

func getRequest(path string) (*http.Response, error) {
	req, err := ServerBase.New().
		Get(path).Request()
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

func putRequest(path string, payload map[string]interface{}) (*http.Response, error) {

	req, err := ServerBase.New().
		Post(path).BodyJSON(payload).Request()
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

func patchRequest(path string, payload map[string]interface{}, etag string) (*http.Response, error) {
	req, err := ServerBase.New().
		Set("Content-Type", "application/json").
		Set("If-Match", etag).
		Put(path).BodyJSON(payload).Request()
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

func doRequest(request *http.Request) (*http.Response, error) {
	dump, _ := httputil.DumpRequest(request, true)
	log.Debug(string(dump))
	client := &http.Client{}
	resp, err := client.Do(request)
	dump, _ = httputil.DumpResponse(resp, true)
	log.Debug(string(dump))
	if resp != nil {
		defer resp.Body.Close()
	}
	return resp, err
}

func init() {
	ServerBase = sling.New().
		SetBasicAuth(config.Username, config.Password).
		Base(config.BaseUrl).
		Set("Accept", "application/vnd.go.cd.v3+json")
}
