package entity

import "time"

type PhotoAlbumMap map[string]PhotoAlbum

func (m PhotoAlbumMap) Get(name, photoUUID, albumUUID string) PhotoAlbum {
	if result, ok := m[name]; ok {
		return result
	}

	return *NewPhotoAlbum(photoUUID, albumUUID)
}

func (m PhotoAlbumMap) Pointer(name, photoUUID, albumUUID string) *PhotoAlbum {
	if result, ok := m[name]; ok {
		return &result
	}

	return NewPhotoAlbum(photoUUID, albumUUID)
}

var PhotoAlbumFixtures = PhotoAlbumMap{
	"1": {
		PhotoUUID: "pt9jtdre2lvl0yh7",
		AlbumUUID: "at9lxuqxpogaaba8",
		Order:     0,
		CreatedAt: time.Date(2020, 3, 6, 2, 6, 51, 0, time.UTC),
		UpdatedAt: time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		Photo:     PhotoFixtures.Pointer("19800101_000002_D640C559"),
		Album:     AlbumFixtures.Pointer("holiday-2030"),
	},
	"2": {
		PhotoUUID: "pt9jtdre2lvl0y11",
		AlbumUUID: "at9lxuqxpogaaba9",
		Order:     0,
		CreatedAt: time.Date(2020, 2, 6, 2, 6, 51, 0, time.UTC),
		UpdatedAt: time.Date(2020, 4, 28, 14, 6, 0, 0, time.UTC),
		Photo:     PhotoFixtures.Pointer("Photo04"),
		Album:     AlbumFixtures.Pointer("berlin-2019"),
	},
	"3": {
		PhotoUUID: "pt9jtdre2lvl0yh8",
		AlbumUUID: "at9lxuqxpogaaba9",
		Order:     0,
		CreatedAt: time.Date(2020, 2, 6, 2, 6, 51, 0, time.UTC),
		UpdatedAt: time.Date(2020, 4, 28, 14, 6, 0, 0, time.UTC),
		Photo:     PhotoFixtures.Pointer("Photo01"),
		Album:     AlbumFixtures.Pointer("berlin-2019"),
	},
}

// CreatePhotoAlbumFixtures inserts known entities into the database for testing.
func CreatePhotoAlbumFixtures() {
	for _, entity := range PhotoAlbumFixtures {
		Db().Create(&entity)
	}
}
