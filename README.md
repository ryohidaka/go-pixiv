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

## Usage

### API functions

```go
import "github.com/ryohidaka/go-pixiv"

// Create a new Pixiv App API client
app, err := pixiv.NewApp("<YOUR_REFRESH_TOKEN>")

// Fetch user details
user, err := app.UserDetail(11)

// Fetch user illusts
illusts, next, err := app.UserIllusts(11)

// Fetch user bookmarks illust
illusts, next, err := app.UserBookmarksIllust(11)

// Fetch illust from user follows
illusts, next, err := app.IllustFollow()

// Fetch illust details
illust, err := app.IllustDetail(129899459)

// Fetch illust bookmark details
bookmark, err := app.IllustBookmarkDetail(129899459)

// Fetch user following
users, next, err := api.UserFollowing(11)

// Fetch user follower
users, next, err := api.UserFollower(11)

// Follow user with user ID
ok, err := api.UserFollowAdd(11)

// Unfollow user with user ID
ok, err := api.UserFollowDelete(11)
```

### Crawler

```go
import "github.com/ryohidaka/go-pixiv/crawler"

// Create a new Pixiv Crawler
crawler, err := crawler.NewCrawler(refreshToken)

// Fetch all user illusts
illusts, err := crawler.FetchAllUserIllusts(11)

// Fetch all user bookmarks illust
illusts, err := crawler.FetchAllBookmarkedIllusts(11)

// Fetch all user follow illust
illusts, err := crawler.FetchAllIllustFollows()

// Fetch all user following
users, err := crawler.FetchAllUserFollowing(11)

// Fetch all user follower
users, err := crawler.FetchAllUserFollowers(11)

// Follow user with user IDs
processed, err := c.UserFollowAddMultiple([]uint64{11})

// Unfollow user with user IDs
processed, err := c.UserFollowDeleteMultiple([]uint64{11})
```

## API functions

### App-API (6.0 - app-api.pixiv.net)

```go
// 用户详情
UserDetail(uid uint64, opts ...UserDetailOptions) (*models.UserDetail, error) {...}

// 用户作品列表
UserIllusts(uid uint64, opts ...UserIllustsOptions) ([]models.Illust, int, error) {...}

// 用户收藏作品列表
UserBookmarksIllust(uid uint64, opts ...UserBookmarksIllustOptions) ([]models.Illust, int, error) {...}

// 关注用户的新作
IllustFollow(opts ...IllustFollowOptions) ([]models.Illust, int, error) {...}

// 作品详情 (类似PAPI.works(),iOS中未使用)
IllustDetail(id uint64) (*models.Illust, error) {...}

// 作品收藏详情
IllustBookmarkDetail(id uint64) (*models.IllustBookmarkDetail, error) {...}

// Following用户列表
UserFollowing(uid uint64, opts ...UserFollowingOptions) ([]models.UserPreview, int, error) {...}

// Followers用户列表
UserFollower(uid uint64, opts ...UserFollowerOptions) ([]models.UserPreview, int, error) {...}

// 关注用户
UserFollowAdd(uid uint64, restrict ...models.Restrict) {...}

// 取消关注用户
UserFollowDelete(uid uint64) (bool, error) {...}
```

## Crawler

```go
// 获取指定用户的作品列表
FetchAllUserIllusts(uid uint64, opts *UserIllustsOptions, sleepMs ...int) ([]models.Illust, error) {...}

// 获取指定用户的收藏列表
FetchAllBookmarkedIllusts(uid uint64, opts *UserBookmarksIllustOptions, sleepMs ...int) ([]models.Illust, error) {...}

FetchAllIllustFollows(opts *IllustFollowOptions, sleepMs ...int) ([]models.Illust, error) {...}

// 获取指定用户跟踪的用户列表
FetchAllUserFollowing(uid uint64, opts *UserFollowingOptions, sleepMs ...int) ([]models.UserPreview, error) {...}

FetchAllUserFollowers(uid uint64, opts *UserFollowerOptions, sleepMs ...int) ([]models.UserPreview, error) {...}

UserFollowAddMultiple(uids []uint64, restrict ...models.Restrict) ([]uint64, error) {...}

UserFollowDeleteMultiple(uids []uint64) ([]uint64, error) {...}
```

## Link

- [pixivpy](https://github.com/upbit/pixivpy/tree/master)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
