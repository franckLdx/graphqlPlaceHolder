package httpClient

import "fmt"

type Album struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	UserID int    `json:"userId"`
}

const albumsFilter = "albums"

func FetchAlbums() (*[]Album, error) {
	var albums []Album
	if err := fetch(albumsFilter, &albums); err != nil {
		return nil, fmt.Errorf("Failed to get albums list", err)
	}
	return &albums, nil
}

func FetchAlbum(id int) (*Album, error) {
	filter := fmt.Sprintf("%s/%d", albumsFilter, id)
	var album Album
	if err := fetch(filter, &album); err != nil {
		return nil, fmt.Errorf("Failed to get album %d: %v", id, err)
	}
	return &album, nil
}
