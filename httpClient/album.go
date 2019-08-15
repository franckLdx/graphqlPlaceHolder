package httpClient

type Album struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	UserID int    `json:"userId"`
}

const AlbumResource Resource = "albums"

func FetchAlbums() (*[]Album, error) {
	var albums []Album
	err := FetchResources(AlbumResource, &albums)
	return &albums, err
}

func FetchAlbum(id int) (*Album, error) {
	var album Album
	err := FetchResource(AlbumResource, id, &album)
	return &album, err
}
