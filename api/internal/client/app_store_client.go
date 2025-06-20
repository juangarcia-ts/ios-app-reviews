package client

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"strconv"
)

type AppStoreReview struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
	Rating int `json:"rating"`
	SubmittedAt string `json:"submittedAt"`
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
	ID struct {
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
		Author struct{}    `json:"author"`
		Entry  []RssFeedEntry   `json:"entry"`
	} `json:"feed"`
}

type AppStoreClient struct {}

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
		appStoreReviews[i] = AppStoreReview{
			Title: review.Title.Label,
			Content: review.Content.Label,
			Author: review.Author.Name.Label,
			Rating: rating,
			SubmittedAt: review.Updated.Label,
		}
	}

	return appStoreReviews, nil
}