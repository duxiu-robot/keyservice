// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 07f56d626f
// Version Date: 2020-12-21T21:40:22Z

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	pb "github.com/duxiu-robot/keyservice/interface-defs"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	EncryptEndpoint      endpoint.Endpoint
	EncryptBatchEndpoint endpoint.Endpoint
	DecryptEndpoint      endpoint.Endpoint
	DecryptBatchEndpoint endpoint.Endpoint
	KeysEndpoint         endpoint.Endpoint
	PingEndpoint         endpoint.Endpoint
}

// Endpoints

func (e Endpoints) Encrypt(ctx context.Context, in *pb.EncryptRequest) (*pb.Response, error) {
	response, err := e.EncryptEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}

func (e Endpoints) EncryptBatch(ctx context.Context, in *pb.EncryptBatchRequest) (*pb.BatchResponse, error) {
	response, err := e.EncryptBatchEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.BatchResponse), nil
}

func (e Endpoints) Decrypt(ctx context.Context, in *pb.DecryptRequest) (*pb.Response, error) {
	response, err := e.DecryptEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}

func (e Endpoints) DecryptBatch(ctx context.Context, in *pb.DecryptBatchRequest) (*pb.BatchResponse, error) {
	response, err := e.DecryptBatchEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.BatchResponse), nil
}

func (e Endpoints) Keys(ctx context.Context, in *pb.KeyRequest) (*pb.KeyResponse, error) {
	response, err := e.KeysEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.KeyResponse), nil
}

func (e Endpoints) Ping(ctx context.Context, in *pb.Empty) (*pb.Response, error) {
	response, err := e.PingEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}

// Make Endpoints

func MakeEncryptEndpoint(s pb.KeyServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.EncryptRequest)
		v, err := s.Encrypt(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeEncryptBatchEndpoint(s pb.KeyServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.EncryptBatchRequest)
		v, err := s.EncryptBatch(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeDecryptEndpoint(s pb.KeyServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.DecryptRequest)
		v, err := s.Decrypt(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeDecryptBatchEndpoint(s pb.KeyServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.DecryptBatchRequest)
		v, err := s.DecryptBatch(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeKeysEndpoint(s pb.KeyServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.KeyRequest)
		v, err := s.Keys(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakePingEndpoint(s pb.KeyServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.Empty)
		v, err := s.Ping(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"Encrypt":      {},
		"EncryptBatch": {},
		"Decrypt":      {},
		"DecryptBatch": {},
		"Keys":         {},
		"Ping":         {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "Encrypt" {
			e.EncryptEndpoint = middleware(e.EncryptEndpoint)
		}
		if inc == "EncryptBatch" {
			e.EncryptBatchEndpoint = middleware(e.EncryptBatchEndpoint)
		}
		if inc == "Decrypt" {
			e.DecryptEndpoint = middleware(e.DecryptEndpoint)
		}
		if inc == "DecryptBatch" {
			e.DecryptBatchEndpoint = middleware(e.DecryptBatchEndpoint)
		}
		if inc == "Keys" {
			e.KeysEndpoint = middleware(e.KeysEndpoint)
		}
		if inc == "Ping" {
			e.PingEndpoint = middleware(e.PingEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"Encrypt":      {},
		"EncryptBatch": {},
		"Decrypt":      {},
		"DecryptBatch": {},
		"Keys":         {},
		"Ping":         {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "Encrypt" {
			e.EncryptEndpoint = middleware("Encrypt", e.EncryptEndpoint)
		}
		if inc == "EncryptBatch" {
			e.EncryptBatchEndpoint = middleware("EncryptBatch", e.EncryptBatchEndpoint)
		}
		if inc == "Decrypt" {
			e.DecryptEndpoint = middleware("Decrypt", e.DecryptEndpoint)
		}
		if inc == "DecryptBatch" {
			e.DecryptBatchEndpoint = middleware("DecryptBatch", e.DecryptBatchEndpoint)
		}
		if inc == "Keys" {
			e.KeysEndpoint = middleware("Keys", e.KeysEndpoint)
		}
		if inc == "Ping" {
			e.PingEndpoint = middleware("Ping", e.PingEndpoint)
		}
	}
}
