package main

import (
	"gin-rest/handler"
	"gin-rest/infra"
	"gin-rest/usecase"
	"fmt"

	appvalidator "gin-rest/handler/validator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

