# go-pixiv

[![Go Reference](https://pkg.go.dev/badge/github.com/ryohidaka/go-pixiv.svg)](https://pkg.go.dev/github.com/ryohidaka/go-pixiv)
![GitHub Release](https://img.shields.io/github/v/release/ryohidaka/go-pixiv)
[![codecov](https://codecov.io/gh/ryohidaka/go-pixiv/graph/badge.svg?token=Q7U8FMv9bn)](https://codecov.io/gh/ryohidaka/go-pixiv)
[![Go Report Card](https://goreportcard.com/badge/github.com/ryohidaka/go-pixiv)](https://goreportcard.com/report/github.com/ryohidaka/go-pixiv)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Pixiv API for Golang

Inspired by [pixivpy](https://github.com/upbit/pixivpy/tree/master)

> [!IMPORTANT]
> Only authentication with a refresh token is supported.
>
> Please check [pixivpy's README](https://github.com/upbit/pixivpy/tree/master) for instructions on how to obtain a refresh token.

## Installation

```bash
go get github.com/ryohidaka/go-pixiv
```

## Documentation

Read [GoDoc](https://pkg.go.dev/github.com/ryohidaka/go-pixiv)

## API functions

### App-API (6.0 - app-api.pixiv.net)

```go
// 用户详情
UserDetail(uid uint64, opts *UserDetailOptions) (*models.UserDetail, error) {...}

// 用户作品列表
UserIllusts(uid uint64, opts *UserIllustsOptions) ([]models.Illust, int, error) {...}

// 用户收藏作品列表
UserBookmarksIllust(uid uint64, opts *UserBookmarksIllustOptions) ([]models.Illust, int, error) {...}

// 关注用户的新作
IllustFollow(opts *IllustFollowOptions) ([]models.Illust, int, error) {...}

// 作品详情 (类似PAPI.works(),iOS中未使用)
IllustDetail(id uint64) (*models.Illust, error) {...}

// 作品收藏详情
IllustBookmarkDetail(id uint64) (*models.IllustBookmarkDetail, error) {...}

// Following用户列表
UserFollowing(userID uint64, opts *UserFollowingOptions) ([]models.UserPreview, int, error) {...}

// Followers用户列表
UserFollower(userID uint64, opts *UserFollowerOptions) ([]models.UserPreview, int, error) {...}

FetchAllUserIllusts(uid uint64, opts *UserIllustsOptions, sleepMs ...int) ([]models.Illust, error)

FetchAllBookmarkedIllusts(uid uint64, opts *UserBookmarksIllustOptions, sleepMs ...int) ([]models.Illust, error)
```

## Usage

```go
// Create a new Pixiv App API client
app, err := pixiv.NewApp("<YOUR_REFRESH_TOKEN>")

// Fetch user details
user, err := app.UserDetail(11)

// Fetch user illusts
illusts, next, err := app.UserIllusts(11, nil)

// Fetch user bookmarks illust
illusts, next, err := app.UserBookmarksIllust(11, nil)

// Fetch illust from user follows
illusts, next, err := app.IllustFollow(nil)

// Fetch illust details
illust, err := app.IllustDetail(129899459)

// Fetch illust bookmark details
bookmark, err := app.IllustBookmarkDetail(129899459)

// Fetch user following
users, next, err := api.UserFollowing(11, nil)

// Fetch user follower
users, next, err := api.UserFollower(11, nil)

// Fetch all user illusts
illusts, err := app.FetchAllUserIllusts(11, nil)

// Fetch all user bookmarks illust
illusts, err := app.FetchAllBookmarkedIllusts(11, nil)
```

## Link

- [pixivpy](https://github.com/upbit/pixivpy/tree/master)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
