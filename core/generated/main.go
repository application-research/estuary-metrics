package main

import (
	"fmt"
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

	dcsEstuary, ok := viper.Get("DCS_ESTUARY").(string)
	if !ok {
		log.Fatalf("Error while reading Estuary database config")
	}

	dcsMetrics, ok := viper.Get("DCS_METRICS").(string)
	if !ok {
		log.Fatalf("Error while reading Metrics database config")
	}

	setupMainDb(dcsEstuary)
	setupMetricsDb(dcsMetrics)
}

func setupMetricsDb(dsn string) {

	fmt.Println(dsn)
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "generated/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	g.UseDB(db)

	// @jcace
	// For first run - comment out the "ApplyBasic" lines, and only run GenerateModel
	// Then subsequent migrations, leave ApplyBasic commented in
	g.ApplyBasic(model.RetrievalEvent{})
	g.GenerateModel("retrieval_events")

	g.Execute()
}

func setupMainDb(dsn string) {
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
	// ? Need to manually specify Autoretrieves otherwise it will use "autoretriefes" model name
	g.GenerateModelAs("autoretrieves", "Autoretrieve")
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
