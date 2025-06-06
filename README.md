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

#### App-API (6.0 - app-api.pixiv.net)

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
users, next, err := app.UserFollowing(11)

// Fetch user follower
users, next, err := app.UserFollower(11)

// Follow user with user ID
ok, err := app.UserFollowAdd(11)

// Unfollow user with user ID
ok, err := app.UserFollowDelete(11)
```

#### Web-API (ajax)

```go
import "github.com/ryohidaka/go-pixiv"

// Create a new Pixiv Web API client
app, err := pixiv.NewWebApp("<YOUR_PHPSESSID>")

// Fetch a user information in a simplified format.
user, err := app.UserShort(11)

// Fetch a full user information.
user, err := app.UserFull(11)
```

### Crawler

#### App-API (6.0 - app-api.pixiv.net)

```go
import "github.com/ryohidaka/go-pixiv"

// Create a new Pixiv Crawler
crawler, err := pixiv.NewCrawler(refreshToken)

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

### Downloader

```go
import "github.com/ryohidaka/go-pixiv"

downloader := pixiv.NewDownloader()

// Download as byte
data, err := downloader.DownloadBytes("https://i.pximg.net/...")

// Download as file
len, err := downloader.DownloadFile("https://i.pximg.net/...", &pixiv.DownloadFileOptions{
    Dir:     ".tmp",
    Name:    "test.jpg",
    Replace: true,
})
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

### Crawler

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

### Downloader

```go
DownloadBytes(url string) ([]byte, error) {...}

DownloadFile(url string, opts *DownloadFileOptions) (int64, error) {...}
```

### Web API

```go
UserShort(uid uint64) (*webmodel.UserShort, error) {...}

UserFull(uid uint64) (*user.User, error) {...}
```

## Link

- [pixivpy](https://github.com/upbit/pixivpy/tree/master)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
