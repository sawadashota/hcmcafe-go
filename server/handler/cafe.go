package handler

import "net/http"

type Cafe struct{}

type CafeEmptyArgs struct {
}

type CafeListArgs struct {
	Limit  int    `json:"limit"`
	Sort   string `json:"sort"`
	Filter Filter `json:"filter"`
}

type Filter struct {
	IncludeDraft     bool `json:"include_draft"`
	IncludeClose     bool `json:"include_close"`
	IncludeUnpopular bool `json:"include_unpopular"`
}

type CafeListReply struct {
	Message string `json:"message"`
}

type CafeDetailReply struct {
	Message string `json:"message"`
}

// List of all cafes
func (c *Cafe) List(r *http.Request, args *CafeEmptyArgs, reply *CafeListReply) error {
	reply.Message = "Cafe List here..."
	return nil
}

//
func (c *Cafe) Detail(r *http.Request, args *CafeEmptyArgs, reply *CafeDetailReply) error {
	reply.Message = "Cafe Detail"
	return nil
}
