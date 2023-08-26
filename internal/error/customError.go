package myError

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var LoggedOutErr = errors.New("user logged out")
var UserNotExist = status.Error(codes.InvalidArgument, "user does not exist")
