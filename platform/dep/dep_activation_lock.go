package dep

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/liuds832/micromdm/dep"
	"github.com/liuds832/micromdm/pkg/httputil"
	"github.com/micromdm/micromdm/pkg/activationlock"
)

func (svc *DEPService) SetActivationLock(ctx context.Context, p *dep.ActivationLockRequest) (*dep.ActivationLockResponse, error) {
	if svc.client == nil {
		return nil, errors.New("DEP not configured yet. add a DEP token to enable DEP")
	}
	return svc.client.ActivationLock(p)
}

type activationLockRequest struct {
	*dep.ActivationLockRequest
}

type activationLockResponse struct {
	*dep.ActivationLockResponse
	BypassCode       string `json:"bypasscode"`
	Err error `json:"err,omitempty"`
}

func (r activationLockResponse) Failed() error { return r.Err }

func decodeActivationLockRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req activationLockRequest
	err := httputil.DecodeJSONRequest(r, &req)
	return req, err
}

func decodeActivationLockResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var resp activationLockResponse
	err := httputil.DecodeJSONResponse(r, &resp)
	return resp, err
}

func MakeSetActivationLockEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		/*生成Bypass code*/
		code, cerr := activationlock.Create(nil)
		if cerr == nil {
			return activationLockResponse{nil, Err: cerr}, nil 
		}

		req := request.(activationLockRequest)
		rep.ActivationLockRequest.EscrowKey = code.Hash()
		resp, err := svc.SetActivationLock(ctx, req.ActivationLockRequest)//调下面Endpoints的SetActivationLock
		return activationLockResponse{ActivationLockResponse: resp, BypassCode: code.String(), Err: err}, nil
	}
}

func (e Endpoints) SetActivationLock(ctx context.Context, p *dep.ActivationLockRequest) (*dep.ActivationLockResponse, error) {
	request := activationLockRequest{ActivationLockRequest: p}
	response, err := e.SetActivationLockEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(activationLockResponse).ActivationLockResponse, response.(activationLockResponse).Err
}
