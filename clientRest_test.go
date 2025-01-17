package opslevel_test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func ATestRestClient(t *testing.T, endpoint string) *resty.Client {
	return ol.NewRestClient(ol.SetURL(
		autopilot.RegisterEndpoint(
			fmt.Sprintf("/%s", endpoint),
			autopilot.FixtureResponse(fmt.Sprintf("%s_response.json", endpoint)),
			autopilot.SkipRequestValidation(),
		),
	),
	)
}

func TestRestClientQuery(t *testing.T) {
	// Arrange
	client := ATestRestClient(t, "rest/example")
	resp := &ol.RestResponse{}
	// Act
	_, err := client.R().SetResult(resp).Get("")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Hello World!", resp.Result)
}
