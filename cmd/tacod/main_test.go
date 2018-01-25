package main

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/sul-dlss-labs/taco"
	"github.com/sul-dlss-labs/taco/config"
)

func TestRetrieveHappyPath(t *testing.T) {
	config.Init("../../config/test.yaml")
	rt, _ := taco.NewRuntime(viper.GetViper())
	handler := buildAPI(rt).Serve(nil)
	r := gofight.New()
	r.GET("/v1/resource/99").
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestRetrieveNotFound(t *testing.T) {
	config.Init("../../config/test.yaml")
	rt, _ := taco.NewRuntime(viper.GetViper())
	handler := buildAPI(rt).Serve(nil)
	r := gofight.New()
	r.GET("/v1/resource/100").
		Run(handler, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, 404, r.Code)
		})
}
