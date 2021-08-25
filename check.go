package opslevel

import (
	"fmt"

	"github.com/relvacode/iso8601"
	"github.com/shurcooL/graphql"
)

type CheckType string

const (
	CheckTypeHasOwner         CheckType = "has_owner"
	CheckTypeServiceProperty  CheckType = "service_property"
	CheckTypeHasServiceConfig CheckType = "has_service_config"
	CheckTypeHasRepository    CheckType = "has_repository"
	CheckTypeToolUsage        CheckType = "tool_usage"
	CheckTypeTagDefined       CheckType = "tag_defined"
	CheckTypeRepoFile         CheckType = "repo_file"
	CheckTypeRepoSearch       CheckType = "repo_search"
	CheckTypeCustom           CheckType = "custom"
	CheckTypePayload          CheckType = "payload"
	CheckTypeManual           CheckType = "manual"
	CheckTypeGeneric          CheckType = "generic"
)

type CheckOwner struct {
	Team Team `graphql:"... on Team"`
	// User User `graphql:"... on User"` // TODO: will this be public?
}

type Check struct {
	Category    Category   `graphql:"category"`
	Description string     `graphql:"description"`
	Enabled     bool       `graphql:"enabled"`
	Filter      Filter     `graphql:"filter"`
	Id          graphql.ID `graphql:"id"`
	Level       Level      `graphql:"level"`
	Name        string     `graphql:"name"`
	Notes       string     `graphql:"notes"`
	Owner       CheckOwner `graphql:"owner"`
	Type        CheckType  `graphql:"type"`

	CustomEventCheckFragment      `graphql:"... on CustomEventCheck"`
	ManualCheckFragment           `graphql:"... on ManualCheck"`
	RepositoryFileCheckFragment   `graphql:"... on RepositoryFileCheck"`
	RepositorySearchCheckFragment `graphql:"... on RepositorySearchCheck"`
	ServicePropertyCheckFragment  `graphql:"... on ServicePropertyCheck"`
	TagDefinedCheckFragment       `graphql:"... on TagDefinedCheck"`
	ToolUsageCheckFragment        `graphql:"... on ToolUsageCheck"`
}

type CustomEventCheckFragment struct {
	Integration      Integration `graphql:"integration"`
	ResultMessage    string      `graphql:"resultMessage"`
	ServiceSelector  string      `graphql:"serviceSelector"`
	SuccessCondition string      `graphql:"successCondition"`
}

type ManualCheckFragment struct {
	UpdateFrequency       *ManualCheckFrequency `graphql:"updateFrequency"`
	UpdateRequiresComment bool                  `graphql:"updateRequiresComment"`
}

type RepositoryFileCheckFragment struct {
	DirectorySearch       bool       `graphql:"directorySearch"`
	Filepaths             []string   `graphql:"filePaths"`
	FileContentsPredicate *Predicate `graphql:"fileContentsPredicate"`
}

type RepositorySearchCheckFragment struct {
	FileExtensions        []string  `graphql:"fileExtensions"`
	FileContentsPredicate Predicate `graphql:"fileContentsPredicate"`
}

type ServicePropertyCheckFragment struct {
	Property  ServiceProperty `graphql:"serviceProperty"`
	Predicate *Predicate      `graphql:"propertyValuePredicate"`
}

type TagDefinedCheckFragment struct {
	TagKey       string     `graphql:"tagKey"`
	TagPredicate *Predicate `graphql:"tagPredicate"`
}

type ToolUsageCheckFragment struct {
	ToolCategory         ToolCategory `graphql:"toolCategory"`
	ToolNamePredicate    *Predicate   `graphql:"toolNamePredicate"`
	EnvironmentPredicate *Predicate   `graphql:"environmentPredicate"`
}

type CheckConnection struct {
	Nodes      []Check
	PageInfo   PageInfo
	TotalCount graphql.Int
}

type CheckCreateInputProvider interface {
	GetCheckCreateInput() *CheckCreateInput
}

