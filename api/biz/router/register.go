// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	feed "simple-douyin/api/biz/router/feed"
	comment "simple-douyin/api/biz/router/comment"
	favorite "simple-douyin/api/biz/router/favorite"
	message "simple-douyin/api/biz/router/message"
	relation "simple-douyin/api/biz/router/relation"
	publish "simple-douyin/api/biz/router/publish"
	user "simple-douyin/api/biz/router/user"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	feed.Register(r)

	favorite.Register(r)

	comment.Register(r)

	relation.Register(r)

	message.Register(r)

	publish.Register(r)

	user.Register(r)
}
