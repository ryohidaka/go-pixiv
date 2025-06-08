package webapi_test

import (
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/pkg/webapi"
	"github.com/ryohidaka/go-pixiv/testutil/webtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserLatestWorks(t *testing.T) {
	webtest.WithMockHTTP(t, func() {
		// Mock the user following response
		url := webapi.ApiHosts + "user/11/works/latest"
		err := webtest.MockResponseFromFile("GET", url, "user_works")
		assert.NoError(t, err)

		// Initialize the WebPixivAPI instance
		app, err := pixiv.NewWebApp("dummy-phpsessid")
		assert.NoError(t, err)

		// Call the UserLatestWorks method
		works, err := app.UserLatestWorks(11)
		assert.NoError(t, err)
		assert.NotNil(t, works)

		// Expected values (adjust according to the user_works.json mock file)
		// Expected illusts
		expectedIllusts := map[string]interface{}{
			"1580459":   nil,
			"130967689": "comic POOL5月コミックス発売情報",
			"131093702": "マンガ投稿企画『ツキコミ〜裏を見ちゃった！〜』開催",
			"131202453": "pixivが大阪・関西万博で期間限定展示を開催",
		}

		for id, expected := range expectedIllusts {
			illust, exists := works.Illusts[id]
			assert.True(t, exists, "illust %s should exist", id)
			if expected == nil {
				assert.Nil(t, illust, "illust %s should be nil", id)
			} else {
				require.NotNil(t, illust, "illust %s should not be nil", id)
				assert.Equal(t, expected, illust.Title)
			}
		}

		// Expected novels
		expectedNovels := map[string]interface{}{
			"129":      nil,
			"24757014": "小説企画「執筆応援プロジェクト～毒～」開催 ",
			"24874304": "6月投稿企画「RainRainRain2025」開催",
			"24890505": "「第18回らぶドロップス恋愛小説コンテスト」開催",
		}

		for id, expected := range expectedNovels {
			novel, exists := works.Novels[id]
			assert.True(t, exists, "novel %s should exist", id)
			if expected == nil {
				assert.Nil(t, novel, "novel %s should be nil", id)
			} else {
				require.NotNil(t, novel, "novel %s should not be nil", id)
				assert.Equal(t, expected, novel.Title)
			}
		}
	})
}
