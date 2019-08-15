package httpClient

type Photo struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
	UserID       int    `json:"userId"`
	AlbumId      int    `json:"albumId"`
}

const PhotoResource Resource = "photos"

func FetchPhotos() (*[]Photo, error) {
	var photos []Photo
	err := FetchResources(PhotoResource, &photos)
	return &photos, err
}

func FetchPhoto(id int) (*Photo, error) {
	var photo Photo
	err := FetchResource(AlbumResource, id, &photo)
	return &photo, err
}
