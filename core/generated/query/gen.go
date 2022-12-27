package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

var (
	Q                      = new(Query)
	AuthToken              *authToken
	Autoretrieve           *autoretrieve
	Collection             *collection
	CollectionRef          *collectionRef
	Content                *content
	ContentDeal            *contentDeal
	Dealer                 *dealer
	DfeRecord              *dfeRecord
	InviteCode             *inviteCode
	MinerStorageAsk        *minerStorageAsk
	ObjRef                 *objRef
	Object                 *object
	PieceCommRecord        *pieceCommRecord
	ProposalRecord         *proposalRecord
	RetrievalFailureRecord *retrievalFailureRecord
	RetrievalSuccessRecord *retrievalSuccessRecord
	Shuttle                *shuttle
	StorageMiner           *storageMiner
	User                   *user
	PublishedBatch         *publishedBatch
)

func SetDefault(db *gorm.DB) {
	*Q = *Use(db)
	AuthToken = &Q.AuthToken
	Autoretrieve = &Q.Autoretrieve
	Collection = &Q.Collection
	CollectionRef = &Q.CollectionRef
	Content = &Q.Content
	ContentDeal = &Q.ContentDeal
	Dealer = &Q.Dealer
	DfeRecord = &Q.DfeRecord
	InviteCode = &Q.InviteCode
	MinerStorageAsk = &Q.MinerStorageAsk
	ObjRef = &Q.ObjRef
	Object = &Q.Object
	PieceCommRecord = &Q.PieceCommRecord
	ProposalRecord = &Q.ProposalRecord
	RetrievalFailureRecord = &Q.RetrievalFailureRecord
	RetrievalSuccessRecord = &Q.RetrievalSuccessRecord
	Shuttle = &Q.Shuttle
	StorageMiner = &Q.StorageMiner
	User = &Q.User
	PublishedBatch = &Q.PublishedBatch
}

func Use(db *gorm.DB) *Query {
	return &Query{
		db:                     db,
		AuthToken:              newAuthToken(db),
		Autoretrieve:           newAutoretrieve(db),
		Collection:             newCollection(db),
		CollectionRef:          newCollectionRef(db),
		Content:                newContent(db),
		ContentDeal:            newContentDeal(db),
		Dealer:                 newDealer(db),
		DfeRecord:              newDfeRecord(db),
		InviteCode:             newInviteCode(db),
		MinerStorageAsk:        newMinerStorageAsk(db),
		ObjRef:                 newObjRef(db),
		Object:                 newObject(db),
		PieceCommRecord:        newPieceCommRecord(db),
		ProposalRecord:         newProposalRecord(db),
		RetrievalFailureRecord: newRetrievalFailureRecord(db),
		RetrievalSuccessRecord: newRetrievalSuccessRecord(db),
		Shuttle:                newShuttle(db),
		StorageMiner:           newStorageMiner(db),
		User:                   newUser(db),
	}
}

type Query struct {
	db *gorm.DB

	AuthToken              authToken
	Autoretrieve           autoretrieve
	Collection             collection
	CollectionRef          collectionRef
	Content                content
	ContentDeal            contentDeal
	Dealer                 dealer
	DfeRecord              dfeRecord
	InviteCode             inviteCode
	MinerStorageAsk        minerStorageAsk
	ObjRef                 objRef
	Object                 object
	PieceCommRecord        pieceCommRecord
	ProposalRecord         proposalRecord
	RetrievalFailureRecord retrievalFailureRecord
	RetrievalSuccessRecord retrievalSuccessRecord
	Shuttle                shuttle
	StorageMiner           storageMiner
	User                   user
	PublishedBatch         publishedBatch
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                     db,
		AuthToken:              q.AuthToken.clone(db),
		Autoretrieve:           q.Autoretrieve.clone(db),
		Collection:             q.Collection.clone(db),
		CollectionRef:          q.CollectionRef.clone(db),
		Content:                q.Content.clone(db),
		ContentDeal:            q.ContentDeal.clone(db),
		Dealer:                 q.Dealer.clone(db),
		DfeRecord:              q.DfeRecord.clone(db),
		InviteCode:             q.InviteCode.clone(db),
		MinerStorageAsk:        q.MinerStorageAsk.clone(db),
		ObjRef:                 q.ObjRef.clone(db),
		Object:                 q.Object.clone(db),
		PieceCommRecord:        q.PieceCommRecord.clone(db),
		ProposalRecord:         q.ProposalRecord.clone(db),
		RetrievalFailureRecord: q.RetrievalFailureRecord.clone(db),
		RetrievalSuccessRecord: q.RetrievalSuccessRecord.clone(db),
		Shuttle:                q.Shuttle.clone(db),
		StorageMiner:           q.StorageMiner.clone(db),
		User:                   q.User.clone(db),
		PublishedBatch:         q.PublishedBatch.clone(db),
	}
}

type queryCtx struct {
	AuthToken              IAuthTokenDo
	Autoretrieve           IAutoretrieveDo
	Collection             ICollectionDo
	CollectionRef          ICollectionRefDo
	Content                IContentDo
	ContentDeal            IContentDealDo
	Dealer                 IDealerDo
	DfeRecord              IDfeRecordDo
	InviteCode             IInviteCodeDo
	MinerStorageAsk        IMinerStorageAskDo
	ObjRef                 IObjRefDo
	Object                 IObjectDo
	PieceCommRecord        IPieceCommRecordDo
	ProposalRecord         IProposalRecordDo
	RetrievalFailureRecord IRetrievalFailureRecordDo
	RetrievalSuccessRecord IRetrievalSuccessRecordDo
	Shuttle                IShuttleDo
	StorageMiner           IStorageMinerDo
	User                   IUserDo
	PublishedBatch         IPublishedBatchDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AuthToken:              q.AuthToken.WithContext(ctx),
		Autoretrieve:           q.Autoretrieve.WithContext(ctx),
		Collection:             q.Collection.WithContext(ctx),
		CollectionRef:          q.CollectionRef.WithContext(ctx),
		Content:                q.Content.WithContext(ctx),
		ContentDeal:            q.ContentDeal.WithContext(ctx),
		Dealer:                 q.Dealer.WithContext(ctx),
		DfeRecord:              q.DfeRecord.WithContext(ctx),
		InviteCode:             q.InviteCode.WithContext(ctx),
		MinerStorageAsk:        q.MinerStorageAsk.WithContext(ctx),
		ObjRef:                 q.ObjRef.WithContext(ctx),
		Object:                 q.Object.WithContext(ctx),
		PieceCommRecord:        q.PieceCommRecord.WithContext(ctx),
		ProposalRecord:         q.ProposalRecord.WithContext(ctx),
		RetrievalFailureRecord: q.RetrievalFailureRecord.WithContext(ctx),
		RetrievalSuccessRecord: q.RetrievalSuccessRecord.WithContext(ctx),
		Shuttle:                q.Shuttle.WithContext(ctx),
		StorageMiner:           q.StorageMiner.WithContext(ctx),
		User:                   q.User.WithContext(ctx),
		PublishedBatch:         q.PublishedBatch.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
