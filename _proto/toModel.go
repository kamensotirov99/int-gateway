package __

import "int-gateway/models"

func (u *UploadPostersResponse) ToModel() *models.UploadResponse {
	return &models.UploadResponse{
		PostersPath: u.PostersPath,
	}
}

func (c *Celebrity) ToModel() interface{} {
	return &models.Celebrity{
		ID:           c.Id,
		Name:         c.Name,
		Occupation:   c.Occupation,
		PostersPath:  c.PostersPath,
		DateOfBirth:  c.DateOfBirth.AsTime(),
		DateOfDeath:  c.DateOfDeath.AsTime(),
		PlaceOfBirth: c.PlaceOfBirth,
		Gender:       models.Gender(c.Gender),
		Bio:          c.Bio,
	}
}

func (c *CelebrityListResponse) ToModel() models.Celebrities {
	celebs := models.Celebrities{}
	for _, celeb := range c.Celebrities {
		celebs = append(celebs, celeb.ToModel().(*models.Celebrity))
	}
	return celebs
}

func (s *Season) ToModel() interface{} {
	return &models.Season{
		ID:          s.Id,
		ShowID:      s.ShowId,
		Title:       s.Title,
		TrailerURL:  s.TrailerUrl,
		PostersPath: s.PostersPath,
		Resume:      s.Resume,
		Rating:      s.Rating,
		ReleaseDate: s.ReleaseDate.AsTime(),
		WrittenBy:   s.WrittenBy.ToModel(),
		ProducedBy:  s.ProducedBy.ToModel(),
		DirectedBy:  s.DirectedBy.ToModel(),
		Episodes:    s.Episodes.ToModel().(models.ShortEpisodes),
	}
}

func (s *ShortEpisodeList) ToModel() interface{} {
	shortEpisodesModel := models.ShortEpisodes{}
	for _, shortEpisode := range s.ShortEpisodes {
		shortEpisodesModel = append(shortEpisodesModel, shortEpisode.ToModel().(*models.ShortEpisode))
	}
	return shortEpisodesModel
}

func (s *ShortEpisode) ToModel() interface{} {
	return &models.ShortEpisode{
		ID:          s.Id,
		Title:       s.Title,
		PostersPath: s.PostersPath,
		Rating:      s.Rating,
		Resume:      s.Resume,
	}
}

func (s *ListSeasonResponse) ToModel() models.Seasons {
	seasons := models.Seasons{}
	for _, season := range s.Seasons {
		seasons = append(seasons, season.ToModel().(*models.Season))
	}
	return seasons
}

func (g *GenreListResponse) ToModel() models.Genres {
	genres := models.Genres{}
	for _, g := range g.Genres {
		genres = append(genres, g.ToModel().(*models.Genre))
	}
	return genres
}

func (g *Genre) ToModel() interface{} {
	return &models.Genre{
		ID:          g.Id,
		Name:        g.Name,
		Description: g.Description,
	}
}
func (a *Article) ToModel() interface{} {
	return &models.Article{
		ID:          a.Id,
		Title:       a.Title,
		ReleaseDate: a.ReleaseDate.AsTime(),
		PostersPath: a.PostersPath,
		Description: a.Description,
		Journalist:  models.ShortJournalist{ID: a.Journalist.Id},
	}
}

func (a *ArticleListResponse) ToModel() models.Articles {
	articles := models.Articles{}
	for _, article := range a.Articles {
		articles = append(articles, article.ToModel().(*models.Article))
	}
	return articles
}

func (s *ShowListResponse) ToModel() models.Shows {
	shows := models.Shows{}
	for _, show := range s.Shows {
		shows = append(shows, show.ToModel().(*models.Show))
	}
	return shows
}

func (s *Show) ToModel() interface{} {
	return &models.Show{
		ID:          s.Id,
		Title:       s.Title,
		Type:        s.Type,
		PostersPath: s.PostersPath,
		ReleaseDate: s.ReleaseDate.AsTime(),
		EndDate:     s.EndDate.AsTime(),
		Rating:      s.Rating,
		Length:      *s.Length.ToModel(),
		TrailerURL:  s.TrailerUrl,
		Genres:      s.Genres.ToModel(),
		WrittenBy:   s.WrittenBy.ToModel(),
		ProducedBy:  s.ProducedBy.ToModel(),
		DirectedBy:  s.DirectedBy.ToModel(),
		Starring:    s.Starring.ToModel(),
		Description: s.Description,
		Seasons:     s.Seasons.ToModel(),
	}
}

func (s *ShortSeasons) ToModel() models.ShortSeasons {
	seasons := models.ShortSeasons{}
	for _, season := range s.Seasons {
		seasons = append(seasons, season.ToModel())
	}
	return seasons
}

func (s *ShortSeason) ToModel() *models.ShortSeason {
	return &models.ShortSeason{
		ID:          s.Id,
		Title:       s.Title,
		PostersPath: s.PostersPath,
		Rating:      s.Rating,
	}
}

func (g *ShortGenres) ToModel() models.ShortGenres {
	genres := models.ShortGenres{}
	for _, genre := range g.Genres {
		genres = append(genres, genre.ToModel())
	}
	return genres
}

func (s *ShortGenre) ToModel() *models.ShortGenre {
	return &models.ShortGenre{
		ID:   s.Id,
		Name: s.Name,
	}
}

func (e *ListEpisodeResponse) ToModel() models.Episodes {
	episodes := models.Episodes{}
	for _, episode := range e.Episodes {
		episodes = append(episodes, episode.ToModel().(*models.Episode))
	}
	return episodes
}

func (e *Episode) ToModel() interface{} {
	return &models.Episode{
		ID:          e.Id,
		SeasonID:    e.SeasonId,
		Title:       e.Title,
		PostersPath: e.PostersPath,
		TrailerURL:  e.TrailerUrl,
		Length:      *e.ShowLength.ToModel(),
		Rating:      e.Rating,
		Resume:      e.Resume,
		WrittenBy:   e.WrittenBy.ToModel(),
		ProducedBy:  e.ProducedBy.ToModel(),
		DirectedBy:  e.DirectedBy.ToModel(),
		Starring:    e.Starring.ToModel(),
	}
}

func (s *ShowLength) ToModel() *models.ShowLength {
	return &models.ShowLength{
		Hours:   int(s.Hours),
		Minutes: int(s.Minutes),
	}
}

func (s *ShortCelebrities) ToModel() models.ShortCelebrities {
	celebs := models.ShortCelebrities{}
	for _, celeb := range s.ShortCelebs {
		celebs = append(celebs, celeb.ToModel())
	}
	return celebs
}

func (s *ShortCelebrity) ToModel() *models.ShortCelebrity {
	return &models.ShortCelebrity{
		ID:          s.Id,
		Name:        s.Name,
		RoleName:    s.RoleName,
		PostersPath: s.PostersPath,
	}
}

func (f *FilmCrew) ToModel() models.FilmCrews {
	crews := models.FilmCrews{}
	for _, crew := range f.FilmCrew {
		crews = append(crews, crew.ToModel())
	}
	return crews
}

func (f *FilmStaff) ToModel() *models.FilmCrew {
	return &models.FilmCrew{
		ID:          f.Id,
		Name:        f.Name,
		PostersPath: f.PostersPath,
	}
}
