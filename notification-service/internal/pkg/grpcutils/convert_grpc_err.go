package grpcutils

import (
	"encoding/json"

	pkgerrors "github.com/pkg/errors"
)

type GRPCError struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

func ConvertGRPCError(strErr string) (GRPCError, error) {
	grpcErr := GRPCError{}
	if err := json.Unmarshal([]byte(strErr), &grpcErr); err != nil {
		return GRPCError{}, pkgerrors.WithStack(err)
	}

	return grpcErr, nil
}
