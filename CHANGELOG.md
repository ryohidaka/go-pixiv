# Changelog

## [0.30.0](https://github.com/ryohidaka/go-pixiv/compare/v0.29.0...v0.30.0) (2025-06-08)


### Features

* **api:** Add UserBookmarksIllusts() ([853fae8](https://github.com/ryohidaka/go-pixiv/commit/853fae84f179f0b6742eb3adb750f94549c71022))

## [0.29.0](https://github.com/ryohidaka/go-pixiv/compare/v0.28.0...v0.29.0) (2025-06-08)


### Features

* **api:** Add UserLatestWorks() ([18c8f9e](https://github.com/ryohidaka/go-pixiv/commit/18c8f9e33d062abe966b97416b2b09a53530b1a7))

## [0.28.0](https://github.com/ryohidaka/go-pixiv/compare/v0.27.0...v0.28.0) (2025-06-08)


### Features

* **api:** Add UserFollowers() ([4c2c3e1](https://github.com/ryohidaka/go-pixiv/commit/4c2c3e13633e83a989baed1bf2f0736ba52fe35b))


### Bug Fixes

* **api:** Fix referers for UserFollowing() ([91349a5](https://github.com/ryohidaka/go-pixiv/commit/91349a513081d120d8da2a7bcb38a90b0550fd02))

## [0.27.0](https://github.com/ryohidaka/go-pixiv/compare/v0.26.0...v0.27.0) (2025-06-08)


### Features

* **api:** Add UserFollowing() ([09d0218](https://github.com/ryohidaka/go-pixiv/commit/09d0218e0a02bf069c8e7371519ed9edf8a2aaa2))

## [0.26.0](https://github.com/ryohidaka/go-pixiv/compare/v0.25.0...v0.26.0) (2025-06-06)


### Features

* **api:** Add UserProfile() ([167cd0d](https://github.com/ryohidaka/go-pixiv/commit/167cd0d607c778375b13e697592bfd5ad9f95083))


### Bug Fixes

* **api:** Expose WebAPI methods in wapi.go ([889abaa](https://github.com/ryohidaka/go-pixiv/commit/889abaaba17fcd6710f586c052b4cce0161c810d))

## [0.25.0](https://github.com/ryohidaka/go-pixiv/compare/v0.24.0...v0.25.0) (2025-06-06)


### Features

* **api:** Add doc link. ([54ec650](https://github.com/ryohidaka/go-pixiv/commit/54ec65063860d279135de64eb8a3d35ade22dea3))
* **api:** Add UserFull() ([c5cd380](https://github.com/ryohidaka/go-pixiv/commit/c5cd3807e6819c298506a5f7f5ff6e335f304fe1))

## [0.24.0](https://github.com/ryohidaka/go-pixiv/compare/v0.23.1...v0.24.0) (2025-06-05)


### Features

* **api:** Add UserShort() ([018be36](https://github.com/ryohidaka/go-pixiv/commit/018be3647f1451f8faef86eb4d4ac1c75cb0622f))
* **api:** Add WebAPI client. ([f1482b3](https://github.com/ryohidaka/go-pixiv/commit/f1482b3308267f91759f0bd4ac2e4db9048c7837))

## [0.23.1](https://github.com/ryohidaka/go-pixiv/compare/v0.23.0...v0.23.1) (2025-06-05)


### Bug Fixes

* **crawler:** Change NewCrawler loading destination of test to the root of the package. ([2416761](https://github.com/ryohidaka/go-pixiv/commit/24167617ced0267cfc4c85cf2a14655c27632418))

## [0.23.0](https://github.com/ryohidaka/go-pixiv/compare/v0.22.0...v0.23.0) (2025-06-05)


### Features

* **api:** Changed AppPixivAPI method caller to the root of the library. ([2568de4](https://github.com/ryohidaka/go-pixiv/commit/2568de42de7320c136e65243ffbeb81efa30388d))
* **client:** Changed Get and Post functions from AppPixivAPI methods to normal functions. ([42c64e7](https://github.com/ryohidaka/go-pixiv/commit/42c64e71a01b8b269a0edefe53088ab0e3e0a63d))
* **testutil:** Embed fixtures using go:embed ([16b1021](https://github.com/ryohidaka/go-pixiv/commit/16b1021791753d878d3374577da34f9008f024d4))
* **testutil:** Move testutil functions to the appapi directory. ([55aa2ae](https://github.com/ryohidaka/go-pixiv/commit/55aa2ae41c13d7b34b82ffb1e36932bbdc055e51))

## [0.22.0](https://github.com/ryohidaka/go-pixiv/compare/v0.21.0...v0.22.0) (2025-06-05)


### Features

* **api:** Move AppApi functions to the pkg directory. ([ad3c66c](https://github.com/ryohidaka/go-pixiv/commit/ad3c66c735ee256d427fddd2236b02cb7c9cc185))
* **client:** Move api.go to pkg directory. ([240fd2c](https://github.com/ryohidaka/go-pixiv/commit/240fd2c1f3e0c6e98f274b648b46b34200310d37))
* **crawler:** Move crawler functions into pkg directory. ([006805b](https://github.com/ryohidaka/go-pixiv/commit/006805b4e2aab802b2b78c55ff8301a6e1477722))
* **downloader:** Move downloader functions into pkg directory. ([d996c4d](https://github.com/ryohidaka/go-pixiv/commit/d996c4decfd1aaf0aad86861977cda30492b8f15))

## [0.21.0](https://github.com/ryohidaka/go-pixiv/compare/v0.20.2...v0.21.0) (2025-06-03)


### Features

* **client:** Add download client. ([6dfeabd](https://github.com/ryohidaka/go-pixiv/commit/6dfeabdee5db322018b93baa4ecb7fa196126f60))
* **downloader:** Add DownloadBytes(). ([d838617](https://github.com/ryohidaka/go-pixiv/commit/d838617a90943a876a5b70c4def7f3b8d8147b2f))
* **downloader:** Add DownloadFile(). ([aa637eb](https://github.com/ryohidaka/go-pixiv/commit/aa637ebf060f5fcdb0bf81c45c4153634a9c95d8))

## [0.20.2](https://github.com/ryohidaka/go-pixiv/compare/v0.20.1...v0.20.2) (2025-06-03)


### Bug Fixes

* **api:** Remove crawler logs. ([6fa06d1](https://github.com/ryohidaka/go-pixiv/commit/6fa06d149aebbde5bc9a4467b67f94a515978bfc))
* **api:** Remove debug logs. ([ac13a6d](https://github.com/ryohidaka/go-pixiv/commit/ac13a6d3335bb40c55036ff308bc7b96338ed691))
* **api:** Remove error logs. ([ae4df8d](https://github.com/ryohidaka/go-pixiv/commit/ae4df8d2ccf2fa6b0a80ef62b6d23524a7c42209))
* **api:** Remove info logs. ([63c711a](https://github.com/ryohidaka/go-pixiv/commit/63c711a91e9b5fb160a417e086b2f6f7cbeb1179))

## [0.20.1](https://github.com/ryohidaka/go-pixiv/compare/v0.20.0...v0.20.1) (2025-05-16)


### Bug Fixes

* **api:** Add sleep processing for each request ([98587da](https://github.com/ryohidaka/go-pixiv/commit/98587da922000fb87f5c0dac8d5d5140aed4067d))

## [0.20.0](https://github.com/ryohidaka/go-pixiv/compare/v0.19.0...v0.20.0) (2025-05-16)


### Features

* **api:** Add UserFollowAddMultiple() ([2f38acd](https://github.com/ryohidaka/go-pixiv/commit/2f38acd05aa7f973f6ef439d37e25525613d4b33))
* **api:** Add UserFollowDeleteMultiple() ([4e651fb](https://github.com/ryohidaka/go-pixiv/commit/4e651fbb71a88b2002df9dfd763e0a1bf464b86a))

## [0.19.0](https://github.com/ryohidaka/go-pixiv/compare/v0.18.0...v0.19.0) (2025-05-16)


### Features

* **api:** Add UserFollowAdd() ([5c8b745](https://github.com/ryohidaka/go-pixiv/commit/5c8b74580702faf7d2c7122670e471512c5b0974))
* **api:** Add UserFollowDelete() ([14852d3](https://github.com/ryohidaka/go-pixiv/commit/14852d317d04a8dec7c228397c8582e3e44ef8a1))


### Bug Fixes

* **api:** Support for empty API responses ([6c444ad](https://github.com/ryohidaka/go-pixiv/commit/6c444ade919a8dc8d57b9089584bf393ce866a6d))

## [0.18.0](https://github.com/ryohidaka/go-pixiv/compare/v0.17.0...v0.18.0) (2025-05-16)


### Features

* **api:** Add wrapper method to make GET requests ([211900a](https://github.com/ryohidaka/go-pixiv/commit/211900a786ff69c8e6ad2885f9b73a255716d2b1))
* **api:** Add wrapper method to make POST requests ([9fcba9f](https://github.com/ryohidaka/go-pixiv/commit/9fcba9fb2d7ea7a2f8073a35375141dadfdb91a8))
* **api:** Support for HTTP requests other than GET ([f6cbb07](https://github.com/ryohidaka/go-pixiv/commit/f6cbb07620e75ff8e86375e01d217d9369b19a73))

## [0.17.0](https://github.com/ryohidaka/go-pixiv/compare/v0.16.0...v0.17.0) (2025-05-15)


### Features

* **crawler:** Add NewCrawler() ([8c1e2a2](https://github.com/ryohidaka/go-pixiv/commit/8c1e2a202b89b671d2f1bf6f27caebb7dbd909a9))
* **crawler:** Move FetchAllBookmarkedIllusts() to crawler ([f855374](https://github.com/ryohidaka/go-pixiv/commit/f8553749b2850b66cb65e67c5852f2f1ad340a87))
* **crawler:** Move FetchAllIllustFollows() to crawler ([cb17c4d](https://github.com/ryohidaka/go-pixiv/commit/cb17c4d3fbd32d635e8c66539648e51af32132a0))
* **crawler:** Move FetchAllUserFollowers() to crawler ([988b413](https://github.com/ryohidaka/go-pixiv/commit/988b413cc02e6d8c1f986a22a6924106e1062bbb))
* **crawler:** Move FetchAllUserFollowing() to crawler ([4e9a96a](https://github.com/ryohidaka/go-pixiv/commit/4e9a96a3e3f3c670df793eb2dc80f015d03cc366))
* **crawler:** Move FetchAllUserIllusts() to crawler ([c50f7c2](https://github.com/ryohidaka/go-pixiv/commit/c50f7c29db8114df264f6af14db8ffea5a98a4fb))
* **crawler:** Move getSleepDuration() to crawler ([c70e540](https://github.com/ryohidaka/go-pixiv/commit/c70e540a3f33660dd13b18bd4c6c7a4eeaff92ef))

## [0.16.0](https://github.com/ryohidaka/go-pixiv/compare/v0.15.0...v0.16.0) (2025-05-15)


### Features

* **api:** Change query parameter to variable argument ([426c73a](https://github.com/ryohidaka/go-pixiv/commit/426c73a3ebe8fa17fe0ae966e080d07099304c90))

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
