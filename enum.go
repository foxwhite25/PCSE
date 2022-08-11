package main

type savePool int8

const (
	ManualSave savePool = iota + 1
	AutoSave
)
