package models

import (
	"time"
)

type UploadResponse struct {
	PostersPath []string
}

type Celebrities []*Celebrity

type Gender string

type Celebrity struct {
	ID           string
	Name         string
	Occupation   []string
	PostersPath  []string
	DateOfBirth  time.Time
	DateOfDeath  time.Time
	PlaceOfBirth string
	Gender       Gender
	Bio          string
}

type Genres []*Genre

type Genre struct {
	ID          string
	Name        string
	Description string
}

type Journalist struct {
	ID   string
	Name string
}

type Journalists []*Journalist

type FilmCrews []*FilmCrew

type FilmCrew struct {
	ID          string
	Name        string
	PostersPath []string
}

type Season struct {
	ID          string
	ShowID      string
	Title       string
	TrailerURL  string
	PostersPath []string
	Resume      string
	Rating      float64
	ReleaseDate time.Time
	WrittenBy   FilmCrews
	ProducedBy  FilmCrews
	DirectedBy  FilmCrews
	Episodes    ShortEpisodes
}

type Seasons []*Season

type ShortEpisodes []*ShortEpisode

type ShortEpisode struct {
	ID          string
	Title       string
	PostersPath []string
	Rating      float64
	Resume      string
}

type Article struct {
	ID          string
	Title       string
	ReleaseDate time.Time
	PostersPath []string
	Description string
	Journalist  ShortJournalist
}

type Articles []*Article

type ShortJournalist struct {
	ID string
}

type Shows []*Show

type Show struct {
	ID          string
	Title       string
	Type        string
	PostersPath []string
	ReleaseDate time.Time
	EndDate     time.Time
	Rating      float64
	Length      ShowLength
	TrailerURL  string
	Genres      ShortGenres
	DirectedBy  FilmCrews
	ProducedBy  FilmCrews
	WrittenBy   FilmCrews
	Starring    ShortCelebrities
	Description string
	Seasons     ShortSeasons
}

type Episodes []*Episode

type Episode struct {
	ID          string
	SeasonID    string
	Title       string
	PostersPath []string
	TrailerURL  string
	Length      ShowLength
	Rating      float64
	Resume      string
	WrittenBy   FilmCrews
	ProducedBy  FilmCrews
	DirectedBy  FilmCrews
	Starring    ShortCelebrities
}

type ShowLength struct {
	Hours   int
	Minutes int
}

type ShortCelebrities []*ShortCelebrity

type ShortCelebrity struct {
	ID          string
	Name        string
	RoleName    string
	PostersPath []string
}

type ShortSeasons []*ShortSeason

type ShortSeason struct {
	ID          string
	Title       string
	PostersPath []string
	Rating      float64
}

type ShortGenres []*ShortGenre

type ShortGenre struct {
	ID   string
	Name string
}
