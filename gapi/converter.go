package gapi

import (
	db "github.com/dokilife/im/db/sqlc"
	"github.com/dokilife/im/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertMessage(message db.Message) *pb.Message {
	return &pb.Message{
		From:      message.From,
		Message:   message.Message,
		CreatedAt: timestamppb.New(message.CreatedAt),
	}
}
