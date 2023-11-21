package resp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"one_ser.com/errno"
)

type H struct {
	Code  int
	Msg   string
	Data  interface{}
	Rows  interface{}
	Total interface{}
}

func Res(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}

func ResList(w http.ResponseWriter, code int, data interface{}, total interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}

func ResFail(w http.ResponseWriter, msg string) {
	code:=errno.ERROR
	Res(w, code, nil, msg)
}

func ResOK(w http.ResponseWriter, data interface{}, msg string) {
	code:=errno.SUCCESS
	Res(w, code, data, msg)
}

func ResOKList(w http.ResponseWriter, data interface{}, total interface{}) {
	code:=errno.SUCCESS
	ResList(w, code, data, total)
}
