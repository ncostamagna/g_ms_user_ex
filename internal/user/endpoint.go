package user

import (
	"context"

	"github.com/ncostamagna/g_ms_client/meta"
	"github.com/ncostamagna/g_ms_user_ex/pkg/response"
)

//Endpoints struct
type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	GetAllReq struct {
		FirstName string
		LastName  string
	}

	UpdateReq struct {
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Email     *string `json:"email"`
		Phone     *string `json:"phone"`
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
		Meta   *meta.Meta  `json:"meta,omitempty"`
	}
)

//MakeEndpoints handler endpoints
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		/*Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),*/
	}
}

func makeCreateEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateReq)

		if req.FirstName == "" {
			return nil, response.BadRequest("first name is required")
		}

		if req.LastName == "" {
			return nil, response.BadRequest("last name is required")
		}

		user, err := s.Create(req.FirstName, req.LastName, req.Email, req.Phone)
		if err != nil {
			return nil, response.BadRequest(err.Error())
		}

		return response.Created("success", user, nil), nil
	}
}

/*
func makeGetEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		path := mux.Vars(r)
		id := path["id"]

		user, err := s.Get(id)
		if err != nil {
			return nil, &Response{Status: 404, Err: "user doesn't exist"}
		}

		return &Response{Status: 200, Data: user}, nil
	}
}

func makeGetAllEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		v := r.URL.Query()

		filters := Filters{
			FirstName: v.Get("first_name"),
			LastName:  v.Get("last_name"),
		}

		limit, _ := strconv.Atoi(v.Get("limit"))
		page, _ := strconv.Atoi(v.Get("page"))

		count, err := s.Count(filters)
		if err != nil {
			return nil, &Response{Status: 500, Err: err.Error()}
		}

		meta, err := meta.New(page, limit, count)
		if err != nil {
			return nil, &Response{Status: 500, Err: err.Error()}
		}

		users, err := s.GetAll(filters, meta.Offset(), meta.Limit())
		if err != nil {
			return nil, &Response{Status: 400, Err: err.Error()}
		}

		return &Response{Status: 200, Data: users, Meta: meta}, nil
	}
}

func makeUpdateEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var req UpdateReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, &Response{Status: 400, Err: "invalid request format"}
		}

		if req.FirstName != nil && *req.FirstName == "" {
			return nil, &Response{Status: 400, Err: "first name is required"}
		}

		if req.LastName != nil && *req.LastName == "" {
			return nil, &Response{Status: 400, Err: "last name is required"}
		}

		path := mux.Vars(r)
		id := path["id"]

		if err := s.Update(id, req.FirstName, req.LastName, req.Email, req.Phone); err != nil {
			return nil, &Response{Status: 404, Err: "user doesn't exist"}
		}

		return &Response{Status: 200, Data: "success"}, nil
	}
}

func makeDeleteEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		path := mux.Vars(r)
		id := path["id"]

		if err := s.Delete(id); err != nil {
			return nil, &Response{Status: 404, Err: "user doesn't exist"}
		}

		return &Response{Status: 200, Data: "success"}, nil
	}
}
*/
