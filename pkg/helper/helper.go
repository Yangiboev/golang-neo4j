package helper

import (
	"github.com/Yangiboev/golang-neo4j/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleError(log logger.Logger, err error, message string, req interface{}) error {

	if err != nil {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return status.Error(codes.Internal, message)
	}
	return nil
}
