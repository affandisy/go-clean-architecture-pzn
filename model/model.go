package model

type WebResponse[T any] struct {
	Data T `json:"data,omitempty"`
	// PagingResponse *PagingResponse `json:"paging,omitempty"`
}

type PagingResponse[T any] struct {
	Data         []T          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"paging,omitempty"`
}

type PageMetadata struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}
