package controllers

import (
	"fmt"
	"html/template"

	"github.com/revel/revel"
	"github.com/russross/blackfriday"
)

func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*GormController).Begin, revel.BEFORE)
	revel.InterceptMethod(PublicApp.AddUser, revel.BEFORE)
	revel.InterceptMethod(App.checkUser, revel.BEFORE)
	revel.InterceptMethod((*GormController).Commit, revel.AFTER)
	revel.InterceptMethod((*GormController).Rollback, revel.FINALLY)

	revel.TemplateFuncs["markdown"] = func(str interface{}) string {
		s := blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", str)))
		return string(template.HTML(s))
	}
}
