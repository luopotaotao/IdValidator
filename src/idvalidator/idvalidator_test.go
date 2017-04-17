package idvalidator

import (
	"testing"
)

func TestValidate(t *testing.T) {
	id:="510723198006202551"
	_,err:=Validate(id)
	if err!=nil{
		t.Error(err)
	}
}