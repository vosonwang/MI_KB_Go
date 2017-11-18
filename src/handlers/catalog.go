package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"db"

)

func Catalog(w http.ResponseWriter, r *http.Request, item httprouter.Params)  {
	//u := uuid.NewV4()
	//menu := &db.Menu{u,"测试",false,2,1,uuid.Nil}
	//fmt.Print(*menu)
	db.FindAllTitle()
	db.AddTitle("测试",false,2,"")

}
