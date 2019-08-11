package httpClient

import "fmt"

type Photo struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
	UserID       int    `json:"userId"`
	AlbumId      int    `json:"albumId"`
}

const photosFilter = "photos"

func FetchPhotos() (*[]Photo, error) {
	var photos []Photo
	if err := fetch(photosFilter, &photos); err != nil {
		return nil, fmt.Errorf("Failed to get photos list", err)
	}
	return &photos, nil
}

func FetchPhoto(id int) (*Photo, error) {
	filter := fmt.Sprintf("%s/%d", photosFilter, id)
	var photo Photo
	if err := fetch(filter, &photo); err != nil {
		return nil, fmt.Errorf("Failed to get photo %d: %v", id, err)
	}
	return &photo, nil
}
