package grpc

import (
	"context"
	"errors"
	"github.com/thegodeveloper/movieapp/metadata/internal/controller/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/thegodeveloper/movieapp/gen"
	"github.com/thegodeveloper/movieapp/metadata/internal/controller"
	"github.com/thegodeveloper/movieapp/metadata/internal/repository"
	"github.com/thegodeveloper/movieapp/metadata/pkg/model"
)

// Handler defines a movie metadata gRPC handler.
type Handler struct {
	gen.UnimplementedMetadataServiceServer
	ctrl *controller.MetadataService
}

// New creates a new movie metadata gRPC handler.
func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// GetMetadataByID returns movie metadata by id.
func (h *Handler) GetMetadata(ctx context.Context, req *gen.GetMetadataRequest) (*gen.GetMetadataResponse, error) {
	if req == nil || req.MovieId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.ctrl.Get(ctx, req.MovieId)
	if err != nil && errors.Is(err, controller.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetMetadataResponse{
		Metadata: model.MetadataToProto(m)
	}, nil
}