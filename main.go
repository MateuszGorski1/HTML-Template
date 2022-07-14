package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorski.mateusz/htmltemplate/person"
	"gorski.mateusz/htmltemplate/templatehtml"
)

func main() {
	const tplPerson = `<html>
	<head>
		<meta charset="UTF-8">
		<title>Person</title>
		<link rel="stylesheet" href="/css/person.css">
	</head>
	<body>
	<div class="wrapper">
		<div class="fancy">{{.Name}}</div>
		<div class="fancy">{{.Surname}}</div>
		<div class="fancy">{{.Age}}</div>
	</div>
	</body>
</html>`

	Person := person.CreatePerson("John", "Davis", 23)
	templatehtml.GenerateTemplate(tplPerson, Person)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/css", "./assets/css")
	r.SetTrustedProxies([]string{"135.245.48.34"})
	r.LoadHTMLFiles("./assets/person.html")
	r.GET("/person", func(c *gin.Context) {
		c.HTML(http.StatusOK, "person.html", nil)
	})

	r.Run(":8080")
}
