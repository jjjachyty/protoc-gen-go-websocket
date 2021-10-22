// Code generated by protoc-gen-go-websocket. DO NOT EDIT.
// versions:
// protoc-gen-go-websocket 1.0

package pb

import (
	context "context"
	json "encoding/json"
	errors "errors"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	melody_v1 "gopkg.in/olahol/melody.v1"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ json.Delim
var _ http.ConnState

const _ = http1.SupportPackageIsVersion1

type UserServiceWebSocketRequest struct {
	Action string
	Input  string
}

type UserServiceWebSocketResponse struct {
	Action  string
	Code    int
	Message string
	Data    interface{}
}

func (ws *UserServiceWebSocketResponse) Error(ctx context.Context) (rs []byte) {
	rs, _ = json.Marshal(ws)
	return
}

func RegisterUserServiceWebSocketServer(s *http1.Server, srv UserServiceServer, sessionHandler func(ctx context.Context, sid string) (context.Context, error)) {
	var m = melody_v1.New()

	s.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		m.HandleRequest(rw, r)
	})

	m.HandleMessage(func(s *melody_v1.Session, msg []byte) {
		var rs []byte

		wsmsg := &UserServiceWebSocketRequest{}
		err := json.Unmarshal(msg, wsmsg)
		if err != nil {
			rs, _ = json.Marshal(UserServiceWebSocketResponse{Action: wsmsg.Action, Code: 500, Message: err.Error()})
			s.Write(rs)
			return
		}

		ctx, err := sessionHandler(s.Request.Context(), s.Request.Header.Get("sid"))
		if err != nil {
			rs, _ = json.Marshal(UserServiceWebSocketResponse{Action: wsmsg.Action, Code: 500, Message: err.Error()})
			s.Write(rs)
			return
		}

		var out interface{}
		if out, err = wsmsg.HandleMessage(ctx, srv); err != nil {
			rs, _ = json.Marshal(UserServiceWebSocketResponse{Action: wsmsg.Action, Code: 500, Message: err.Error()})
			s.Write(rs)
			return
		}

		rs, _ = json.Marshal(UserServiceWebSocketResponse{Action: wsmsg.Action, Code: 200, Data: out})
		s.Write(rs)
	})
}

func (ws *UserServiceWebSocketRequest) HandleMessage(ctx context.Context, srv UserServiceServer) (interface{}, error) {
	var err error
	switch ws.Action {
	case "/user.UserService/UserRegister":
		var in UserRegisterReq
		if err = json.Unmarshal([]byte(ws.Input), &in); err != nil {
			return nil, err
		}
		return srv.UserRegister(ctx, &in)
	}
	return nil, errors.New("action not found")
}