package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ContentResponseMeta struct {
	Total int `json:"total"`
}

type Content struct {
	Body string `json:"body"`
}

type ContentResponse struct {
	Meta ContentResponseMeta `json:"meta"`
	Data []Content `json:"data"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {

		return c.File("../templates/index.html")
		
	})

	res := &ContentResponse{
		Meta: ContentResponseMeta{},
		Data: []Content{},
	}

	e.GET("/content", func(c echo.Context) error {
		return c.JSON(http.StatusOK, res)
	})

  
  type Content struct {
    Username string `json:"username"`
    ContentID string `json:"id"`
    ContentBody string `json:"content_body"`
  }

  e.POST("/postTest", func(c echo.Context) (err error) {
    content := new(Content)
    if err = c.Bind(content); err != nil {
      return
    }

    return c.JSON(http.StatusOK, content)
  })

	e.Logger.Fatal(e.Start(":1323"))
}
