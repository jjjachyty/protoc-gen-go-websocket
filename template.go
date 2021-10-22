package main

import (
	"bytes"
	"strings"
	"text/template"
)

var WebSocketTemplate = `
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}


type {{.ServiceType}}WebSocketRequest struct {
	Action  string
	Input  string
}

type {{.ServiceType}}WebSocketResponse struct {
	Action string
	Code  int
	Message string
	Data    interface{}
}




func (ws *{{.ServiceType}}WebSocketResponse)Error(ctx context.Context) (rs []byte){
	rs,_ = json.Marshal(ws)
	return 
}



func Register{{.ServiceType}}WebSocketServer(s *http1.Server, srv {{.ServiceType}}Server,sessionHandler func(ctx context.Context, sid string) (context.Context, error)) {
	var m = melody_v1.New()

	s.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		m.HandleRequest(rw, r)
	})

	m.HandleMessage(func(s *melody_v1.Session, msg []byte) {
		var rs []byte
	
		wsmsg := &{{.ServiceType}}WebSocketRequest{}
		err := json.Unmarshal(msg, wsmsg)
		if err != nil {
			rs,_ = json.Marshal({{$svrType}}WebSocketResponse{Action:wsmsg.Action,Code:500,Message:err.Error()})
			s.Write(rs)
			return
		}

		ctx, err := sessionHandler(s.Request.Context(), s.Request.Header.Get("sid"))
		if err != nil {
			rs,_ = json.Marshal({{$svrType}}WebSocketResponse{Action:wsmsg.Action,Code:500,Message:err.Error()})
			s.Write(rs)
			return
		}

		var out interface{}
		if out,err =  wsmsg.HandleMessage(ctx,srv);err !=nil{
			rs,_ = json.Marshal({{$svrType}}WebSocketResponse{Action:wsmsg.Action,Code:500,Message:err.Error()})
			s.Write(rs)
			return
		}

		rs,_ = json.Marshal({{$svrType}}WebSocketResponse{Action:wsmsg.Action,Code:200,Data:out})
		s.Write(rs)
	})
}




func (ws *{{.ServiceType}}WebSocketRequest)HandleMessage(ctx context.Context,srv {{.ServiceType}}Server) (interface{},error){	
	var err error
	switch ws.Action {
		{{- range .Methods}}
	case "/{{$svrName}}/{{.Name}}":
		var in {{.Request}}
		if err = json.Unmarshal([]byte(ws.Input), &in); err != nil {
			return nil,err
		}
		 return srv.{{.Name}}(ctx, &in)
		{{- end}}
	}
	return nil, errors.New("action not found")
}
`

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/helloworld/helloworld.proto
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
}

type methodDesc struct {
	// method
	Name     string
	Num      int
	Request  string
	Response string
	Reply    string
	// WebSocket_rule
	Path         string
	Method       string
	HasVars      bool
	HasBody      bool
	Body         string
	ResponseBody string
}

func (s *serviceDesc) execute() string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("WebSocket").Parse(strings.TrimSpace(WebSocketTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
