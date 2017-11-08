package main

type SDNPlugin interface {
	Setup(map[string]interface{})(error)
	AllocateNic()(int,error)
	DeleteNic(int)(error)
}