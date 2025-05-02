# go-pixiv

Pixiv API for Golang

Inspired by [pixivpy](https://github.com/upbit/pixivpy/tree/master)

> [!IMPORTANT]
> Only authentication with a refresh token is supported.
>
> Please check [pixivpy's README](https://github.com/upbit/pixivpy/tree/master) for instructions on how to obtain a refresh token.

## API functions

### App-API (6.0 - app-api.pixiv.net)

```go
// 用户详情
UserDetail(uid uint64, opts *UserDetailOptions) (*models.UserDetail, error) {...}
```

## Usage

```go
// Create a new Pixiv App API client
app, _ := pixiv.NewApp("<YOUR_REFRESH_TOKEN>")

// Fetch user details
user, _ := app.UserDetail(11)
fmt.Println("Name:", user.User.Name)
// Outputs: pixiv事務局
```

## Link

- [pixivpy](https://github.com/upbit/pixivpy/tree/master)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
