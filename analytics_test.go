package analytics

import (
	"testing"

	"github.com/getlantern/testify/assert"
)

func samplePayload() *Payload {
	payload := &Payload{
		TrackingId: "UA-21815217-2",
		InstanceId: "test-client-555",
	}

	return payload
}

func TestAnalyticsRequest(t *testing.T) {

	p := samplePayload()

	status, err := SendRequest(p)
	assert.Nil(t, err)
	assert.Equal(t, true, status)

}
