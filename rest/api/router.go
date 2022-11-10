package api

import (
	"context"
	"encoding/json"
	"fmt"
	auth "github.com/alvin-reyes/estuary-auth"
	devicesapi "github.com/application-research/estuary-metrics/rest/api/devices-api"
	"github.com/application-research/estuary-metrics/rest/api/objects-api"
	"github.com/application-research/estuary-metrics/rest/api/reporting-api"
	_ "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/core/generated/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

var (
	_             = time.Second // import time.Second for unknown usage in objects-api
	crudEndpoints map[string]*CrudAPI
)

// CrudAPI describes requests available for tables in the database
type CrudAPI struct {
	Name            string           `json:"name"`
	CreateURL       string           `json:"create_url"`
	RetrieveOneURL  string           `json:"retrieve_one_url"`
	RetrieveManyURL string           `json:"retrieve_many_url"`
	UpdateURL       string           `json:"update_url"`
	DeleteURL       string           `json:"delete_url"`
	FetchDDLURL     string           `json:"fetch_ddl_url"`
	TableInfo       *model.TableInfo `json:"table_info"`
}

// PagedResults results for pages GetAll results.
type PagedResults struct {
	Page         int64       `json:"page"`
	PageSize     int64       `json:"page_size"`
	Data         interface{} `json:"data"`
	TotalRecords int         `json:"total_records"`
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// ConfigGinRouter configure gin router
func ConfigGinRouter(router gin.IRoutes) {

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
	pushapi.ConfigMetricsPushRouter(router)

	//	TODO: blockstore objects-api

	//	devices-api
	devicesapi.ConfigEquinixDevicesRouter(router)

	router.GET("/ddl/:argID", ConverHttpRouterToGin(GetDdl))
	router.GET("/ddl", ConverHttpRouterToGin(GetDdlEndpoints))
	return
}

// ConverHttpRouterToGin wrap httprouter.Handle to gin.HandlerFunc
func ConverHttpRouterToGin(f httprouter.Handle) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params httprouter.Params
		_len := len(c.Params)
		if _len == 0 {
			params = nil
		} else {
			params = ((*[1 << 10]httprouter.Param)(unsafe.Pointer(&c.Params[0])))[:_len]
		}

		f(c.Writer, c.Request, params)
	}
}

func InitializeContext(r *http.Request) (ctx context.Context) {
	if ContextInitializer != nil {
		ctx = ContextInitializer(r)
	} else {
		ctx = r.Context()
	}
	return ctx
}

func ValidateRequest(ctx context.Context, r *http.Request, table string, action model.Action) error {
	if RequestValidator != nil {
		return RequestValidator(ctx, r, table, action)
	}

	return nil
}

type RequestValidatorFunc func(ctx context.Context, r *http.Request, table string, action model.Action) error

var RequestValidator RequestValidatorFunc

type ContextInitializerFunc func(r *http.Request) (ctx context.Context)

var ContextInitializer ContextInitializerFunc

func ReadInt(r *http.Request, param string, v int64) (int64, error) {
	p := r.FormValue(param)
	if p == "" {
		return v, nil
	}

	return strconv.ParseInt(p, 10, 64)
}

func WriteJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func WriteRowsAffected(w http.ResponseWriter, rowsAffected int64) {
	data, _ := json.Marshal(rowsAffected)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func ReadJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, v)
}

func ReturnError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	status := 0
	switch err {
	case dao.ErrNotFound:
		status = http.StatusBadRequest
	case dao.ErrUnableToMarshalJSON:
		status = http.StatusBadRequest
	case dao.ErrUpdateFailed:
		status = http.StatusBadRequest
	case dao.ErrInsertFailed:
		status = http.StatusBadRequest
	case dao.ErrDeleteFailed:
		status = http.StatusBadRequest
	case dao.ErrBadParams:
		status = http.StatusBadRequest
	default:
		status = http.StatusBadRequest
	}
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	objectsapi.SendJSON(w, r, er.Code, er)
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

