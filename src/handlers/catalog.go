package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"db"
)

func Catalog(w http.ResponseWriter, r *http.Request, item httprouter.Params)  {
	db.AddItem("测试",false,2)
}
