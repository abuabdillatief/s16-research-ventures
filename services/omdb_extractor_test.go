package services

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/abuabdillatief/s16-research-ventures/generated"
)

func TestExtractGetMovieByIDResponse(t *testing.T) {
	successBody := `
	{"Title":"Batman: The Killing Joke","Year":"2016","Rated":"R","Released":"25 Jul 2016","Runtime":"76 min","Genre":"Animation, Action, Crime","Director":"Sam Liu","Writer":"Brian Azzarello, Brian Bolland, Bob Kane","Actors":"Kevin Conroy, Mark Hamill, Tara Strong","Plot":"As Batman hunts for the escaped Joker, the Clown Prince of Crime attacks the Gordon family to prove a diabolical point mirroring his own fall into madness.","Language":"English","Country":"United States","Awards":"1 win & 2 nominations","Poster":"https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"6.4/10"},{"Source":"Rotten Tomatoes","Value":"36%"}],"Metascore":"N/A","imdbRating":"6.4","imdbVotes":"59,539","imdbID":"tt4853102","Type":"movie","DVD":"02 Aug 2016","BoxOffice":"$3,775,000","Production":"N/A","Website":"N/A","Response":"True"}
	`
	type args struct {
		resp *http.Response
		req  *generated.GetMovieByIDRequest
	}

	tests := []struct {
		name    string
		args    args
		want    *generated.GetMovieByIDResponse
		wantErr bool
	}{
		{
			name: "Success Case",
			args: args{
				resp: &http.Response{
					Body: io.NopCloser(strings.NewReader(successBody)),
				},
				req: &generated.GetMovieByIDRequest{
					Id: "tt4853102",
				},
			},
			want: &generated.GetMovieByIDResponse{
				Id:       "tt4853102",
				Title:    "Batman: The Killing Joke",
				Year:     "2016",
				Rated:    "R",
				Genre:    "Animation, Action, Crime",
				Plot:     "As Batman hunts for the escaped Joker, the Clown Prince of Crime attacks the Gordon family to prove a diabolical point mirroring his own fall into madness.",
				Director: "Sam Liu",
				Actors: []string{
					"Kevin Conroy", "Mark Hamill", "Tara Strong",
				},
				Language:  "English",
				Country:   "United States",
				Type:      "movie",
				PosterUrl: "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractGetMovieByIDResponse(tt.args.resp, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractGetMovieByIDResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractGetMovieByIDResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractSearchMovieResponse(t *testing.T) {
	// correct json response from omdb
	successBody := `
	{"Search":[{"Title":"Spiderman and Grandma","Year":"2009","imdbID":"tt1433184","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMjE3Mzg0MjAxMl5BMl5BanBnXkFtZTcwNjIyODg5Mg@@._V1_SX300.jpg"}],"totalResults":"1","Response":"True"}
	`
	// wrong json response from omdb
	failBody := `
	{"result":[{"Title":"Spiderman and Grandma","Year":"2009","imdbID":"tt1433184","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMjE3Mzg0MjAxMl5BMl5BanBnXkFtZTcwNjIyODg5Mg@@._V1_SX300.jpg"}],"totalResults":"1","Response":"True"}
	`

	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *generated.SearchMoviesResponse
		wantErr bool
	}{
		{
			name: "Success case",
			args: args{
				resp: &http.Response{
					Body: io.NopCloser(strings.NewReader(successBody)),
				},
			},
			want: &generated.SearchMoviesResponse{
				Movies: []*generated.MovieResult{{
					Id:        "tt1433184",
					Title:     "Spiderman and Grandma",
					Year:      "2009",
					Type:      "movie",
					PosterUrl: "https://m.media-amazon.com/images/M/MV5BMjE3Mzg0MjAxMl5BMl5BanBnXkFtZTcwNjIyODg5Mg@@._V1_SX300.jpg",
				}},
				TotalResults: 1,
			},
			wantErr: false,
		},
		{
			name: "Fail case when json response from omdb is invalid, causing empty array in response",
			args: args{
				resp: &http.Response{
					Body: io.NopCloser(strings.NewReader(failBody)),
				},
			},
			want: &generated.SearchMoviesResponse{
				TotalResults: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractSearchMovieResponse(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractSearchMovieResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractSearchMovieResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