func parseUint8(ps httprouter.Params, key string) (uint8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return uint8(id), err
	}
	return uint8(id), err
}
func parseUint16(ps httprouter.Params, key string) (uint16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return uint16(id), err
	}
	return uint16(id), err
}
func parseUint32(ps httprouter.Params, key string) (uint32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return uint32(id), err
	}
	return uint32(id), err
}
func parseUint64(ps httprouter.Params, key string) (uint64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return uint64(id), err
	}
	return uint64(id), err
}
func parseInt(ps httprouter.Params, key string) (int, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return -1, err
	}
	return int(id), err
}
func parseInt8(ps httprouter.Params, key string) (int8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return -1, err
	}
	return int8(id), err
}
func parseInt16(ps httprouter.Params, key string) (int16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return -1, err
	}
	return int16(id), err
}
func parseInt32(ps httprouter.Params, key string) (int32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(id), err
}
func ParseInt64(ps httprouter.Params, key string) (int64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 54)
	if err != nil {
		return -1, err
	}
	return id, err
}
func ParseString(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}
func ParseUUID(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}

// GetDdl is a function to get table info for a table in the estuary database
// @Summary Get table info for a table in the estuary database by argID
// @Tags TableInfo
//
// @Description GetDdl is a function to get table info for a table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} objects-api.CrudAPI
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /ddl/{argID} [get]
// http "http://localhost:3030/ddl/xyz" X-Api-User:user123
func GetDdl(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := InitializeContext(r)

	argID := ps.ByName("argID")

	if err := ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		ReturnError(ctx, w, r, err)
		return
	}

	record, ok := crudEndpoints[argID]
	if !ok {
		ReturnError(ctx, w, r, fmt.Errorf("unable to find table: %s", argID))
		return
	}

	WriteJSON(ctx, w, record)
}

// GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the estuary database
// @Summary Gets a list of ddl endpoints available for tables in the estuary database
// @Tags TableInfo
// @Description GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the estuary database
// @Accept  json
// @Produce  json
// @Success 200 {object} objects-api.CrudAPI
// @Router /ddl [get]
// http "http://localhost:3030/ddl" X-Api-User:user123
func GetDdlEndpoints(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := InitializeContext(r)

	if err := ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		ReturnError(ctx, w, r, err)
		return
	}

	WriteJSON(ctx, w, crudEndpoints)
}

