// Code generated by goctl. DO NOT EDIT.
package types

type Article struct {
	Type    *string `json:"type,optional"`
	Content *string `json:"content,optional"`
	Author  string  `json:"author"`
}

type ArticleRes struct {
	Article
	Id int64 `json:"id"`
}

type ArticleId struct {
	Id int64 `path:"id"`
}
