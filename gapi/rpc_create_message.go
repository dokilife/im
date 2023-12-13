package gapi

import (
	"context"

	db "github.com/dokilife/im/db/sqlc"
	"github.com/dokilife/im/pb"
	"github.com/dokilife/im/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateMessage(ctx context.Context, req *pb.CreateMessageRequest) (*pb.CreateMessageResponse, error) {
	violations := validateCreateMessageRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateMessageTxParams{
		CreateMessageParams: db.CreateMessageParams{
			From:    req.GetFrom(),
			Message: req.GetMessage(),
		},
	}

	txResult, err := server.store.CreateMessageTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create message: %s", err)
	}

	rsp := &pb.CreateMessageResponse{
		Message: convertMessage(txResult.Message),
	}

	return rsp, nil
}

func validateCreateMessageRequest(req *pb.CreateMessageRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateMessage(req.GetMessage()); err != nil {
		violations = append(violations, fieldViolation("message", err))
	}

	return violations
}
