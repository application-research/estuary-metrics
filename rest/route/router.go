package route

import (
	"fmt"
	auth "github.com/alvin-reyes/estuary-auth"
	"github.com/application-research/estuary-metrics/core/dao"
	devicesapi "github.com/application-research/estuary-metrics/rest/api/devices-api"
	objectsapi "github.com/application-research/estuary-metrics/rest/api/objects-api"
	"github.com/application-research/estuary-metrics/rest/api/reporting-api"
	statsapi "github.com/application-research/estuary-metrics/rest/api/stats-api"
	"github.com/gin-gonic/gin"
	_ "github.com/satori/go.uuid"
	"net/http"
	"strings"
)

// ConfigRouter configure gin route
// @Security BearerAuth
func ConfigRouter(router *gin.Engine) {
	group := router.Group("/api/v1")
	unprotectedRouter := router.Group("/api/v1")
	protectedRouter := group.Use(func(c *gin.Context) {
		// authenticate here
		authServer := auth.Init().SetDB(dao.DB).Connect()
		authorizationString := c.GetHeader("Authorization")
		authParts := strings.Split(authorizationString, " ")
		fmt.Println("authParts: ", authParts)
		//	authparts
		if len(authParts) != 2 {
			http.Error(c.Writer, "invalid authorization header", http.StatusUnauthorized)
		}
		// 	tokens
		token, err := authServer.CheckAuthorizationToken(authParts[1], 100) // user needs to be 100 (super admin to access)
		if err != nil {
			http.Error(c.Writer, "invalid authorization token", http.StatusUnauthorized)
		}
		//	only admins
		if token.Perm < 100 {
			http.Error(c.Writer, "permission denied", http.StatusUnauthorized)

		}
	})

	//	all estuary objects objects-api
	objectsapi.ConfigAuthTokensRouter(protectedRouter)
	objectsapi.ConfigAutoretrievesRouter(protectedRouter)
	objectsapi.ConfigUnProtectedAutoRetrievesRouter(unprotectedRouter)

	objectsapi.ConfigCollectionRefsRouter(protectedRouter)
	objectsapi.ConfigCollectionsRouter(protectedRouter)

	objectsapi.ConfigContentDealsRouter(protectedRouter)
	objectsapi.ConfigUnProtectedContentDealsRouter(unprotectedRouter)

	objectsapi.ConfigContentsRouter(protectedRouter)
	objectsapi.ConfigUnProtectedContentsRouter(unprotectedRouter)

	objectsapi.ConfigDealersRouter(protectedRouter)
	objectsapi.ConfigDfeRecordsRouter(protectedRouter)
	objectsapi.ConfigInviteCodesRouter(protectedRouter)
	objectsapi.ConfigMinerStorageAsksRouter(protectedRouter)
	objectsapi.ConfigObjRefsRouter(protectedRouter)
	objectsapi.ConfigObjectsRouter(protectedRouter)
	objectsapi.ConfigPieceCommRecordsRouter(protectedRouter)
	objectsapi.ConfigProposalRecordsRouter(protectedRouter)
	objectsapi.ConfigRetrievalFailureRecordsRouter(protectedRouter)
	objectsapi.ConfigRetrievalSuccessRecordsRouter(protectedRouter)
	objectsapi.ConfigShuttlesRouter(protectedRouter)
	objectsapi.ConfigStorageMinersRouter(protectedRouter)
	objectsapi.ConfigUsersRouter(protectedRouter)

	//	reporting objects-api
	reportingapi.ConfigMetricsPushRouter(unprotectedRouter)

	//	devices-api
	devicesapi.ConfigEquinixDevicesRouter(unprotectedRouter)
	devicesapi.ConfigAwsDevicesRouter(unprotectedRouter)

	//	stats-api
	statsapi.ConfigStatsRouter(unprotectedRouter)
	statsapi.ConfigRankingRoute(unprotectedRouter)
	statsapi.ConfigDistributionRoute(unprotectedRouter)
	statsapi.ConfigHeartbeatRoute(unprotectedRouter)
	statsapi.ConfigLocationRoute(unprotectedRouter)

	//	TODO: Blockstore API

	//	DDL
	objectsapi.ConfigDDLRouter(unprotectedRouter)
	return
}
