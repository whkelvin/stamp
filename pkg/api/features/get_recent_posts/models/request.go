package models

type Request struct {
	Size int `query:"size"`
	Page int `query:"page"`
}
