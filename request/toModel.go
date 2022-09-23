package request

import "int-gateway/models"

func (s *ShowLength) ToModel() *models.ShowLength {
	return &models.ShowLength{
		Hours:   s.Hours,
		Minutes: s.Minutes,
	}
}

func (g *ShortGenres) ToModel() models.ShortGenres {
	genres := models.ShortGenres{}
	for _, genre := range *g {
		genres = append(genres, genre.ToModel())
	}
	return genres
}

func (g *ShortGenre) ToModel() *models.ShortGenre {
	return &models.ShortGenre{
		ID:   g.ID,
		Name: g.Name,
	}
}

func (s *ShortCelebrities) ToModel() models.ShortCelebrities {
	celebs := models.ShortCelebrities{}
	for _, celeb := range *s {
		celebs = append(celebs, celeb.ToModel())
	}
	return celebs
}

func (s *ShortCelebrity) ToModel() *models.ShortCelebrity {
	return &models.ShortCelebrity{
		ID:          s.ID,
		Name:        s.Name,
		RoleName:    s.RoleName,
		PostersPath: s.PostersPath,
	}
}

func (f *FilmCrews) ToModel() models.FilmCrews {
	crews := models.FilmCrews{}
	for _, crew := range *f {
		crews = append(crews, crew.ToModel())
	}
	return crews
}

func (f *FilmCrew) ToModel() *models.FilmCrew {
	return &models.FilmCrew{
		ID:          f.ID,
		Name:        f.Name,
		PostersPath: f.PostersPath,
	}
}

func (e *ShortEpisode) ToModel() *models.ShortEpisode {
	return &models.ShortEpisode{
		ID:          e.ID,
		Title:       e.Title,
		PostersPath: e.PostersPath,
		Rating:      e.Rating,
		Resume:      e.Resume,
	}
}

func (se *ShortEpisodes) ToModel() models.ShortEpisodes {
	episodes := models.ShortEpisodes{}
	for _, ep := range *se {
		episodes = append(episodes, ep.ToModel())
	}
	return episodes
}

func (s *ShortSeason) ToModel() *models.ShortSeason {
	return &models.ShortSeason{
		ID:          s.ID,
		Title:       s.Title,
		PostersPath: s.PostersPath,
		Rating:      s.Rating,
	}
}

func (s *ShortSeasons) ToModel() models.ShortSeasons {
	seasons := models.ShortSeasons{}
	for _, season := range *s {
		seasons = append(seasons, season.ToModel())
	}
	return seasons
}
