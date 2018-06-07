package handler

import (
	"net/http"
	"github.com/sawadashota/hcmcafe/server/repository"
	"github.com/sawadashota/hcmcafe/server/domain"
)

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

type CafeStoreRequest struct {
	Name string `json:"name"`
}

type CafeStoreResponse struct {
	Status int `json:"status"`
}

// List of all cafes
func (c *Cafe) List(r *http.Request, args *CafeEmptyArgs, reply *CafeListReply) error {
	reply.Message = "Cafe List here..."
	return nil
}

// Detail shows all information of cafe
func (c *Cafe) Detail(r *http.Request, args *CafeEmptyArgs, reply *CafeDetailReply) error {
	reply.Message = "Cafe Detail"
	return nil
}

// StoreCafe cafe data
func (c *Cafe) Store(r *http.Request, args *CafeStoreRequest, reply *CafeStoreResponse) error {
	cafe := domain.NewCafe(args.Name, 10000, 30000)
	if err := repository.StoreCafe(cafe, r); err == nil {
		reply.Status = http.StatusOK
	} else {
		reply.Status = http.StatusInternalServerError
	}
	return nil
}
