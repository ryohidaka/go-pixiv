package models

import "github.com/ryohidaka/go-pixiv/models/appmodel"



type UserDetail = appmodel.UserDetail
type Illust = appmodel.Illust
type IllustBookmarkDetail = appmodel.IllustBookmarkDetail
type IllustType = appmodel.IllustType
type UserPreview = appmodel.UserPreview
type Restrict = appmodel.Restrict

const (
	Public  = appmodel.Public
	Private = appmodel.Private
)
