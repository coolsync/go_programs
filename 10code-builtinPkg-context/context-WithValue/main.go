package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func main() {
    // error massages encountered, use http.Handle
	// cannot use injectMsgID(sayHoloHandler) (value of type http.Handler) 
    // as func(http.ResponseWriter, *http.Request) value in argument to http.HandleFunc
    
    sayHoloHandler := http.HandlerFunc(sayHolo)
    // http.HandleFunc(pattern string, handler func(ResponseWriter, *Request))
    http.Handle("/holo", injectMsgID(sayHoloHandler))
    http.ListenAndServe(":9090", nil)
}

func injectMsgID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
        // 生成 msg id
        MsgID := uuid.New().String()
   
        // 注入 request_scoped
        ctx := context.WithValue(r.Context(), "msg_id", MsgID)
        // copy 新生成的 context
        // req := r.WithContext(ctx)
        req := r.Clone(ctx)
        // pass copy context to sayHolo
        next.ServeHTTP(rw, req)
    })
}

func sayHolo(w http.ResponseWriter, r *http.Request) {
    // 获取 cp context value, 并使用确定类型 获取 具体值
    MsgId := ""

    if m := r.Context().Value("msg_id"); m != nil {
        if val, ok := m.(string); ok {
            MsgId = val
        }
    }

    w.Header().Add("msg_id", MsgId)
    w.Write([]byte("holo web"))
}