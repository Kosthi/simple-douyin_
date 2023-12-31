/*
IDL 注解说明
https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/usage/annotation/
*/
namespace go publish

include "./common.thrift"

struct PublishActionRequest {
    1: required string token (api.form = "token"), // 用户鉴权token
    2: required binary data (api.form = "data"),   // 视频数据
    3: required string title (api.form = "title"), // 视频标题
}

struct PublishActionResponse {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
}

struct PublishListRequest {
    1: required i64 user_id (api.query = "user_id"), // 用户id
    2: required string token (api.query = "token"),  // 用户鉴权token
}

struct PublishListResponse {
    1: required i32 status_code,               // 状态码，0-成功，其他值-失败
    2: optional string status_msg,             // 返回状态描述
    3: optional list<common.Video> video_list, // 用户发布的视频列表
}

service PublishService {
    PublishActionResponse PublishAction(1: PublishActionRequest req) (api.post = "/douyin/publish/action/"),
    PublishListResponse PublishList(1: PublishListRequest req) (api.get = "/douyin/publish/list/"),
}
