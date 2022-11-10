package route

import (
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
func ConfigRouter(router gin.IRoutes) {

	router.Use(func(c *gin.Context) {
		// authenticate here
		authServer := auth.Init().SetDB(dao.DB).Connect()
		authorizationString := c.GetHeader("Authorization")
		authParts := strings.Split(authorizationString, " ")
		//	authparts
		if len(authParts) != 2 {
			http.Error(c.Writer, "invalid authorization header", http.StatusUnauthorized)

		}
		// 	tokens
		token, err := authServer.CheckAuthorizationToken(authParts[1], 100)
		if err != nil {
			http.Error(c.Writer, "invalid authorization token", http.StatusUnauthorized)
		}
		//	only admins
		if token.Perm < 100 {
			http.Error(c.Writer, "permission denied", http.StatusUnauthorized)

		}
	})

	//	all estuary objects objects-api
	objectsapi.ConfigAuthTokensRouter(router)
	objectsapi.ConfigAutoretrievesRouter(router)
	objectsapi.ConfigCollectionRefsRouter(router)
	objectsapi.ConfigCollectionsRouter(router)
	objectsapi.ConfigContentDealsRouter(router)
	objectsapi.ConfigContentsRouter(router)
	objectsapi.ConfigDealersRouter(router)
	objectsapi.ConfigDfeRecordsRouter(router)
	objectsapi.ConfigInviteCodesRouter(router)
	objectsapi.ConfigMinerStorageAsksRouter(router)
	objectsapi.ConfigObjRefsRouter(router)
	objectsapi.ConfigObjectsRouter(router)
	objectsapi.ConfigPieceCommRecordsRouter(router)
	objectsapi.ConfigProposalRecordsRouter(router)
	objectsapi.ConfigRetrievalFailureRecordsRouter(router)
	objectsapi.ConfigRetrievalSuccessRecordsRouter(router)
	objectsapi.ConfigShuttlesRouter(router)
	objectsapi.ConfigStorageMinersRouter(router)
	objectsapi.ConfigUsersRouter(router)

	//	reporting objects-api
	reportingapi.ConfigMetricsPushRouter(router)

	//	TODO: blockstore objects-api

	//	devices-api
	devicesapi.ConfigEquinixDevicesRouter(router)

	//	stats-api
	statsapi.ConfigStatsRouter(router)

	objectsapi.ConfigDDLRouter(router)
	return
}
