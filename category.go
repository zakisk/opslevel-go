package opslevel

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/hasura/go-graphql-client"
)

type Category struct {
	Id   ID `json:"id"`
	Name string
}

type CategoryConnection struct {
	Nodes      []Category
	PageInfo   PageInfo
	TotalCount graphql.Int
}

type CategoryCreateInput struct {
	Name string `json:"name"`
}

type CategoryUpdateInput struct {
	Id   ID     `json:"id"`
	Name string `json:"name"`
}

type CategoryDeleteInput struct {
	Id ID `json:"id"`
}

func (self *Category) Alias() string {
	return slug.Make(self.Name)
}

func (conn *CategoryConnection) Hydrate(client *Client) error {
	var q struct {
		Account struct {
			Rubric struct {
				Categories CategoryConnection `graphql:"categories(after: $after, first: $first)"`
			}
		}
	}
	v := PayloadVariables{
		"first": client.pageSize,
	}
	q.Account.Rubric.Categories.PageInfo = conn.PageInfo
	for q.Account.Rubric.Categories.PageInfo.HasNextPage {
		v["after"] = q.Account.Rubric.Categories.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Rubric.Categories.Nodes {
			conn.Nodes = append(conn.Nodes, item)
		}
	}
	return nil
}

//#region Create

func (client *Client) CreateCategory(input CategoryCreateInput) (*Category, error) {
	var m struct {
		Payload struct {
			Category Category
			Errors   []OpsLevelErrors
		} `graphql:"categoryCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v)
	return &m.Payload.Category, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (client *Client) GetCategory(id ID) (*Category, error) {
	var q struct {
		Account struct {
			Category Category `graphql:"category(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	err := client.Query(&q, v)
	if q.Account.Category.Id == "" {
		err = fmt.Errorf("Category with ID '%s' not found!", id)
	}
	return &q.Account.Category, HandleErrors(err, nil)
}

func (client *Client) ListCategories() ([]Category, error) {
	var q struct {
		Account struct {
			Rubric struct {
				Categories CategoryConnection
			}
		}
	}
	if err := client.Query(&q, nil); err != nil {
		return q.Account.Rubric.Categories.Nodes, err
	}
	if err := q.Account.Rubric.Categories.Hydrate(client); err != nil {
		return q.Account.Rubric.Categories.Nodes, err
	}
	return q.Account.Rubric.Categories.Nodes, nil
}

//#endregion

//#region Update

func (client *Client) UpdateCategory(input CategoryUpdateInput) (*Category, error) {
	var m struct {
		Payload struct {
			Category Category
			Errors   []OpsLevelErrors
		} `graphql:"categoryUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v)
	return &m.Payload.Category, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteCategory(id ID) error {
	var m struct {
		Payload struct {
			Id     ID `graphql:"deletedCategoryId"`
			Errors []OpsLevelErrors
		} `graphql:"categoryDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": CategoryDeleteInput{Id: id},
	}
	err := client.Mutate(&m, v)
	return HandleErrors(err, m.Payload.Errors)
}

//#endregion
