package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateTool(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation ToolCreate($input:ToolCreateInput!){toolCreate(input: $input){tool{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},errors{message,path}}}",
    "variables": {
		"input": {
			"category": "other",
			"displayName": "example",
			"serviceId": "{{ template "id1" }}",
			"url": "https://example.com"
		}
}}`
	response := `{"data": {
		"toolCreate": {
			"tool": {{ template "tool_1" }},
			"errors": []
		}
}}`
	client := ABetterTestClient(t, "toolCreate", request, response)
	// Act
	result, err := client.CreateTool(ol.ToolCreateInput{
		Category:    ol.ToolCategoryOther,
		DisplayName: "example",
		ServiceId:   "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
		Url:         "https://example.com",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"), result.Service.Id)
	autopilot.Equals(t, ol.ToolCategoryOther, result.Category)
	autopilot.Equals(t, "Example", result.DisplayName)
	autopilot.Equals(t, "https://example.com", result.Url)
}

func TestUpdateTool(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation ToolUpdate($input:ToolUpdateInput!){toolUpdate(input: $input){tool{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},errors{message,path}}}",
    "variables": {
		"input": {
			"id": "{{ template "id1" }}",
			"category": "deployment"
		}
}}`
	response := `{"data": {
		"toolUpdate": {
			"tool": {{ template "tool_1_update" }},
			"errors": []
		}
}}`
	client := ABetterTestClient(t, "toolUpdate", request, response)
	// Act
	result, err := client.UpdateTool(ol.ToolUpdateInput{
		Id:       "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
		Category: ol.ToolCategoryDeployment,
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ToolCategoryDeployment, result.Category)
	autopilot.Equals(t, "prod", result.Environment)
}

func TestDeleteTool(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation ToolDelete($input:ToolDeleteInput!){toolDelete(input: $input){errors{message,path}}}",
    "variables": {
		"input": {
			"id": "{{ template "id1" }}"
		}
}}`
	response := `{"data": {
		"toolDelete": {
			"errors": []
		}
}}`
	client := ABetterTestClient(t, "toolDelete", request, response)
	// Act
	err := client.DeleteTool("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx")
	// Assert
	autopilot.Ok(t, err)
}
