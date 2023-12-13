package gapi

import (
	"github.com/dokilife/im/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetMessageStream implements the GetMessageStream RPC method
func (server *Server) GetMessageStream(req *emptypb.Empty, stream pb.IM_GetMessageStreamServer) error {
	// You can add any additional logic before entering the loop, if necessary

	// Create a new connection to listen for NOTIFY events
	conn, err := server.store.AcquireConn(stream.Context())
	if err != nil {
		return status.Errorf(codes.Internal, "failed to acquire connection: %s", err)
	}
	defer conn.Release()

	// Listen for NOTIFY events on the 'message_insert' channel
	_, err = conn.Exec(stream.Context(), "LISTEN message_insert")
	if err != nil {
		return status.Errorf(codes.Internal, "failed to execute LISTEN command: %s", err)
	}

	for {
		// Check if the client has requested to stop streaming
		if stream.Context().Err() != nil {
			return nil
		}

		// Wait for a NOTIFY event
		_, err := conn.Conn().WaitForNotification(stream.Context())
		if err != nil {
			return status.Errorf(codes.Internal, "failed to wait for notification: %s", err)
		}

		// Fetch the new message from the database
		message, err := server.store.GetMessage(stream.Context())
		if err != nil {
			return status.Errorf(codes.Internal, "failed to fetch message: %s", err)
		}

		// Send the new message to the client
		if err := stream.Send(&pb.GetMessageStreamResponse{
			Message: convertMessage(message),
		}); err != nil {
			// Check if the client has closed the connection
			if stream.Context().Err() != nil {
				return nil
			}

			return status.Errorf(codes.Internal, "failed to send message: %s", err)
		}
	}
}
