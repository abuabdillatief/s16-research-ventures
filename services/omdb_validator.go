package services

import (
	"errors"
	"strings"

	"github.com/abuabdillatief/s16-research-ventures/generated"
)

func ValidateGetMovieByIDRequest(req *generated.GetMovieByIDRequest) error {
	if req.Id == "" {
		return errors.New("id is invalid")
	}

	return nil
}

func ValidateSearchMovieRequest(req *generated.SearchMoviesRequest) error {
	// shows in the omdb documentation that empty string
	// is valid request. adding this as a requirement
	if req.Query == "" {
		return errors.New("query cannot be empty")
	}

	if req.Page < 1 || req.Page > 100 {
		return errors.New("page number is invalid. valid page: 1 - 100")
	}

	switch strings.ToLower(req.Type) {
	case "movie", "series", "episode":
	default:
		return errors.New("type is not valid. valid types: movie, series, episode")
	}

	return nil
}
