package main

import (
	"testing"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"github.com/sawadashota/hcmcafe/server/handler"
)

const EndPoint = "http://localhost:9000/rpc"

type HealthCheckRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Id      int    `json:"id"`
}

type HealthCheckResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result handler.PongReply `json:"result"`
}

// Run server before run test
func TestHealthCheck(t *testing.T) {
	params, err := json.Marshal(HealthCheckRequest{
		Jsonrpc: "2.0",
		Method:  "HealthCheck.Ping",
		Id:      1,
	})

	if err != nil {
		t.Errorf("%v\n", err)
	}

	resp, err := http.Post(EndPoint, "application/json", bytes.NewBuffer(params))

	if err != nil {
		t.Errorf("Run server before run test\n$ make serve\n")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("%v\n", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Http response status should be 200 but %d\n", resp.StatusCode)
	}

	output := HealthCheckResponse{}
	err = json.Unmarshal(body, &output)

	if err != nil {
		t.Errorf("%v\n", err)
	}

	if output.Result.Message != "pong!" {
		t.Errorf("Pong message expect 'pong!' but %s", output.Result.Message)
	}
}
