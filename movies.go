package trakt

import "fmt"

var (
	MovieURL         = Hyperlink("movies/{traktID}")
	MoviesPopularURL = Hyperlink("movies/popular")
	MoviesSearchURL  = Hyperlink("search/movie?query={query}")
	MovieByIDURL     = Hyperlink("search/{id_type}/{id}?type=movie")
)

// Create a MoviesService with the base url.URL
func (c *Client) Movies() (movies *MoviesService) {
	movies = &MoviesService{client: c}
	return
}

type MoviesService struct {
	client *Client
}

// One returns a single movie identified by a Trakt ID. It also returns a Result
// object to inspect the returned response of the server.
func (r *MoviesService) One(traktID int) (movie *Movie, result *Result) {
	url, _ := MovieURL.Expand(M{"traktID": fmt.Sprintf("%d", traktID)})
	result = r.client.get(url, &movie)
	return
}

func (r *MoviesService) OneOfType(id string, idType string) (movie *Movie, result *Result) {
	movies := []MovieResult{}
	url, _ := MovieByIDURL.Expand(M{"id_type": idType, "id": id})
	result = r.client.get(url, &movies)
	if len(movies) > 0 {
		return movies[0].Movie, result
	}
	return nil, result
}

func (r *MoviesService) AllPopular() (movies []Movie, result *Result) {
	url, _ := MoviesPopularURL.Expand(M{})
	result = r.client.get(url, &movies)
	return
}

func (r *MoviesService) Search(query string) (movies []MovieResult, result *Result) {
	url, _ := MoviesSearchURL.Expand(M{"query": query})
	result = r.client.get(url, &movies)
	return
}

type MovieResult struct {
	Score float64 `json:"score"`
	Movie *Movie  `json:"movie"`
	Type  string  `json:"type"`
}
