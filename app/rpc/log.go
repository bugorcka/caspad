package rpc

import (
	"github.com/casklas/kaspaclassic/infrastructure/logger"
	"github.com/casklas/kaspaclassic/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
