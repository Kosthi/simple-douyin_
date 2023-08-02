package main

import (
	"context"
	relation "simple-douyin/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAdd implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAdd(ctx context.Context, request *relation.RelationAddRequest) (resp *relation.RelationAddResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationRemove implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationRemove(ctx context.Context, request *relation.RelationRemoveRequest) (resp *relation.RelationRemoveResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, request *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, request *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, request *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowCount implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowCount(ctx context.Context, request *relation.RelationFollowCountRequest) (resp *relation.RelationFollowCountResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowerCount implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerCount(ctx context.Context, request *relation.RelationFollowerCountRequest) (resp *relation.RelationFollowerCountResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationIsFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationIsFollow(ctx context.Context, request *relation.RelationIsFollowRequest) (resp *relation.RelationIsFollowResponse, err error) {
	// TODO: Your code here...
	return
}