type CheckCreateInput struct {
	Name     string      `json:"name"`
	Enabled  bool        `json:"enabled"`
	Category graphql.ID  `json:"categoryId"`
	Level    graphql.ID  `json:"levelId"`
	Owner    *graphql.ID `json:"ownerId,omitempty"`
	Filter   *graphql.ID `json:"filterId,omitempty"`
	Notes    string      `json:"notes,omitempty"`
}

func (c *CheckCreateInput) GetCheckCreateInput() *CheckCreateInput {
	return c
}

type CheckUpdateInputProvider interface {
	GetCheckUpdateInput() *CheckUpdateInput
}

type CheckUpdateInput struct {
	Id       graphql.ID  `json:"id"`
	Name     string      `json:"name,omitempty"`
	Enabled  *bool       `json:"enabled,omitempty"`
	Category *graphql.ID `json:"categoryId,omitempty"`
	Level    *graphql.ID `json:"levelId,omitempty"`
	Owner    *graphql.ID `json:"ownerId,omitempty"`
	Filter   *graphql.ID `json:"filterId,omitempty"`
	Notes    string      `json:"notes,omitempty"`
}

func (c *CheckUpdateInput) GetCheckUpdateInput() *CheckUpdateInput {
	return c
}

type CheckCustomEventCreateInput struct {
	CheckCreateInput

	Integration      graphql.ID `json:"integrationId"`
	ServiceSelector  string     `json:"serviceSelector"`
	SuccessCondition string     `json:"successCondition"`
	Message          string     `json:"resultMessage,omitempty"`
}

type CheckCustomEventUpdateInput struct {
	CheckUpdateInput

	ServiceSelector  string      `json:"serviceSelector,omitempty"`
	SuccessCondition string      `json:"successCondition,omitempty"`
	Message          string      `json:"resultMessage,omitempty"`
	Integration      *graphql.ID `json:"integrationId,omitempty"`
}

// FrequencyTimeScale represents the time scale type for the frequency.
type FrequencyTimeScale string

// The time scale type for the frequency.
const (
	FrequencyTimeScaleDay   FrequencyTimeScale = "day"   // Consider the time scale of days.
	FrequencyTimeScaleWeek  FrequencyTimeScale = "week"  // Consider the time scale of weeks.
	FrequencyTimeScaleMonth FrequencyTimeScale = "month" // Consider the time scale of months.
	FrequencyTimeScaleYear  FrequencyTimeScale = "year"  // Consider the time scale of years.
)

func GetFrequencyTimeScales() []string {
	return []string{
		string(FrequencyTimeScaleDay),
		string(FrequencyTimeScaleWeek),
		string(FrequencyTimeScaleMonth),
		string(FrequencyTimeScaleYear),
	}
}

type ManualCheckFrequency struct {
	StartingDate       iso8601.Time       `graphql:"startingDate"`
	FrequencyTimeScale FrequencyTimeScale `graphql:"frequencyTimeScale"`
	FrequencyValue     int                `graphql:"frequencyValue"`
}

type ManualCheckFrequencyInput struct {
	StartingDate       iso8601.Time       `json:"startingDate"`
	FrequencyTimeScale FrequencyTimeScale `json:"frequencyTimeScale"`
	FrequencyValue     int                `json:"frequencyValue"`
}

func NewManualCheckFrequencyInput(startingDate string, timeScale FrequencyTimeScale, value int) *ManualCheckFrequencyInput {
	return &ManualCheckFrequencyInput{
		StartingDate:       NewISO8601Date(startingDate),
		FrequencyTimeScale: timeScale,
		FrequencyValue:     value,
	}
}

type CheckManualCreateInput struct {
	CheckCreateInput

	UpdateFrequency       *ManualCheckFrequencyInput `json:"updateFrequency,omitempty"`
	UpdateRequiresComment bool                       `json:"updateRequiresComment"`
}

type CheckManualUpdateInput struct {
	CheckUpdateInput

	UpdateFrequency       *ManualCheckFrequencyInput `json:"updateFrequency,omitempty"`
	UpdateRequiresComment bool                       `json:"updateRequiresComment,omitempty"`
}