func init() {
	crudEndpoints = make(map[string]*CrudAPI)

	var tmp *CrudAPI

	tmp = &CrudAPI{
		Name:            "auth_tokens",
		CreateURL:       "/authtokens",
		RetrieveOneURL:  "/authtokens",
		RetrieveManyURL: "/authtokens",
		UpdateURL:       "/authtokens",
		DeleteURL:       "/authtokens",
		FetchDDLURL:     "/ddl/auth_tokens",
	}

	tmp.TableInfo, _ = model.GetTableInfo("auth_tokens")
	crudEndpoints["auth_tokens"] = tmp

	tmp = &CrudAPI{
		Name:            "autoretrieves",
		CreateURL:       "/autoretrieves",
		RetrieveOneURL:  "/autoretrieves",
		RetrieveManyURL: "/autoretrieves",
		UpdateURL:       "/autoretrieves",
		DeleteURL:       "/autoretrieves",
		FetchDDLURL:     "/ddl/autoretrieves",
	}

	tmp.TableInfo, _ = model.GetTableInfo("autoretrieves")
	crudEndpoints["autoretrieves"] = tmp

	tmp = &CrudAPI{
		Name:            "collection_refs",
		CreateURL:       "/collectionrefs",
		RetrieveOneURL:  "/collectionrefs",
		RetrieveManyURL: "/collectionrefs",
		UpdateURL:       "/collectionrefs",
		DeleteURL:       "/collectionrefs",
		FetchDDLURL:     "/ddl/collection_refs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("collection_refs")
	crudEndpoints["collection_refs"] = tmp

	tmp = &CrudAPI{
		Name:            "collections",
		CreateURL:       "/collections",
		RetrieveOneURL:  "/collections",
		RetrieveManyURL: "/collections",
		UpdateURL:       "/collections",
		DeleteURL:       "/collections",
		FetchDDLURL:     "/ddl/collections",
	}

	tmp.TableInfo, _ = model.GetTableInfo("collections")
	crudEndpoints["collections"] = tmp

	tmp = &CrudAPI{
		Name:            "content_deals",
		CreateURL:       "/contentdeals",
		RetrieveOneURL:  "/contentdeals",
		RetrieveManyURL: "/contentdeals",
		UpdateURL:       "/contentdeals",
		DeleteURL:       "/contentdeals",
		FetchDDLURL:     "/ddl/content_deals",
	}

	tmp.TableInfo, _ = model.GetTableInfo("content_deals")
	crudEndpoints["content_deals"] = tmp

	tmp = &CrudAPI{
		Name:            "contents",
		CreateURL:       "/contents",
		RetrieveOneURL:  "/contents",
		RetrieveManyURL: "/contents",
		UpdateURL:       "/contents",
		DeleteURL:       "/contents",
		FetchDDLURL:     "/ddl/contents",
	}

	tmp.TableInfo, _ = model.GetTableInfo("contents")
	crudEndpoints["contents"] = tmp

	tmp = &CrudAPI{
		Name:            "dealers",
		CreateURL:       "/dealers",
		RetrieveOneURL:  "/dealers",
		RetrieveManyURL: "/dealers",
		UpdateURL:       "/dealers",
		DeleteURL:       "/dealers",
		FetchDDLURL:     "/ddl/dealers",
	}

	tmp.TableInfo, _ = model.GetTableInfo("dealers")
	crudEndpoints["dealers"] = tmp

	tmp = &CrudAPI{
		Name:            "dfe_records",
		CreateURL:       "/dferecords",
		RetrieveOneURL:  "/dferecords",
		RetrieveManyURL: "/dferecords",
		UpdateURL:       "/dferecords",
		DeleteURL:       "/dferecords",
		FetchDDLURL:     "/ddl/dfe_records",
	}

	tmp.TableInfo, _ = model.GetTableInfo("dfe_records")
	crudEndpoints["dfe_records"] = tmp

	tmp = &CrudAPI{
		Name:            "invite_codes",
		CreateURL:       "/invitecodes",
		RetrieveOneURL:  "/invitecodes",
		RetrieveManyURL: "/invitecodes",
		UpdateURL:       "/invitecodes",
		DeleteURL:       "/invitecodes",
		FetchDDLURL:     "/ddl/invite_codes",
	}

	tmp.TableInfo, _ = model.GetTableInfo("invite_codes")
	crudEndpoints["invite_codes"] = tmp

	tmp = &CrudAPI{
		Name:            "miner_storage_asks",
		CreateURL:       "/minerstorageasks",
		RetrieveOneURL:  "/minerstorageasks",
		RetrieveManyURL: "/minerstorageasks",
		UpdateURL:       "/minerstorageasks",
		DeleteURL:       "/minerstorageasks",
		FetchDDLURL:     "/ddl/miner_storage_asks",
	}

	tmp.TableInfo, _ = model.GetTableInfo("miner_storage_asks")
	crudEndpoints["miner_storage_asks"] = tmp

	tmp = &CrudAPI{
		Name:            "obj_refs",
		CreateURL:       "/objrefs",
		RetrieveOneURL:  "/objrefs",
		RetrieveManyURL: "/objrefs",
		UpdateURL:       "/objrefs",
		DeleteURL:       "/objrefs",
		FetchDDLURL:     "/ddl/obj_refs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("obj_refs")
	crudEndpoints["obj_refs"] = tmp

	tmp = &CrudAPI{
		Name:            "objects",
		CreateURL:       "/objects",
		RetrieveOneURL:  "/objects",
		RetrieveManyURL: "/objects",
		UpdateURL:       "/objects",
		DeleteURL:       "/objects",
		FetchDDLURL:     "/ddl/objects",
	}

	tmp.TableInfo, _ = model.GetTableInfo("objects")
	crudEndpoints["objects"] = tmp

	tmp = &CrudAPI{
		Name:            "piece_comm_records",
		CreateURL:       "/piececommrecords",
		RetrieveOneURL:  "/piececommrecords",
		RetrieveManyURL: "/piececommrecords",
		UpdateURL:       "/piececommrecords",
		DeleteURL:       "/piececommrecords",
		FetchDDLURL:     "/ddl/piece_comm_records",
	}

	tmp.TableInfo, _ = model.GetTableInfo("piece_comm_records")
	crudEndpoints["piece_comm_records"] = tmp

	tmp = &CrudAPI{
		Name:            "piece_comm_records_backup_20220125",
		CreateURL:       "/piececommrecordsbackup20220125",
		RetrieveOneURL:  "/piececommrecordsbackup20220125",
		RetrieveManyURL: "/piececommrecordsbackup20220125",
		UpdateURL:       "/piececommrecordsbackup20220125",
		DeleteURL:       "/piececommrecordsbackup20220125",
		FetchDDLURL:     "/ddl/piece_comm_records_backup_20220125",
	}

	tmp.TableInfo, _ = model.GetTableInfo("piece_comm_records_backup_20220125")
	crudEndpoints["piece_comm_records_backup_20220125"] = tmp

	tmp = &CrudAPI{
		Name:            "proposal_records",
		CreateURL:       "/proposalrecords",
		RetrieveOneURL:  "/proposalrecords",
		RetrieveManyURL: "/proposalrecords",
		UpdateURL:       "/proposalrecords",
		DeleteURL:       "/proposalrecords",
		FetchDDLURL:     "/ddl/proposal_records",
	}

	tmp.TableInfo, _ = model.GetTableInfo("proposal_records")
	crudEndpoints["proposal_records"] = tmp

	tmp = &CrudAPI{
		Name:            "retrieval_failure_records",
		CreateURL:       "/retrievalfailurerecords",
		RetrieveOneURL:  "/retrievalfailurerecords",
		RetrieveManyURL: "/retrievalfailurerecords",
		UpdateURL:       "/retrievalfailurerecords",
		DeleteURL:       "/retrievalfailurerecords",
		FetchDDLURL:     "/ddl/retrieval_failure_records",
	}

	tmp.TableInfo, _ = model.GetTableInfo("retrieval_failure_records")
	crudEndpoints["retrieval_failure_records"] = tmp

	tmp = &CrudAPI{
		Name:            "retrieval_success_records",
		CreateURL:       "/retrievalsuccessrecords",
		RetrieveOneURL:  "/retrievalsuccessrecords",
		RetrieveManyURL: "/retrievalsuccessrecords",
		UpdateURL:       "/retrievalsuccessrecords",
		DeleteURL:       "/retrievalsuccessrecords",
		FetchDDLURL:     "/ddl/retrieval_success_records",
	}

	tmp.TableInfo, _ = model.GetTableInfo("retrieval_success_records")
	crudEndpoints["retrieval_success_records"] = tmp

	tmp = &CrudAPI{
		Name:            "shuttles",
		CreateURL:       "/shuttles",
		RetrieveOneURL:  "/shuttles",
		RetrieveManyURL: "/shuttles",
		UpdateURL:       "/shuttles",
		DeleteURL:       "/shuttles",
		FetchDDLURL:     "/ddl/shuttles",
	}

	tmp.TableInfo, _ = model.GetTableInfo("shuttles")
	crudEndpoints["shuttles"] = tmp

	tmp = &CrudAPI{
		Name:            "storage_miners",
		CreateURL:       "/storageminers",
		RetrieveOneURL:  "/storageminers",
		RetrieveManyURL: "/storageminers",
		UpdateURL:       "/storageminers",
		DeleteURL:       "/storageminers",
		FetchDDLURL:     "/ddl/storage_miners",
	}

	tmp.TableInfo, _ = model.GetTableInfo("storage_miners")
	crudEndpoints["storage_miners"] = tmp

	tmp = &CrudAPI{
		Name:            "users",
		CreateURL:       "/users",
		RetrieveOneURL:  "/users",
		RetrieveManyURL: "/users",
		UpdateURL:       "/users",
		DeleteURL:       "/users",
		FetchDDLURL:     "/ddl/users",
	}

	tmp.TableInfo, _ = model.GetTableInfo("users")
	crudEndpoints["users"] = tmp

}
