package rpc

import (
	"context"
	"errors"

	con "github.com/abuabdillatief/s16-research-ventures/context"
	"github.com/abuabdillatief/s16-research-ventures/generated"
	"github.com/abuabdillatief/s16-research-ventures/services"
	"google.golang.org/grpc"
)

type OmdbServer struct {
	generated.UnimplementedOMDBServiceServer
}

func NewOMDBService(srv *grpc.Server) {
	generated.RegisterOMDBServiceServer(srv, &OmdbServer{})
}

func (o *OmdbServer) GetMovieByID(ctx context.Context, req *generated.GetMovieByIDRequest) (*generated.GetMovieByIDResponse, error) {
	if !con.IsAuthenticated(ctx) {
		return nil, errors.New("request is not authenticated")
	}
	return services.GetMovieByID(ctx, req)
}

func (o *OmdbServer) SearchMovies(ctx context.Context, req *generated.SearchMoviesRequest) (*generated.SearchMoviesResponse, error) {
	if !con.IsAuthenticated(ctx) {
		return nil, errors.New("request is not authenticated")
	}
	return services.SearchMovies(ctx, req)
}
