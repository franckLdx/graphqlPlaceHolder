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

func FetchAlbum(albumId int) (*Album, error) {
	var album Album
	err := FetchResource(AlbumResource, albumId, &album)
	return &album, err
}

func FetchPhotosOfAlbum(albumId int) (*[]Photo, error) {
	var photos []Photo
	err := FetchSubResources(AlbumResource, albumId, PhotoResource, &photos)
	return &photos, err
}
