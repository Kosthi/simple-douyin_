package main

import (
	"context"
	servLog "github.com/sirupsen/logrus"
	"simple-douyin/kitex_gen/feed"
	"simple-douyin/service/feed/service"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	// 返回时间小于latestTime的最多30个视频
	// 可以根据req.userId推荐个性化视频

	resp, err := service.Feed(ctx, req)
	if err != nil {
		servLog.Error(err)
		if resp == nil {
			resp = feed.NewFeedResponse()
		}
		resp.StatusCode = 57002
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		return resp, err
	}

	return resp, nil
}
