package models

import "int-gateway/response"

func (u *UploadResponse) ToResponse() *response.UploadImages {
	images := response.UploadImages{}
	images.PostersPath = append(images.PostersPath, u.PostersPath...)
	return &images
}

func (g *Gender) ToResponse() response.Gender {
	return response.Gender(*g)
}

func (c *Celebrities) ToResponse() response.Celebrities {
	celebs := response.Celebrities{}
	for _, celeb := range *c {
		celebs = append(celebs, celeb.ToResponse())
	}
	return celebs
}

func (c *Celebrity) ToResponse() *response.Celebrity {
	return &response.Celebrity{
		ID:           c.ID,
		Name:         c.Name,
		Occupation:   c.Occupation,
		PostersPath:  c.PostersPath,
		DateOfBirth:  c.DateOfBirth,
		DateOfDeath:  c.DateOfDeath,
		PlaceOfBirth: c.PlaceOfBirth,
		Gender:       c.Gender.ToResponse(),
		Bio:          c.Bio,
	}
}

func (s *Season) ToResponse() *response.Season {
	return &response.Season{
		ID:          s.ID,
		ShowID:      s.ShowID,
		Title:       s.Title,
		TrailerURL:  s.TrailerURL,
		PostersPath: s.PostersPath,
		Resume:      s.Resume,
		Rating:      s.Rating,
		ReleaseDate: s.ReleaseDate,
		WrittenBy:   s.WrittenBy.ToResponse(),
		ProducedBy:  s.ProducedBy.ToResponse(),
		DirectedBy:  s.DirectedBy.ToResponse(),
		Episodes:    s.Episodes.ToResponse(),
	}
}

func (ss *Seasons) ToResponse() response.Seasons {
	seasons := response.Seasons{}
	for _, season := range *ss {
		seasons = append(seasons, season.ToResponse())
	}
	return seasons
}

func (f *FilmCrews) ToResponse() response.FilmCrews {
	crews := response.FilmCrews{}
	for _, crew := range *f {
		crews = append(crews, crew.ToResponse())
	}
	return crews
}

func (f *FilmCrew) ToResponse() *response.FilmCrew {
	return &response.FilmCrew{
		ID:          f.ID,
		Name:        f.Name,
		PostersPath: f.PostersPath,
	}
}

func (e *ShortEpisode) ToResponse() *response.ShortEpisode {
	return &response.ShortEpisode{
		ID:          e.ID,
		Title:       e.Title,
		PostersPath: e.PostersPath,
		Rating:      e.Rating,
		Resume:      e.Resume,
	}
}

func (e *ShortEpisodes) ToResponse() response.ShortEpisodes {
	shortEps := response.ShortEpisodes{}
	for _, shortEp := range *e {
		shortEps = append(shortEps, shortEp.ToResponse())
	}
	return shortEps
}

func (g *Genres) ToResponse() response.Genres {
	genres := response.Genres{}
	for _, g := range *g {
		genres = append(genres, g.ToResponse())
	}
	return genres
}

func (g *Genre) ToResponse() *response.Genre {
	return &response.Genre{
		ID:          g.ID,
		Name:        g.Name,
		Description: g.Description,
	}
}

func (a *Article) ToResponse() *response.Article {
	return &response.Article{
		ID:          a.ID,
		Title:       a.Title,
		ReleaseDate: a.ReleaseDate,
		PostersPath: a.PostersPath,
		Description: a.Description,
		Journalist:  response.ShortJournalist(a.Journalist),
	}
}

func (a *Articles) ToResponse() response.Articles {
	articles := response.Articles{}
	for _, article := range *a {
		articles = append(articles, article.ToResponse())
	}
	return articles
}

func (j *Journalists) ToResponse() response.Journalists {
	var resp []*response.Journalist
	for _, jst := range *j {
		resp = append(resp, jst.ToResponse())
	}
	return resp
}

func (j *Journalist) ToResponse() *response.Journalist {
	return &response.Journalist{
		ID:   j.ID,
		Name: j.Name,
	}
}

func (s *Shows) ToResponse() response.Shows {
	shows := response.Shows{}
	for _, show := range *s {
		shows = append(shows, show.ToResponse())
	}
	return shows
}

func (s *Show) ToResponse() *response.Show {
	return &response.Show{
		ID:          s.ID,
		Title:       s.Title,
		TrailerURL:  s.TrailerURL,
		Type:        s.Type,
		PostersPath: s.PostersPath,
		ReleaseDate: s.ReleaseDate,
		EndDate:     s.EndDate,
		Rating:      s.Rating,
		Length:      *s.Length.ToResponse(),
		Genres:      s.Genres.ToResponse(),
		DirectedBy:  s.DirectedBy.ToResponse(),
		ProducedBy:  s.ProducedBy.ToResponse(),
		WrittenBy:   s.WrittenBy.ToResponse(),
		Starring:    s.Starring.ToResponse(),
		Description: s.Description,
		Seasons:     s.Seasons.ToResponse(),
	}
}

func (s *ShortSeasons) ToResponse() response.ShortSeasons {
	seasons := response.ShortSeasons{}
	for _, season := range *s {
		seasons = append(seasons, season.ToResponse())
	}
	return seasons
}

func (s *ShortSeason) ToResponse() *response.ShortSeason {
	return &response.ShortSeason{
		ID:          s.ID,
		Title:       s.Title,
		PostersPath: s.PostersPath,
		Rating:      s.Rating,
	}
}

func (g *ShortGenres) ToResponse() response.ShortGenres {
	genres := response.ShortGenres{}
	for _, genre := range *g {
		genres = append(genres, genre.ToResponse())
	}
	return genres
}

func (g *ShortGenre) ToResponse() *response.ShortGenre {
	return &response.ShortGenre{
		ID:   g.ID,
		Name: g.Name,
	}
}

func (e *Episodes) ToResponse() response.Episodes {
	episodes := response.Episodes{}
	for _, episode := range *e {
		episodes = append(episodes, episode.ToResponse())
	}
	return episodes
}

func (e *Episode) ToResponse() *response.Episode {
	return &response.Episode{
		ID:          e.ID,
		SeasonID:    e.SeasonID,
		Title:       e.Title,
		PostersPath: e.PostersPath,
		TrailerURL:  e.TrailerURL,
		Length:      *e.Length.ToResponse(),
		Rating:      e.Rating,
		Resume:      e.Resume,
		WrittenBy:   e.WrittenBy.ToResponse(),
		ProducedBy:  e.ProducedBy.ToResponse(),
		DirectedBy:  e.DirectedBy.ToResponse(),
		Starring:    e.Starring.ToResponse(),
	}
}

func (s *ShowLength) ToResponse() *response.ShowLength {
	return &response.ShowLength{
		Hours:   s.Hours,
		Minutes: s.Minutes,
	}
}

func (s *ShortCelebrities) ToResponse() response.ShortCelebrities {
	celebs := response.ShortCelebrities{}
	for _, celeb := range *s {
		celebs = append(celebs, celeb.ToResponse())
	}
	return celebs
}

func (s *ShortCelebrity) ToResponse() *response.ShortCelebrity {
	return &response.ShortCelebrity{
		ID:          s.ID,
		Name:        s.Name,
		RoleName:    s.RoleName,
		PostersPath: s.PostersPath,
	}
}
