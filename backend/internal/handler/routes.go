// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	collection "backend/internal/handler/collection"
	comment "backend/internal/handler/comment"
	statistics "backend/internal/handler/statistics"
	tag "backend/internal/handler/tag"
	"backend/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 增加页面
				Method:  http.MethodPost,
				Path:    "/collection",
				Handler: collection.AddCollectionHandler(serverCtx),
			},
			{
				// 删除页面
				Method:  http.MethodDelete,
				Path:    "/collection",
				Handler: collection.DeleteCollectionHandler(serverCtx),
			},
			{
				// 更新页面
				Method:  http.MethodPut,
				Path:    "/collection",
				Handler: collection.UpdateCollectionHandler(serverCtx),
			},
			{
				// 单页面
				Method:  http.MethodGet,
				Path:    "/collection/:id",
				Handler: collection.CollectionHandler(serverCtx),
			},
			{
				// 页面集合
				Method:  http.MethodGet,
				Path:    "/collections",
				Handler: collection.CollectionsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 发布评论
				Method:  http.MethodPost,
				Path:    "/comment",
				Handler: comment.AddCommentHandler(serverCtx),
			},
			{
				// 删除评论
				Method:  http.MethodDelete,
				Path:    "/comment",
				Handler: comment.DeleteCommentHandler(serverCtx),
			},
			{
				// 获取页面评论
				Method:  http.MethodGet,
				Path:    "/comments",
				Handler: comment.CommentHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 增加页面浏览量
				Method:  http.MethodPost,
				Path:    "/statistics/:id",
				Handler: statistics.StatisticsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 单tag
				Method:  http.MethodGet,
				Path:    "/tag/:id",
				Handler: tag.TagHandler(serverCtx),
			},
			{
				// tags页面集合
				Method:  http.MethodGet,
				Path:    "/tags",
				Handler: tag.TagsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}
