# Changelog

## [0.15.0](https://github.com/ryohidaka/go-pixiv/compare/v0.14.3...v0.15.0) (2025-05-11)


### Features

* **api:** Changed the method of obtaining refresh tokens. ([cb14c52](https://github.com/ryohidaka/go-pixiv/commit/cb14c5278cfebc919c06512a6d339636269f6b7a))

## [0.14.3](https://github.com/ryohidaka/go-pixiv/compare/v0.14.2...v0.14.3) (2025-05-08)


### Bug Fixes

* **api:** Change log level. ([1e85f3e](https://github.com/ryohidaka/go-pixiv/commit/1e85f3e1d82e450361c4fc25887729aa6cd130bc))
* **models:** Fixed Illust.Caption to optional ([cb5e4e2](https://github.com/ryohidaka/go-pixiv/commit/cb5e4e26e63ec670a3dcbd86b9f300eba4f3cd35))

## [0.14.2](https://github.com/ryohidaka/go-pixiv/compare/v0.14.1...v0.14.2) (2025-05-08)


### Bug Fixes

* **models:** Add missing fields for Illust struct. ([dc8611d](https://github.com/ryohidaka/go-pixiv/commit/dc8611de6c4d62009c4d67e5dcd13c9cf83e8d04))
* **models:** Add request fields for Illust struct. ([31a56b1](https://github.com/ryohidaka/go-pixiv/commit/31a56b18a0e668293f258abd56de317d0580e2ad))
* **models:** Fixed Illust.Type types to IllustType ([54800c1](https://github.com/ryohidaka/go-pixiv/commit/54800c1faaeab415c3b59ae528daf2f0590a31ff))
* **models:** Fixed Tag.TranslatedName types to nullable ([f5826e2](https://github.com/ryohidaka/go-pixiv/commit/f5826e2d8f073cbda579276d23c62d022402ac90))

## [0.14.1](https://github.com/ryohidaka/go-pixiv/compare/v0.14.0...v0.14.1) (2025-05-08)


### Bug Fixes

* **models:** Add missing fields for User struct. ([39e2cae](https://github.com/ryohidaka/go-pixiv/commit/39e2caeef018367c77ad06b952cd7e755a6c15fb))
* **models:** Modified Comment in User struct to be optional. ([c3f3c1a](https://github.com/ryohidaka/go-pixiv/commit/c3f3c1a9530a096a883a6e4b6f686eaa278e9c8c))

## [0.14.0](https://github.com/ryohidaka/go-pixiv/compare/v0.13.0...v0.14.0) (2025-05-06)


### Features

* **api:** Add FetchAllUserFollowers() ([7f47db5](https://github.com/ryohidaka/go-pixiv/commit/7f47db5077590b9749f0aa2ae602270c654d7fcb))
* **api:** Modified userID to uid ([b08d176](https://github.com/ryohidaka/go-pixiv/commit/b08d176923c9b17553506beacfe986c9ad00c81f))

## [0.13.0](https://github.com/ryohidaka/go-pixiv/compare/v0.12.0...v0.13.0) (2025-05-06)


### Features

* **api:** Add FetchAllUserFollowing() ([a856d0e](https://github.com/ryohidaka/go-pixiv/commit/a856d0effc9806d4baad192cac14da71b7a2600c))

## [0.12.0](https://github.com/ryohidaka/go-pixiv/compare/v0.11.0...v0.12.0) (2025-05-06)


### Features

* **api:** Add FetchAllIllustFollows() ([f8f0294](https://github.com/ryohidaka/go-pixiv/commit/f8f0294a0991592cbfd2e3de3407e85f6fedce98))

## [0.11.0](https://github.com/ryohidaka/go-pixiv/compare/v0.10.0...v0.11.0) (2025-05-04)


### Features

* **api:** Add FetchAllBookmarkedIllusts() ([e327c71](https://github.com/ryohidaka/go-pixiv/commit/e327c71a50f3ff88e10c55bd12aaf58a4ae4c54e))

## [0.10.0](https://github.com/ryohidaka/go-pixiv/compare/v0.9.0...v0.10.0) (2025-05-04)


### Features

* **api:** Add FetchAllUserIllusts() ([f876a9c](https://github.com/ryohidaka/go-pixiv/commit/f876a9c7062336e363b3806699a0c92fd5199cdb))

## [0.9.0](https://github.com/ryohidaka/go-pixiv/compare/v0.8.0...v0.9.0) (2025-05-04)


### Features

* **api:** Add UserFollower() ([ad0c8e5](https://github.com/ryohidaka/go-pixiv/commit/ad0c8e55d56490091820eac6be85838188216c98))


### Bug Fixes

* **api:** Modified type name ([a24f6f0](https://github.com/ryohidaka/go-pixiv/commit/a24f6f0798503fadec1f6571aa686dd5369bb967))

## [0.8.0](https://github.com/ryohidaka/go-pixiv/compare/v0.7.1...v0.8.0) (2025-05-04)


### Features

* **api:** Add UserFollowing() ([2e3cf7b](https://github.com/ryohidaka/go-pixiv/commit/2e3cf7b6eaddf2d0b948e1c9c4de1c86b433cbd2))

## [0.7.1](https://github.com/ryohidaka/go-pixiv/compare/v0.7.0...v0.7.1) (2025-05-04)


### Bug Fixes

* **model:** Remove unused model. ([c9c8971](https://github.com/ryohidaka/go-pixiv/commit/c9c897176efcdc92ab0e00e1f456a632d5863327))

## [0.7.0](https://github.com/ryohidaka/go-pixiv/compare/v0.6.0...v0.7.0) (2025-05-04)


### Features

* **api:** Add IllustBookmarkDetail() ([3c11009](https://github.com/ryohidaka/go-pixiv/commit/3c110098c7ed4b5f2ccba6967d4cab931e31c4c5))
* **api:** Add IllustBookmarkDetail() ([85f8cbf](https://github.com/ryohidaka/go-pixiv/commit/85f8cbff25b9c7939814d9a6fe3d00b176ee36d1))

## [0.6.0](https://github.com/ryohidaka/go-pixiv/compare/v0.5.1...v0.6.0) (2025-05-03)


### Features

* **api:** Added log output with slog ([711c265](https://github.com/ryohidaka/go-pixiv/commit/711c2653167382876576e12eb43c02ef458b1194))

## [0.5.1](https://github.com/ryohidaka/go-pixiv/compare/v0.5.0...v0.5.1) (2025-05-03)


### Bug Fixes

* Rename user_bookmarks_illust.go ([6af1c49](https://github.com/ryohidaka/go-pixiv/commit/6af1c4999d979e286741791ea2c9e5efc40845aa))

## [0.5.0](https://github.com/ryohidaka/go-pixiv/compare/v0.4.0...v0.5.0) (2025-05-03)


### Features

* **api:** Add IllustDetail() ([1c7a32b](https://github.com/ryohidaka/go-pixiv/commit/1c7a32b6577695f23723e3d59f95c82b75b54b0f))
* **api:** Add IllustDetail() ([e917ee5](https://github.com/ryohidaka/go-pixiv/commit/e917ee5b79fe9e3bdd101dc04a85509f66abf25f))

## [0.4.0](https://github.com/ryohidaka/go-pixiv/compare/v0.3.1...v0.4.0) (2025-05-03)


### Features

* **api:** Add IllustFollow() ([d71e454](https://github.com/ryohidaka/go-pixiv/commit/d71e454a5b343897eaca9e27747efd30f3f0e393))

## [0.3.1](https://github.com/ryohidaka/go-pixiv/compare/v0.3.0...v0.3.1) (2025-05-03)


### Bug Fixes

* Fix type name ([1d08fb9](https://github.com/ryohidaka/go-pixiv/commit/1d08fb920707f1e8ced7be53c64edef41b426f93))
* **model:** Add Restrict type. ([085a0f1](https://github.com/ryohidaka/go-pixiv/commit/085a0f138e0915dac5c498eb7b6b44128c4b0ef5))

## [0.3.0](https://github.com/ryohidaka/go-pixiv/compare/v0.2.0...v0.3.0) (2025-05-03)


### Features

* **api:** Add UserBookmarksIllust() ([2a430c8](https://github.com/ryohidaka/go-pixiv/commit/2a430c8dbe1a51e2c8c1cac1b3f9aa136a180366))
* **utils:** Add getRestrict() ([527c991](https://github.com/ryohidaka/go-pixiv/commit/527c9913c2a53a49252cce1b99f100cc20ceb8b9))

## [0.2.0](https://github.com/ryohidaka/go-pixiv/compare/v0.1.0...v0.2.0) (2025-05-03)


### Features

* **api:** Add UserIllusts() ([6b5dc86](https://github.com/ryohidaka/go-pixiv/commit/6b5dc869ad06ebd9dfe4c194deb4a4ca02b0ccdb))
* **utils:** Add parseNextPageOffset() ([3efaa4f](https://github.com/ryohidaka/go-pixiv/commit/3efaa4f1d212fc9d8da6b431e738014c2ab09865))

## 0.1.0 (2025-05-02)


### Features

* **api:** Add API client. ([86e61d4](https://github.com/ryohidaka/go-pixiv/commit/86e61d4d4afae8f38f24eebff0949dbcb40e635a))
* **api:** Add UserDetail() ([10222be](https://github.com/ryohidaka/go-pixiv/commit/10222bea2036709e32ff80dc43879b2c72b7088b))
* **auth:** Added API authentication process. ([0cbfeff](https://github.com/ryohidaka/go-pixiv/commit/0cbfeff8003a7da57815b038b2ee202a61894074))
* **testutil:** Add process to get refresh-token from environments. ([4a1114f](https://github.com/ryohidaka/go-pixiv/commit/4a1114f9bd56fbd148333a103473c6d6d258c672))


### Bug Fixes

* Change module name. ([f9d2e0a](https://github.com/ryohidaka/go-pixiv/commit/f9d2e0a88015fa19b8eafe7063879c0c1a536516))


### Miscellaneous Chores

* release 0.1.0 ([37d9d37](https://github.com/ryohidaka/go-pixiv/commit/37d9d37ee3a2d9390c931d3e96e22362f241df34))
