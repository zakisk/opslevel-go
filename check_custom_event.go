package opslevel

type CustomEventCheckFragment struct {
	Integration      IntegrationId `graphql:"integration"`
	PassPending      bool          `graphql:"passPending"`
	ResultMessage    string        `graphql:"resultMessage"`
	ServiceSelector  string        `graphql:"serviceSelector"`
	SuccessCondition string        `graphql:"successCondition"`
}

type CheckCustomEventCreateInput struct {
	CheckCreateInput

	Integration      ID     `json:"integrationId"`
	ServiceSelector  string `json:"serviceSelector"`
	SuccessCondition string `json:"successCondition"`
	Message          string `json:"resultMessage,omitempty"`
	PassPending      *bool  `json:"passPending,omitempty"`
}

type CheckCustomEventUpdateInput struct {
	CheckUpdateInput

	ServiceSelector  string  `json:"serviceSelector,omitempty"`
	SuccessCondition string  `json:"successCondition,omitempty"`
	Message          *string `json:"resultMessage,omitempty"`
	PassPending      *bool   `json:"passPending,omitempty"`
	Integration      *ID     `json:"integrationId,omitempty"`
}

func (client *Client) CreateCheckCustomEvent(input CheckCustomEventCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkCustomEventCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckCustomEventCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckCustomEvent(input CheckCustomEventUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkCustomEventUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckCustomEventUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
