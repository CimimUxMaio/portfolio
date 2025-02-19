package main

import (
	"log"
	"net/http"

	"github.com/CimimUxMaio/portfolio/model"
	"github.com/CimimUxMaio/portfolio/templates"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())

	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		renderComponent(c, templates.Index())
	})

	info := r.Group("/html/info")
	info.GET("/whoami", withPortfolio(func(c *gin.Context, portfolio model.Portfolio) {
		renderComponent(c, templates.WhoAmI(portfolio.Profile))
	}))

	info.GET("/mywork", withPortfolio(func(c *gin.Context, portfolio model.Portfolio) {
		renderComponent(c, templates.MyWork(portfolio.Projects))
	}))

	info.GET("/contact", withPortfolio(func(c *gin.Context, portfolio model.Portfolio) {
		renderComponent(c, templates.Contact(portfolio.Contact))
	}))

	info.GET("/projects/:title", withPortfolio(func(c *gin.Context, portfolio model.Portfolio) {
		title := c.Param("title")

		var targetProject *model.Project = nil
		for _, project := range portfolio.Projects {
			if project.Title == title {
				targetProject = &project
				break
			}
		}

		if targetProject == nil {
			c.Status(http.StatusNotFound)
			return
		}

		renderComponent(c, templates.Project(*targetProject))
	}))

	commands := r.Group("/html/commands")
	commands.GET("/whoami", command("whoami", "/html/info/whoami", ""))
	commands.GET("/mywork", command("ls ~/projects/ ", "/html/info/mywork", ""))
	commands.GET("/contact", command("cat /usr/contact-info.json", "/html/info/contact", ""))
	commands.GET("/projects/:title", func(c *gin.Context) {
		title := c.Param("title")
		command("cat ~/projects/"+title+"/README.md", "/html/info/projects/"+title, "")(c)
	})

	commands.GET("/clear", command("clear", "", "clearContent();"))

	error := r.Run()
	if error != nil {
		panic(error)
	}
}

func withPortfolio(handler func(*gin.Context, model.Portfolio)) gin.HandlerFunc {
	return func(c *gin.Context) {
		portfolio, err := model.LoadPortfolio()
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		handler(c, portfolio)
	}
}

func renderComponent(ctx *gin.Context, component templ.Component) {
	ctx.Status(http.StatusOK)
	err := component.Render(ctx.Request.Context(), ctx.Writer)
	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
	}
}

func command(text string, requestUrl string, afterRequest string) gin.HandlerFunc {
	return func(c *gin.Context) {
		commandInfo := templates.CommandInfo{
			Text:         text,
			RequestUrl:   requestUrl,
			AfterRequest: afterRequest,
		}
		renderComponent(c, templates.Command(commandInfo))
	}
}
