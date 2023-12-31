package main

import (
	"context"
	comment "simple-douyin/kitex_gen/comment"
	"simple-douyin/service/comment/service"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAddAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAddAction(ctx context.Context, req *comment.CommentAddActionRequest) (resp *comment.CommentAddActionResponse, err error) {
	// TODO: Your code here...
	// 更新视频的被评论总数CommentCount
	return
}

// CommentDelAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentDelAction(ctx context.Context, req *comment.CommentDelActionRequest) (resp *comment.CommentDelActionResponse, err error) {
	// TODO: Your code here...
	// 更新视频的CommentCount
	return
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentCount implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentCount(ctx context.Context, req *comment.CommentCountRequest) (resp *comment.CommentCountResponse, err error) {
	resp = new(comment.CommentCountResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.CommentCount(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57005
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
	}
	// 后处理返回结果
	// ...
	return resp, nil
}
