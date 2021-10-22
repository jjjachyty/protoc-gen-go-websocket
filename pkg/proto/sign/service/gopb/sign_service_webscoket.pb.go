// Code generated by protoc-gen-go-websocket. DO NOT EDIT.
// versions:
// protoc-gen-go-websocket 1.0

package gopb

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

type SignServiceWebSocketRequest struct {
	Action string
	Input  string
}

type SignServiceWebSocketResponse struct {
	Action  string
	Code    int
	Message string
	Data    interface{}
}

func (ws *SignServiceWebSocketResponse) Error(ctx context.Context) (rs []byte) {
	rs, _ = json.Marshal(ws)
	return
}

func RegisterSignServiceWebSocketServer(s *http1.Server, srv SignServiceServer, sessionHandler func(ctx context.Context, sid string) (context.Context, error)) {
	var m = melody_v1.New()

	s.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		m.HandleRequest(rw, r)
	})

	m.HandleMessage(func(s *melody_v1.Session, msg []byte) {
		var rs []byte

		wsmsg := &SignServiceWebSocketRequest{}
		err := json.Unmarshal(msg, wsmsg)
		if err != nil {
			rs, _ = json.Marshal(SignServiceWebSocketResponse{Action: wsmsg.Action, Code: 500, Message: err.Error()})
			s.Write(rs)
			return
		}

		ctx, err := sessionHandler(s.Request.Context(), s.Request.Header.Get("sid"))
		if err != nil {
			rs, _ = json.Marshal(SignServiceWebSocketResponse{Action: wsmsg.Action, Code: 500, Message: err.Error()})
			s.Write(rs)
			return
		}

		var out interface{}
		if out, err = wsmsg.HandleMessage(ctx, srv); err != nil {
			rs, _ = json.Marshal(SignServiceWebSocketResponse{Action: wsmsg.Action, Code: 500, Message: err.Error()})
			s.Write(rs)
			return
		}

		rs, _ = json.Marshal(SignServiceWebSocketResponse{Action: wsmsg.Action, Code: 200, Data: out})
		s.Write(rs)
	})
}

func (ws *SignServiceWebSocketRequest) HandleMessage(ctx context.Context, srv SignServiceServer) (interface{}, error) {
	var err error
	switch ws.Action {
	case "/pb.SignService/OrgAccountAdd":
		var in OrgAccountAddReq
		if err = json.Unmarshal([]byte(ws.Input), &in); err != nil {
			return nil, err
		}
		return srv.OrgAccountAdd(ctx, &in)
	case "/pb.SignService/UserAccountAdd":
		var in UserAccountAddReq
		if err = json.Unmarshal([]byte(ws.Input), &in); err != nil {
			return nil, err
		}
		return srv.UserAccountAdd(ctx, &in)
	case "/pb.SignService/TemplateCreate":
		var in CreateByTemplateReq
		if err = json.Unmarshal([]byte(ws.Input), &in); err != nil {
			return nil, err
		}
		return srv.TemplateCreate(ctx, &in)
	case "/pb.SignService/TemplateGet":
		var in TemplateGetReq
		if err = json.Unmarshal([]byte(ws.Input), &in); err != nil {
			return nil, err
		}
		return srv.TemplateGet(ctx, &in)
	case "/pb.SignService/FlowOneStepCreate":
		var in FlowOneStepCreateReq
		if err = json.Unmarshal([]byte(ws.Input), &in); err != nil {
			return nil, err
		}
		return srv.FlowOneStepCreate(ctx, &in)
	case "/pb.SignService/VoucherIssue":
		var in VoucherIssueReq
		if err = json.Unmarshal([]byte(ws.Input), &in); err != nil {
			return nil, err
		}
		return srv.VoucherIssue(ctx, &in)
	case "/pb.SignService/Notice":
		var in NoticeReq
		if err = json.Unmarshal([]byte(ws.Input), &in); err != nil {
			return nil, err
		}
		return srv.Notice(ctx, &in)
	}
	return nil, errors.New("action not found")
}
