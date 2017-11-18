package db

import "github.com/satori/go.uuid"

type Title struct{
	Id uuid.UUID
	Title string
	Expand bool
	Sort int8
	Level int8
	Parent_id uuid.UUID
}

type Menu []Title