package handler

import (
	"errors"
	"fmt"
	"github.com/ismdeep/link/config"
	"github.com/ismdeep/link/model"
	"github.com/ismdeep/link/request"
	"github.com/ismdeep/link/response"
	"github.com/ismdeep/rand"
	"time"
)

func randID() string {
	base := "0123456789abcdefghijklmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
	s := ""
	for i := 0; i < 22; i++ {
		index := rand.Intn(len(base))
		s += base[index : index+1]
	}
	return s
}

type linkHandler struct{}

var Link = &linkHandler{}

func (receiver *linkHandler) Add(req *request.LinkSave) (*response.LinkAdd, error) {

	// 1. check if req.ID is already exists.
	if req.ID != "" {
		var cnt int
		if err := model.DB.Model(&model.Link{}).Where("id=?", req.ID).Error; err != nil {
			return nil, err
		}
		if cnt > 0 {
			return nil, errors.New("already exists")
		}
	}

	// 2. save link
	link := &model.Link{
		ID:          req.ID,
		Target:      req.Target,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if link.ID == "" {
		link.ID = randID()
	}
	if err := model.DB.Create(link).Error; err != nil {
		return nil, err
	}

	return &response.LinkAdd{
		ID:   link.ID,
		Link: fmt.Sprintf("%v/%v", config.Config.Server.Site, link.ID),
	}, nil
}
