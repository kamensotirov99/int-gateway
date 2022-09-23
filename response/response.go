package response

import (
	"time"
)

type Gender string

type UploadImages struct {
	PostersPath []string `json:"postersPath"`
}

type Journalist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Journalists []*Journalist

type ShortJournalist struct {
	ID string `json:"id"`
}

type Celebrities []*Celebrity

type Celebrity struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Occupation   []string  `json:"occupation"`
	PostersPath  []string  `json:"postersPath"`
	DateOfBirth  time.Time `json:"dateOfBirth"`
	DateOfDeath  time.Time `json:"dateOfDeath"`
	PlaceOfBirth string    `json:"placeOfBirth"`
	Gender       Gender    `json:"gender"`
	Bio          string    `json:"bio"`
}

type Season struct {
	ID          string        `json:"id"`
	ShowID      string        `json:"showId"`
	Title       string        `json:"title"`
	TrailerURL  string        `json:"trailerUrl"`
	PostersPath []string      `json:"postersPath"`
	Resume      string        `json:"resume"`
	Rating      float64       `json:"rating"`
	ReleaseDate time.Time     `json:"releaseDate"`
	WrittenBy   FilmCrews     `json:"writtenBy"`
	ProducedBy  FilmCrews     `json:"producedBy"`
	DirectedBy  FilmCrews     `json:"directedBy"`
	Episodes    ShortEpisodes `json:"episodes"`
}

type FilmCrews []*FilmCrew

type FilmCrew struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	PostersPath []string `json:"postersPath"`
}

type Seasons []*Season

type ShortEpisode struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	PostersPath []string `json:"postersPath"`
	Rating      float64  `json:"rating"`
	Resume      string   `json:"resume"`
}

type ShortEpisodes []*ShortEpisode
type Genres []*Genre

type Genre struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Article struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	ReleaseDate time.Time       `json:"releaseDate"`
	PostersPath []string        `json:"postersPath"`
	Description string          `json:"description"`
	Journalist  ShortJournalist `json:"journalist"`
}

type Articles []*Article

type Shows []*Show

type Show struct {
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	Type        string           `json:"type"`
	PostersPath []string         `json:"postersPath"`
	ReleaseDate time.Time        `json:"releaseDate"`
	EndDate     time.Time        `json:"endDate"`
	Rating      float64          `json:"rating"`
	Length      ShowLength       `json:"length"`
	TrailerURL  string           `json:"trailerUrl"`
	Genres      ShortGenres      `json:"genres"`
	DirectedBy  FilmCrews        `json:"directedBy"`
	ProducedBy  FilmCrews        `json:"producedBy"`
	WrittenBy   FilmCrews        `json:"writtenBy"`
	Starring    ShortCelebrities `json:"starring"`
	Description string           `json:"description"`
	Seasons     ShortSeasons     `json:"seasons"`
}

type Episodes []*Episode

type Episode struct {
	ID          string           `json:"id"`
	SeasonID    string           `json:"seasonId"`
	Title       string           `json:"title"`
	PostersPath []string         `json:"postersPath"`
	TrailerURL  string           `json:"trailerUrl"`
	Length      ShowLength       `json:"length"`
	Rating      float64          `json:"rating"`
	Resume      string           `json:"resume"`
	WrittenBy   FilmCrews        `json:"writtenBy"`
	ProducedBy  FilmCrews        `json:"producedBy"`
	DirectedBy  FilmCrews        `json:"directedBy"`
	Starring    ShortCelebrities `json:"starring"`
}

type ShortCelebrities []*ShortCelebrity

type ShortCelebrity struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	RoleName    string   `json:"roleName"`
	PostersPath []string `json:"postersPath"`
}

type ShowLength struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

type ShortSeasons []*ShortSeason

type ShortSeason struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	PostersPath []string `json:"postersPath"`
	Rating      float64  `json:"rating"`
}

type ShortGenres []*ShortGenre

type ShortGenre struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
