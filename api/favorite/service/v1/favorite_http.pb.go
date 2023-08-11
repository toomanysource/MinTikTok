// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.4
// source: favorite/service/v1/favorite.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationFavoriteServiceFavoriteAction = "/api.favorite.service.v1.FavoriteService/FavoriteAction"
const OperationFavoriteServiceGetFavoriteList = "/api.favorite.service.v1.FavoriteService/GetFavoriteList"

type FavoriteServiceHTTPServer interface {
	// FavoriteAction 取消或添加喜爱视频(客户端)
	FavoriteAction(context.Context, *FavoriteActionRequest) (*FavoriteActionReply, error)
	// GetFavoriteList 获取喜爱视频列表(客户端)
	GetFavoriteList(context.Context, *FavoriteListRequest) (*FavoriteListReply, error)
}

func RegisterFavoriteServiceHTTPServer(s *http.Server, srv FavoriteServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/douyin/favorite/list", _FavoriteService_GetFavoriteList0_HTTP_Handler(srv))
	r.POST("/douyin/favorite/action", _FavoriteService_FavoriteAction0_HTTP_Handler(srv))
}

func _FavoriteService_GetFavoriteList0_HTTP_Handler(srv FavoriteServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FavoriteListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFavoriteServiceGetFavoriteList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFavoriteList(ctx, req.(*FavoriteListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FavoriteListReply)
		return ctx.Result(200, reply)
	}
}

func _FavoriteService_FavoriteAction0_HTTP_Handler(srv FavoriteServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FavoriteActionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFavoriteServiceFavoriteAction)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FavoriteAction(ctx, req.(*FavoriteActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FavoriteActionReply)
		return ctx.Result(200, reply)
	}
}

type FavoriteServiceHTTPClient interface {
	FavoriteAction(ctx context.Context, req *FavoriteActionRequest, opts ...http.CallOption) (rsp *FavoriteActionReply, err error)
	GetFavoriteList(ctx context.Context, req *FavoriteListRequest, opts ...http.CallOption) (rsp *FavoriteListReply, err error)
}

type FavoriteServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewFavoriteServiceHTTPClient(client *http.Client) FavoriteServiceHTTPClient {
	return &FavoriteServiceHTTPClientImpl{client}
}

func (c *FavoriteServiceHTTPClientImpl) FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...http.CallOption) (*FavoriteActionReply, error) {
	var out FavoriteActionReply
	pattern := "/douyin/favorite/action"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFavoriteServiceFavoriteAction))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FavoriteServiceHTTPClientImpl) GetFavoriteList(ctx context.Context, in *FavoriteListRequest, opts ...http.CallOption) (*FavoriteListReply, error) {
	var out FavoriteListReply
	pattern := "/douyin/favorite/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationFavoriteServiceGetFavoriteList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}