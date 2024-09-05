package grpc

import (
	"context"
	"strings"

	"github.com/kytruong0712/goffee-shop/menu-service/internal/controller/category"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/handler/grpc/protobuf"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/model"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/pkg/grpcerrorutils"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateCategory is gRPC function to support create new category
func (i impl) CreateCategory(ctx context.Context, req *protobuf.CreateCategoryRequest) (*protobuf.Category, error) {
	inp, err := validateAndMapToCreateCategoryInput(req)
	if err != nil {
		return nil, err
	}

	rs, err := i.categoryCtrl.Create(ctx, inp)
	if err != nil {
		switch err {
		case category.ErrCategoryAlreadyExists:
			return nil, grpcerrorutils.ErrDetails(codes.NotFound, category.ErrCategoryAlreadyExists.Error(), category.ErrCategoryAlreadyExists.Error(), "name")
		default:
			return nil, grpcerrorutils.ErrDetails(codes.Internal, err.Error(), err.Error(), "")
		}
	}

	return toCategory(rs), nil
}

func validateAndMapToCreateCategoryInput(req *protobuf.CreateCategoryRequest) (category.CreateInput, error) {
	if strings.TrimSpace(req.Name) == "" {
		return category.CreateInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, ErrCategoryNameIsRequired.Error(), ErrCategoryNameIsRequired.Error(), "name")
	}

	if strings.TrimSpace(req.Description) == "" {
		return category.CreateInput{}, grpcerrorutils.ErrDetails(codes.InvalidArgument, ErrCategoryDescriptionIsRequired.Error(), ErrCategoryDescriptionIsRequired.Error(), "description")
	}

	return category.CreateInput{
		Name:        req.Name,
		Description: req.Description,
	}, nil
}

func toCategory(rs model.Category) *protobuf.Category {
	var stt protobuf.CategoryStatus
	if rs.Status == model.CategoryStatusActive {
		stt = protobuf.CategoryStatus_ACTIVE
	} else if rs.Status == model.CategoryStatusInactive {
		stt = protobuf.CategoryStatus_INACTIVE
	}

	return &protobuf.Category{
		Id:          rs.ID,
		Name:        rs.Name,
		Description: rs.Description,
		Status:      stt,
		CreatedAt:   timestamppb.New(rs.CreatedAt),
		UpdatedAt:   timestamppb.New(rs.UpdatedAt),
	}
}
