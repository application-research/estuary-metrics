package objectsapi

import (
	"fmt"
	"github.com/application-research/estuary-metrics/core/generated/model"
	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

var (
	_             = time.Second // import time.Second for unknown usage in objects-api
	CrudEndpoints map[string]*CrudAPI
)

func ConfigDDLRouter(router gin.IRoutes) {
	//
	getCrudEndpoints() // get the endpoints
	router.GET("/ddl/:id", api.ConvertHttpRouterToGin(GetDdl))
	router.GET("/ddl", api.ConvertHttpRouterToGin(GetDdlEndpoints))

}

// GetDdl is a function to get table info for a table in the estuary database
// @Summary Get table info for a table in the estuary database by argID
// @Tags TableInfo
//
// @Description GetDdl is a function to get table info for a table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} objectsapi.CrudAPI
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /ddl/{argID} [get]
// http "http://localhost:3030/ddl/xyz" X-Api-User:user123
func GetDdl(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	argID := ps.ByName("argID")

	if err := api.ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, ok := CrudEndpoints[argID]
	if !ok {
		api.ReturnError(ctx, w, r, fmt.Errorf("unable to find table: %s", argID))
		return
	}

	api.WriteJSON(ctx, w, record)
}

// GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the estuary database
// @Summary Gets a list of ddl endpoints available for tables in the estuary database
// @Tags TableInfo
// @Description GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the estuary database
// @Accept  json
// @Produce  json
// @Success 200 {object} objectsapi.CrudAPI
// @Router /ddl [get]
// http "http://localhost:3030/ddl" X-Api-User:user123
func GetDdlEndpoints(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	if err := api.ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, CrudEndpoints)
}

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

