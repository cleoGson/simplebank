package api
 // this server will serve all http request for banking service
 type Server struct {
 store *db.store
 router *gin.Engine

 }