package main

import (
	"log"

	"github.com/application-research/estuary-metrics/core/generated/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {

	// cli get .env settings
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	dbHost, okHost := viper.Get("DB_HOST").(string)
	dbUser, okUser := viper.Get("DB_USER").(string)
	dbPass, okPass := viper.Get("DB_PASS").(string)
	dbName, okName := viper.Get("DB_NAME").(string)
	dbPort, okPort := viper.Get("DB_PORT").(string)
	if !okHost || !okUser || !okPass || !okName || !okPort {
		log.Fatalf("Error while reading database config")
	}

	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "generated/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
		//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		/* FieldCoverable: true,*/
		// if you want generated field with unsigned integer type, set FieldSignable true
		/* FieldSignable: true,*/
		//if you want to generated index tags from database, set FieldWithIndexTag true
		/* FieldWithIndexTag: true,*/
		//if you want to generated type tags from database, set FieldWithTypeTag true
		/* FieldWithTypeTag: true,*/
		//if you need unit tests for query code, set WithUnitTest true
		/* WithUnitTest: true, */
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=prefer TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	g.UseDB(db)

	g.ApplyBasic(model.Content{})
	g.ApplyBasic(model.Autoretrieve{})
	g.ApplyBasic(model.ContentDeal{})
	g.ApplyBasic(model.ObjRef{})
	g.ApplyBasic(model.Object{})
	g.ApplyBasic(model.AuthToken{})
	g.ApplyBasic(model.CollectionRef{})
	g.ApplyBasic(model.Collection{})
	g.ApplyBasic(model.DfeRecord{})
	g.ApplyBasic(model.User{})
	g.ApplyBasic(model.InviteCode{})
	g.ApplyBasic(model.ContentDeal{})
	g.ApplyBasic(model.Dealer{})
	g.ApplyBasic(model.Shuttle{})
	g.ApplyBasic(model.StorageMiner{})
	g.ApplyBasic(model.MinerStorageAsk{})
	g.ApplyBasic(model.RetrievalSuccessRecord{})
	g.ApplyBasic(model.RetrievalFailureRecord{})
	g.ApplyBasic(model.ProposalRecord{})
	g.ApplyBasic(model.PieceCommRecord{})
	g.ApplyBasic(model.PublishedBatch{})

	// apply basic crud rest on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generated table models' code when calling Execute.
	g.GenerateModel("contents")
	g.GenerateModel("content_deals")
	g.GenerateModel("auth_tokens")
	g.GenerateModel("autoretrieves")
	g.GenerateModel("dealers")
	g.GenerateModel("collection_refs")
	g.GenerateModel("collections")
	g.GenerateModel("content_deals")
	g.GenerateModel("contents")
	g.GenerateModel("dfe_records")
	g.GenerateModel("invite_codes")
	g.GenerateModel("miner_storage_asks")
	g.GenerateModel("obj_refs")
	g.GenerateModel("objects")
	g.GenerateModel("piece_comm_records")
	g.GenerateModel("proposal_records")
	g.GenerateModel("retrieval_failure_records")
	g.GenerateModel("retrieval_success_records")
	g.GenerateModel("shuttles")
	g.GenerateModel("storage_miners")
	g.GenerateModel("users")
	g.GenerateModel("published_batches")

	// execute the action of code generation
	g.Execute()
}
