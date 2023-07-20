package services

import (
	"errors"
	"testing"

	"github.com/abuabdillatief/s16-research-ventures/generated"
	"github.com/stretchr/testify/assert"
)

func TestValidateGetMovieByIDRequest(t *testing.T) {
	type args struct {
		req *generated.GetMovieByIDRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err error
	}{
		{
			name: "Success case",
			args: args{
				req: &generated.GetMovieByIDRequest{
					Id: "tt0294568",
				},
			},
			wantErr: false,
		},
		{
			name: "Fail case",
			args: args{
				req: &generated.GetMovieByIDRequest{
					Id: "",
				},
			},
			wantErr: true,
			err: errors.New("id is invalid"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateGetMovieByIDRequest(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("ValidateGetMovieByIDRequest() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

func TestValidateSearchMovieRequest(t *testing.T) {
	type args struct {
		req *generated.SearchMoviesRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Success case",
			args: args{
				req: &generated.SearchMoviesRequest{
					Query: "spiderman",
					Type:  "movie",
					Page:  1,
				},
			},
			wantErr: false,
		},
		{
			name: "Fail case when query is invalid",
			args: args{
				req: &generated.SearchMoviesRequest{
					Query: "",
					Type:  "movie",
					Page:  1,
				},
			},
			err:     errors.New("query cannot be empty"),
			wantErr: true,
		},
		{
			name: "Fail case when type is invalid",
			args: args{
				req: &generated.SearchMoviesRequest{
					Query: "spiderman",
					Type:  "drama",
					Page:  1,
				},
			},
			wantErr: true,
			err:     errors.New("type is not valid. valid types: movie, series, episode"),
		},
		{
			name: "Fail case when page is invalid",
			args: args{
				req: &generated.SearchMoviesRequest{
					Query: "spiderman",
					Type:  "movie",
					Page:  144,
				},
			},
			wantErr: true,
			err:     errors.New("page number is invalid. valid page: 1 - 100"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSearchMovieRequest(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSearchMovieRequest() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.err, err)
			}
		})
	}
}
