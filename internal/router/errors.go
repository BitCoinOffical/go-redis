package router

import "errors"

var ErrorClientDisconnected = errors.New("Client disconnect")
var ErrorEOF = errors.New("EOF")
var ErrorRead = errors.New("Read Error")
var ErrorBadAgument = errors.New("Bad argument")
var ErrorEmptyCommand = errors.New("Empty command")
