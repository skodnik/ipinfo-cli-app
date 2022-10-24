//go:build integration
// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const testIp = "8.8.8.8"

func Test_makeRequest(t *testing.T) {
	resp := makeRequest(testIp, "", false)

	assert.Equal(t, 200, resp.StatusCode)
}

func Test_getBody(t *testing.T) {
	body := getBody(makeRequest(testIp, "", false))

	assert.NotEmpty(t, body)
}

func Test_convertToIpInfo(t *testing.T) {
	body := getBody(makeRequest(testIp, "", false))
	ipData := convertToIpInfo(body)

	assert.Equal(t, ipData.Ip, testIp)
}

func Test_convertToIpInfoWithSample(t *testing.T) {
	body, _ := ioutil.ReadFile("./../../test/sample-response.json")
	ipData := convertToIpInfo(body)

	assert.Equal(t, ipData.Ip, testIp)
}