type CheckRepositoryFileCreateInput struct {
	CheckCreateInput

	DirectorySearch       bool            `json:"directorySearch"`
	Filepaths             []string        `json:"filePaths"`
	FileContentsPredicate *PredicateInput `json:"fileContentsPredicate,omitempty"`
}

type CheckRepositoryFileUpdateInput struct {
	CheckUpdateInput

	DirectorySearch       bool            `json:"directorySearch,omitempty"`
	Filepaths             []string        `json:"filePaths,omitempty"`
	FileContentsPredicate *PredicateInput `json:"fileContentsPredicate,omitempty"`
}

type CheckRepositoryIntegratedCreateInput struct {
	CheckCreateInput
}

type CheckRepositoryIntegratedUpdateInput struct {
	CheckUpdateInput
}

type CheckRepositorySearchCreateInput struct {
	CheckCreateInput

	FileExtensions        []string       `json:"fileExtensions,omitempty"`
	FileContentsPredicate PredicateInput `json:"fileContentsPredicate"`
}

type CheckRepositorySearchUpdateInput struct {
	CheckUpdateInput

	FileExtensions        []string        `json:"fileExtensions,omitempty"`
	FileContentsPredicate *PredicateInput `json:"fileContentsPredicate,omitempty"`
}

type CheckServiceConfigurationCreateInput struct {
	CheckCreateInput
}

type CheckServiceConfigurationUpdateInput struct {
	CheckUpdateInput
}

type CheckServiceOwnershipCreateInput struct {
	CheckCreateInput
}

type CheckServiceOwnershipUpdateInput struct {
	CheckUpdateInput
}

type ServiceProperty string

const (
	ServicePropertyDescription ServiceProperty = "description"
	ServicePropertyName        ServiceProperty = "name"
	ServicePropertyLanguage    ServiceProperty = "language"
	ServicePropertyFramework   ServiceProperty = "framework"
	ServicePropertyProduct     ServiceProperty = "product"
	ServicePropertyLifecycle   ServiceProperty = "lifecycle_index"
	ServicePropertyTier        ServiceProperty = "tier_index"
)

func GetServicePropertyTypes() []string {
	return []string{
		string(ServicePropertyDescription),
		string(ServicePropertyName),
		string(ServicePropertyLanguage),
		string(ServicePropertyFramework),
		string(ServicePropertyProduct),
		string(ServicePropertyLifecycle),
		string(ServicePropertyTier),
	}
}

type CheckServicePropertyCreateInput struct {
	CheckCreateInput

	Property  ServiceProperty `json:"serviceProperty"`
	Predicate *PredicateInput `json:"propertyValuePredicate,omitempty"`
}

type CheckServicePropertyUpdateInput struct {
	CheckUpdateInput

	Property  ServiceProperty `json:"serviceProperty,omitempty"`
	Predicate *PredicateInput `json:"propertyValuePredicate,omitempty"`
}

type CheckTagDefinedCreateInput struct {
	CheckCreateInput

	TagKey       string          `json:"tagKey"`
	TagPredicate *PredicateInput `json:"tagPredicate,omitempty"`
}

type CheckTagDefinedUpdateInput struct {
	CheckUpdateInput

	TagKey       string          `json:"tagKey,omitempty"`
	TagPredicate *PredicateInput `json:"tagPredicate,omitempty"`
}

type CheckToolUsageCreateInput struct {
	CheckCreateInput

	ToolCategory         ToolCategory    `json:"toolCategory"`
	ToolNamePredicate    *PredicateInput `json:"toolNamePredicate,omitempty"`
	EnvironmentPredicate *PredicateInput `json:"environmentPredicate,omitempty"`
}

type CheckToolUsageUpdateInput struct {
	CheckUpdateInput

	ToolCategory         ToolCategory    `json:"toolCategory,omitempty"`
	ToolNamePredicate    *PredicateInput `json:"toolNamePredicate,omitempty"`
	EnvironmentPredicate *PredicateInput `json:"environmentPredicate,omitempty"`
}

