package models

type Status int

const (
	NOT_STARTED Status = iota
	IN_PROGRESS
	FINISHED
	DISCONTINUED
)
