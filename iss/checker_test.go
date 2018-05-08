package iss

import (
	"testing"
	"gopkg.in/jarcoal/httpmock.v1"
	"github.com/stretchr/testify/assert"
)

func TestGetIssCoordinate(test *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://api.open-notify.org/iss-now.json",
		httpmock.NewStringResponder(200, `{"message": "success", "iss_position": { "latitude": "-29.0677", "longitude": "137.3126" }, "timestamp": 1525782417 }`))

	coordinate, err := GetCoordinate()

	assert.Nil(test, err)
	assert.Equal(test, "-29.0677", coordinate.Latitude)
	assert.Equal(test, "137.3126", coordinate.Longitude)
}

func TestGetIssCoordinateWhenEndpointIsDown(test *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://api.open-notify.org/iss-now.json",
		httpmock.NewStringResponder(404, `{}`))

	_, err := GetCoordinate()
	assert.NotNil(test, err)
}