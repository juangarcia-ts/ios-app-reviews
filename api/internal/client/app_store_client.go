package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"ios-app-reviews-viewer.com/m/internal/helpers"
)

type AppStoreReview struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	Rating      int       `json:"rating"`
	SubmittedAt time.Time `json:"submittedAt"`
}

type RssFeedEntry struct {
	Author struct {
		URI struct {
			Label string `json:"label"`
		} `json:"uri"`
		Name struct {
			Label string `json:"label"`
		} `json:"name"`
		Label string `json:"label"`
	} `json:"author"`
	Updated struct {
		Label string `json:"label"`
	} `json:"updated"`
	IMRating struct {
		Label string `json:"label"`
	} `json:"im:rating"`
	IMVersion struct {
		Label string `json:"label"`
	} `json:"im:version"`
	Id struct {
		Label string `json:"label"`
	} `json:"id"`
	Title struct {
		Label string `json:"label"`
	} `json:"title"`
	Content struct {
		Label      string `json:"label"`
		Attributes struct {
			Type string `json:"type"`
		} `json:"attributes"`
	} `json:"content"`
	Link struct {
		Attributes struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"attributes"`
	} `json:"link"`
	IMVoteSum struct {
		Label string `json:"label"`
	} `json:"im:voteSum"`
	IMContentType struct {
		Attributes struct {
			Term  string `json:"term"`
			Label string `json:"label"`
		} `json:"attributes"`
	} `json:"im:contentType"`
	IMVoteCount struct {
		Label string `json:"label"`
	} `json:"im:voteCount"`
}

type RssFeed struct {
	Feed struct {
		Author struct{}       `json:"author"`
		Entry  []RssFeedEntry `json:"entry"`
	} `json:"feed"`
}

type GetAppInfoFromAppStore struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		TrackName     string `json:"trackName"`
		ArtworkUrl512 string `json:"artworkUrl512"`
	} `json:"results"`
}

type AppInfo struct {
	AppName string `json:"app_name"`
	LogoUrl string `json:"logo_url"`
}

type AppStoreClient struct{}

func NewAppStoreClient() *AppStoreClient {
	return &AppStoreClient{}
}

func (c *AppStoreClient) GetReviews(appId string, page int) ([]AppStoreReview, error) {
	rssUrl := fmt.Sprintf("https://itunes.apple.com/us/rss/customerreviews/id=%s/sortBy=mostRecent/page=%d/json", appId, page)
	resp, err := http.Get(rssUrl)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var rssFeed RssFeed
	json.Unmarshal(body, &rssFeed)

	reviews := rssFeed.Feed.Entry
	appStoreReviews := make([]AppStoreReview, len(reviews))

	for i, review := range reviews {
		rating, _ := strconv.Atoi(review.IMRating.Label)
		parsedSubmittedAt, _ := helpers.ParseDateTime(review.Updated.Label)

		appStoreReviews[i] = AppStoreReview{
			Id:          review.Id.Label,
			Title:       review.Title.Label,
			Content:     review.Content.Label,
			Author:      review.Author.Name.Label,
			Rating:      rating,
			SubmittedAt: parsedSubmittedAt,
		}
	}

	return appStoreReviews, nil
}

func (c *AppStoreClient) GetAppInfoFromAppStore(appId string) (*AppInfo, error) {
	url := fmt.Sprintf("https://itunes.apple.com/lookup?id=%s", appId)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response GetAppInfoFromAppStore

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if response.ResultCount == 0 || len(response.Results) == 0 {
		return nil, fmt.Errorf("app not found")
	}

	appInfo := response.Results[0]

	if err != nil {
		return nil, err
	}

	return &AppInfo{
		AppName: appInfo.TrackName,
		LogoUrl: appInfo.ArtworkUrl512,
	}, nil
}
