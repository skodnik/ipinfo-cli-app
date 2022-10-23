//go:build integration
// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testIp = "8.8.8.8"

func Test_makeRequest(t *testing.T) {
	resp := makeRequest(testIp, "")

	assert.Equal(t, 200, resp.StatusCode)
}

func Test_getBody(t *testing.T) {
	body := getBody(makeRequest(testIp, ""))

	assert.NotEmpty(t, body)
}

func Test_convertToIpInfo(t *testing.T) {
	body := getBody(makeRequest(testIp, ""))
	ipData := convertToIpInfo(body)

	assert.Equal(t, ipData.Ip, testIp)
}
