package app
import (
	
	"net/http"
	"github.com/jirawat050/GolangWeb/view"
	"github.com/jirawat050/GolangWeb/model"

)
func adminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost{
	http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
	return
	}
	view.AdminLogin(w, nil)
}
func adminList(w http.ResponseWriter, r *http.Request) {
	view.AdminList(w, nil)
}
func adminCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost{
		n := model.News {
		Title: r.FormValue("title"),
		Detail: r.FormValue("detail"),
	}
	model.CreateNews(n)
		http.Redirect(w, r, "/admin/create", http.StatusSeeOther)
	

	}
	view.AdminCreate(w, nil)
}
func adminEdit(w http.ResponseWriter, r *http.Request) {
	view.AdminEdit(w, nil)
}