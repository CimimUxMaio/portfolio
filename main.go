package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())

	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	r.SetFuncMap(customFuncMap()) // Setup custom function for templates

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", withPortfolioData("index.html"))

	info := r.Group("/html/info")
	info.GET("/whoami", withPortfolioData("whoami.html"))
	info.GET("/mywork", withPortfolioData("mywork.html"))
	info.GET("/contact", withPortfolioData("contact.html"))

	commands := r.Group("/html/commands")
	commands.GET("/whoami", command("whoami", "/html/info/whoami", ""))
	commands.GET("/mywork", command("cat ./projects/summary.md ", "/html/info/mywork", ""))
	commands.GET("/contact", command("cat contact_info.json", "/html/info/contact", ""))

	commands.GET("/clear", command("clear", "", "clearContent();"))

	error := r.Run()
	if error != nil {
		panic(error)
	}
}

func getPortfolioData() gin.H {
	b, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatalf("Failed to read data.json:\n %v\n", err)
	}

	data := map[string]any{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Fatalf("Failed to read data.json:\n %v\n", err)
	}

	return gin.H(data)
}

func withPortfolioData(template string) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := getPortfolioData()
		c.HTML(http.StatusOK, template, gin.H(data))
	}
}

func command(command string, infoUrl string, onExec string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "command.html", commandInfo(command, infoUrl, onExec))
	}
}

func commandInfo(command string, infoUrl string, onExec string) gin.H {
	typingSteps := len(command)
	var charPerSec float32 = 12
	typingDur := float32(typingSteps) / charPerSec
	return gin.H{"command": command, "typingDur": typingDur, "typingSteps": typingSteps, "infoUrl": infoUrl, "onExec": onExec}
}

func customFuncMap() template.FuncMap {
	return template.FuncMap{
		"dict": func(values ...any) (map[string]any, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("Invalid dict call")
			}

			dict := make(map[string]any, len(values)/2)

			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("Dict keys must be strings")
				}

				dict[key] = values[i+1]
			}

			return dict, nil
		},
	}
}
