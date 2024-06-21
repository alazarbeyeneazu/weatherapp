package initiator

import (
	pb "github.com/alazarbeyeneazu/common/api"
	"github.com/alazarbeyeneazu/gateway/internals/glue/auth"
	"github.com/alazarbeyeneazu/gateway/internals/glue/weather"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func initRoute(
	group *gin.RouterGroup,
	handler handler,
	log zap.Logger,
	authGRPCClient pb.AuthServiceClient,

) {
	weather.Init(group, log, handler.weather, authGRPCClient)
	auth.Init(group, log, handler.auth)

}
