package errcode

import (
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	pb "grpc-example/proto"
	"log"
)

var errCodes = map[uint32]*pb.Response{}

type Status struct {
	status *status.Status
}

func New(code codes.Code, msg string) *Status {
	return NewWithData(code, msg, nil)
}

func NewWithData(code codes.Code, msg string, data *any.Any) *Status {
	if _, ok := errCodes[uint32(code)]; ok {
		log.Panicf("Code[%d] is exist ", code)
	}

	errCode := &pb.Response{
		Code: uint32(code),
		Msg:  msg,
		Data: data,
	}

	s, _ := status.New(code, msg).WithDetails(errCode)

	errCodes[uint32(code)] = errCode

	return &Status{s}
}

func (s *Status) Code() uint32 {
	return uint32(s.status.Code())
}

func (s *Status) Response(msgs ...proto.Message) (*pb.Response, error) {
	resp := &pb.Response{
		Code: Success.Code(),
		Msg:  Success.status.Message(),
		Data: nil,
	}

	var _any *any.Any
	var err error
	for _, msg := range msgs {
		if msg != proto.Message(nil) {
			_any, err = anypb.New(msg)
			log.Println("s.Code()", s.Code())
			resp.Code = s.Code()
			log.Println("s.status.Message()", s.status.Message())
			resp.Msg = s.status.Message()
			resp.Data = _any
		}
		break
	}

	return resp, err
}

//
//type InternalError struct {
//	Code uint32       `protobuf:"code" json:"code" example:"200" example:"400" example:"500" example:"502"`
//	Msg  string      `protobuf:"msg" json:"msg" example:"ok"`
//	Data interface{} `protobuf:"data" json:"data,omitempty"`
//}
//
//func (e *InternalError) Error() string {
//	return fmt.Sprintf("%d: %status", e.Code, e.Msg)
//}
//
//func (e *InternalError) StatusCode() int {
//	switch codes.Code(e.Code) {
//	case Success.Code():
//		return http.StatusOK
//	case BadRequest.Code():
//		return http.StatusBadRequest
//	case ServerError.Code():
//		return http.StatusInternalServerError
//	case ErrorAuthCheckTokenFail.Code(), ErrorAuthCheckTokenTimeout.Code(), ErrorAuthToken.Code(), ErrorAuth.Code():
//		return http.StatusUnauthorized
//	default:
//		return http.StatusInternalServerError
//	}
//}
//
//func FromError(err error) *status.Status {
//	status, _ := status.FromError(err)
//	return status
//}
//
//func GetMsg(code uint32) string {
//	e, ok := errCodes[code]
//	if ok {
//		return e.Msg
//	}
//	return ""
//}
//
//type ValiadError struct {
//	Key string
//	Msg string
//}
//
//type ValiadErrors []*ValiadError
//
//func (v *ValiadError) Error() string {
//	return v.Msg
//}
//
//func (v ValiadErrors) Error() string {
//	return strings.Join(v.Errors(), ",")
//}
//
//func (v ValiadErrors) Errors() []string {
//	var errs []string
//	for _, err := range v {
//		errs = append(errs, err.Error())
//	}
//	return errs
//}
