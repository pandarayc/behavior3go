package core

type NodeStatus int32
const (
	STATUS_SUCCESS NodeStatus =  iota
	STATUS_FAILURE
	STATUS_RUNNING 
	STATUS_ERROR
)