type CheckDeleteInput struct {
	Id graphql.ID `json:"id"`
}

// Encompass CheckCreatePayload and CheckUpdatePayload into 1 struct
type CheckResponsePayload struct {
	Check  Check
	Errors []OpsLevelErrors
}

func (conn *CheckConnection) Hydrate(client *Client) error {
	var q struct {
		Account struct {
			Rubric struct {
				Checks CheckConnection `graphql:"checks(after: $after, first: $first)"`
			}
		}
	}
	v := PayloadVariables{
		"first": client.pageSize,
	}
	q.Account.Rubric.Checks.PageInfo = conn.PageInfo
	for q.Account.Rubric.Checks.PageInfo.HasNextPage {
		v["after"] = q.Account.Rubric.Checks.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Rubric.Checks.Nodes {
			conn.Nodes = append(conn.Nodes, item)
		}
	}
	return nil
}

func (p *CheckResponsePayload) Mutate(client *Client, m interface{}, v map[string]interface{}) (*Check, error) {
	if err := client.Mutate(m, v); err != nil {
		return nil, err
	}
	return &p.Check, FormatErrors(p.Errors)
}

//#region Create

func (client *Client) CreateCheckCustomEvent(input CheckCustomEventCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkCustomEventCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckCustomEvent(input CheckCustomEventUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkCustomEventUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckManual(input CheckManualCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkManualCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckManual(input CheckManualUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkManualUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckRepositoryFile(input CheckRepositoryFileCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryFileCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckRepositoryFile(input CheckRepositoryFileUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryFileUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckRepositoryIntegrated(input CheckRepositoryIntegratedCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryIntegratedCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}
func (client *Client) UpdateCheckRepositoryIntegrated(input CheckRepositoryIntegratedUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryIntegratedUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckRepositorySearch(input CheckRepositorySearchCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositorySearchCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckRepositorySearch(input CheckRepositorySearchUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositorySearchUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckServiceConfiguration(input CheckServiceConfigurationCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceConfigurationCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckServiceConfiguration(input CheckServiceConfigurationUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceConfigurationUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckServiceOwnership(input CheckServiceOwnershipCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceOwnershipCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckServiceOwnership(input CheckServiceOwnershipUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceOwnershipUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckServiceProperty(input CheckServicePropertyCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServicePropertyCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckServiceProperty(input CheckServicePropertyUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServicePropertyUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckTagDefined(input CheckTagDefinedCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkTagDefinedCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckTagDefined(input CheckTagDefinedUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkTagDefinedUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckToolUsage(input CheckToolUsageCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkToolUsageCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckToolUsage(input CheckToolUsageUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkToolUsageUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

//#endregion

//#region Retrieve

func (client *Client) GetCheck(id graphql.ID) (*Check, error) {
	var q struct {
		Account struct {
			Check Check `graphql:"check(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	if q.Account.Check.Id == nil {
		return nil, fmt.Errorf("Check with ID '%s' not found!", id)
	}
	return &q.Account.Check, nil
}

func (client *Client) ListChecks() ([]Check, error) {
	var q struct {
		Account struct {
			Rubric struct {
				Checks CheckConnection `graphql:"checks"`
			}
		}
	}
	if err := client.Query(&q, nil); err != nil {
		return q.Account.Rubric.Checks.Nodes, err
	}
	if err := q.Account.Rubric.Checks.Hydrate(client); err != nil {
		return q.Account.Rubric.Checks.Nodes, err
	}
	return q.Account.Rubric.Checks.Nodes, nil
}

//#endregion

//#region Delete

func (client *Client) DeleteCheck(id graphql.ID) error {
	var m struct {
		Payload IdResponsePayload `graphql:"checkDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": CheckDeleteInput{Id: id},
	}
	return m.Payload.Mutate(client, &m, v)
}

//#endregion
