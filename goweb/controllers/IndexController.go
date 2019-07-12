package controllers

import (
	"fmt"
	"net/http"
)

type IndexController struct {
}

func (c *IndexController) Index(w http.ResponseWriter, req *http.Request) {
	fmt.Println("IndexController::Index")
}
