package models

import (
	"github.com/whkelvin/stamp/pkg/api/generated/models"
)

//type Request struct {
//	Size int `query:"size"`
//	Page int `query:"page"`
//}

type Request = models.GetRecentPostsParams
