package qbt

import "strings"

type ActionInterface interface {
	Arguments() []string
}

func NewActionFromString(action string) *Action {
	return &Action{action: action}
}

type Action struct {
	action string
}

func (a *Action) Arguments() []string {
	return strings.Split(a.action, " ")
}
