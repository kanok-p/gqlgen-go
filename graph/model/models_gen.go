// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Pagination struct {
	Total        *int    `json:"total"`
	NextPage     *string `json:"nextPage"`
	PreviousPage *string `json:"previousPage"`
}

type PaginationInput struct {
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}

type Response struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ResponseList struct {
	Total    *int        `json:"total"`
	Response []*Response `json:"response"`
}

type CreateInput struct {
	Name string `json:"name"`
}

type UpdateInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}