func getCrudEndpoints() map[string]*CrudAPI {

	CrudEndpoints = make(map[string]*CrudAPI)

	var tmp *CrudAPI

	tmp = &CrudAPI{
		Name:            "auth_tokens",
		RetrieveOneURL:  "/authtokens",
		RetrieveManyURL: "/authtokens",
		FetchDDLURL:     "/ddl/auth_tokens",
	}

	tmp.TableInfo, _ = model.GetTableInfo("auth_tokens")
	CrudEndpoints["auth_tokens"] = tmp

	tmp = &CrudAPI{
		Name:            "autoretrieves",
		RetrieveOneURL:  "/autoretrieves",
		RetrieveManyURL: "/autoretrieves",
		FetchDDLURL:     "/ddl/autoretrieves",
	}

	tmp.TableInfo, _ = model.GetTableInfo("autoretrieves")
	CrudEndpoints["autoretrieves"] = tmp

	tmp = &CrudAPI{
		Name:            "collection_refs",
		RetrieveOneURL:  "/collectionrefs",
		RetrieveManyURL: "/collectionrefs",
		FetchDDLURL:     "/ddl/collection_refs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("collection_refs")
	CrudEndpoints["collection_refs"] = tmp

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
	CrudEndpoints["collections"] = tmp

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
	CrudEndpoints["content_deals"] = tmp

	tmp = &CrudAPI{
		Name:            "contents",
		RetrieveOneURL:  "/contents",
		RetrieveManyURL: "/contents",
		FetchDDLURL:     "/ddl/contents",
	}

	tmp.TableInfo, _ = model.GetTableInfo("contents")
	CrudEndpoints["contents"] = tmp

	tmp = &CrudAPI{
		Name:            "dealers",
		RetrieveOneURL:  "/dealers",
		RetrieveManyURL: "/dealers",
		FetchDDLURL:     "/ddl/dealers",
	}

	tmp.TableInfo, _ = model.GetTableInfo("dealers")
	CrudEndpoints["dealers"] = tmp

	tmp = &CrudAPI{
		Name:            "dfe_records",
		RetrieveOneURL:  "/dferecords",
		RetrieveManyURL: "/dferecords",
		FetchDDLURL:     "/ddl/dfe_records",
	}

	tmp.TableInfo, _ = model.GetTableInfo("dfe_records")
	CrudEndpoints["dfe_records"] = tmp

	tmp = &CrudAPI{
		Name:            "invite_codes",
		RetrieveOneURL:  "/invitecodes",
		RetrieveManyURL: "/invitecodes",
		FetchDDLURL:     "/ddl/invite_codes",
	}

	tmp.TableInfo, _ = model.GetTableInfo("invite_codes")
	CrudEndpoints["invite_codes"] = tmp

	tmp = &CrudAPI{
		Name:            "miner_storage_asks",
		RetrieveOneURL:  "/minerstorageasks",
		RetrieveManyURL: "/minerstorageasks",
		FetchDDLURL:     "/ddl/miner_storage_asks",
	}

	tmp.TableInfo, _ = model.GetTableInfo("miner_storage_asks")
	CrudEndpoints["miner_storage_asks"] = tmp

	tmp = &CrudAPI{
		Name:            "obj_refs",
		RetrieveOneURL:  "/objrefs",
		RetrieveManyURL: "/objrefs",
		FetchDDLURL:     "/ddl/obj_refs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("obj_refs")
	CrudEndpoints["obj_refs"] = tmp

	tmp = &CrudAPI{
		Name:            "objects",
		RetrieveOneURL:  "/objects",
		RetrieveManyURL: "/objects",
		FetchDDLURL:     "/ddl/objects",
	}

	tmp.TableInfo, _ = model.GetTableInfo("objects")
	CrudEndpoints["objects"] = tmp

	tmp = &CrudAPI{
		Name:            "piece_comm_records",
		RetrieveOneURL:  "/piececommrecords",
		RetrieveManyURL: "/piececommrecords",
		FetchDDLURL:     "/ddl/piece_comm_records",
	}

	tmp.TableInfo, _ = model.GetTableInfo("piece_comm_records")
	CrudEndpoints["piece_comm_records"] = tmp

	tmp = &CrudAPI{
		Name:            "piece_comm_records_backup_20220125",
		RetrieveOneURL:  "/piececommrecordsbackup20220125",
		RetrieveManyURL: "/piececommrecordsbackup20220125",
		FetchDDLURL:     "/ddl/piece_comm_records_backup_20220125",
	}

	tmp.TableInfo, _ = model.GetTableInfo("piece_comm_records_backup_20220125")
	CrudEndpoints["piece_comm_records_backup_20220125"] = tmp

	tmp = &CrudAPI{
		Name:            "proposal_records",
		RetrieveOneURL:  "/proposalrecords",
		RetrieveManyURL: "/proposalrecords",
		FetchDDLURL:     "/ddl/proposal_records",
	}

	tmp.TableInfo, _ = model.GetTableInfo("proposal_records")
	CrudEndpoints["proposal_records"] = tmp

	tmp = &CrudAPI{
		Name:            "retrieval_failure_records",
		RetrieveOneURL:  "/retrievalfailurerecords",
		RetrieveManyURL: "/retrievalfailurerecords",
		FetchDDLURL:     "/ddl/retrieval_failure_records",
	}

	tmp.TableInfo, _ = model.GetTableInfo("retrieval_failure_records")
	CrudEndpoints["retrieval_failure_records"] = tmp

	tmp = &CrudAPI{
		Name:            "retrieval_success_records",
		RetrieveOneURL:  "/retrievalsuccessrecords",
		RetrieveManyURL: "/retrievalsuccessrecords",
		FetchDDLURL:     "/ddl/retrieval_success_records",
	}

	tmp.TableInfo, _ = model.GetTableInfo("retrieval_success_records")
	CrudEndpoints["retrieval_success_records"] = tmp

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
	CrudEndpoints["shuttles"] = tmp

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
	CrudEndpoints["storage_miners"] = tmp

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
	CrudEndpoints["users"] = tmp

	return CrudEndpoints
}
