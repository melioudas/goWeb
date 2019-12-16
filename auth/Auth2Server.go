package auth

import (
	"dataPlatform/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	oredis "gopkg.in/go-oauth2/redis.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"log"
	"net/http"
)

var srvw *server.Server

func InitAuth(engine *gin.Engine) {
	fmt.Println("====>INIT AUTH! ")
	//load conf
	oci := model.ConfigParam["oauth-client-id"]
	ocs := model.ConfigParam["oauth-client-secret"]
	ocd := model.ConfigParam["oauth-client-domain"]
	or := model.ConfigParam["oauth-redis"]

	manager := manage.NewDefaultManager()
	// token memory store
	//manager.MustTokenStorage(store.NewMemoryTokenStore())
	// use redis token store
	manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr: or,
		DB: 0,
	}))
	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set(oci, &models.Client{
		ID:     oci,
		Secret: ocs,
		Domain: ocd,
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	//http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
	//	err := srv.HandleAuthorizeRequest(w, r)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusBadRequest)
	//	}
	//})
	//
	//http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
	//	srv.HandleTokenRequest(w, r)
	//})
	//
	//log.Fatal(http.ListenAndServe(":5555", nil))
	srvw = srv
	//authorize方法重写
	//engine.Any("/authorize1", func(context *gin.Context) {
	//	err := srv.HandleAuthorizeRequest(context.Writer, context.Request)
	//	if err != nil {
	//		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
	//	}
	//})

	//TOKEN方法重写
	engine.Any("/oauth/token", func(context *gin.Context) {
		srv.HandleTokenRequest(context.Writer, context.Request)
	})
	//engine.Any("/authorize", func(context *gin.Context) {
	//	w:=context.Writer
	//	r:=context.Request
	//	token, err := srv.ValidationBearerToken(r)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusBadRequest)
	//		return
	//	}
	//
	//	data := map[string]interface{}{
	//		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
	//		"client_id":  token.GetClientID(),
	//		"user_id":    token.GetUserID(),
	//	}
	//	e := json.NewEncoder(w)
	//	e.SetIndent("", "  ")
	//	e.Encode(data)
	//})

}


func Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		w:=context.Writer
		r:=context.Request
		_, err := srvw.ValidationBearerToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			context.Abort()
			fmt.Println("Abort===============")
			return
		}else {
			context.Next()
			fmt.Println("Next===============")
		}
		//data := map[string]interface{}{
		//	"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		//	"client_id":  token.GetClientID(),
		//	"user_id":    token.GetUserID(),
		//}
		//e := json.NewEncoder(w)
		//e.SetIndent("", "  ")
		//e.Encode(data)

	}
}

