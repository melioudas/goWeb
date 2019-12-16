package auth
//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"io"
//	"net/http"
//	"time"
//
//	"golang.org/x/oauth2"
//	"golang.org/x/oauth2/clientcredentials"
//)
//
//const (
//	authServerURL = "http://localhost:5555"
//)
//
//var (
//	config = oauth2.Config{
//		ClientID:     "000000",
//		ClientSecret: "999999",
//		Scopes:       []string{"all"},
//		RedirectURL:  "http://localhost:5555/oauth2",
//		Endpoint: oauth2.Endpoint{
//			AuthURL:  authServerURL + "/authorize",
//			TokenURL: authServerURL + "/token",
//		},
//	}
//	globalToken *oauth2.Token // Non-concurrent security
//)
//
//
//func OAuth2(engine *gin.Engine) {
//
//	engine.Any(  "/", func(c *gin.Context) {
//		u := config.AuthCodeURL("xyz")
//		http.Redirect(c.Writer, c.Request, u, http.StatusFound)
//	})
//
//
//	engine.Any( "/oauth2", func(c *gin.Context) {
//		w := c.Writer
//		r := c.Request
//
//		r.ParseForm()
//		state := r.Form.Get("state")
//		if state != "xyz" {
//			http.Error(w, "State invalid", http.StatusBadRequest)
//			return
//		}
//		code := r.Form.Get("code")
//		if code == "" {
//			http.Error(w, "Code not found", http.StatusBadRequest)
//			return
//		}
//		token, err := config.Exchange(context.Background(), code)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		globalToken = token
//
//		e := json.NewEncoder(w)
//		e.SetIndent("", "  ")
//		e.Encode(token)
//	})
//
//	engine.Any(  "/refresh", func(i *gin.Context) {
//		w := i.Writer
//		r := i.Request
//
//		if globalToken == nil {
//			http.Redirect(w, r, "/", http.StatusFound)
//			return
//		}
//
//		globalToken.Expiry = time.Now()
//		token, err := config.TokenSource(context.Background(), globalToken).Token()
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		globalToken = token
//		e := json.NewEncoder(w)
//		e.SetIndent("", "  ")
//		e.Encode(token)
//	})
//
//	engine.Any(  "/try", func(i *gin.Context) {
//		w := i.Writer
//		r := i.Request
//		if globalToken == nil {
//			http.Redirect(w, r, "/", http.StatusFound)
//			return
//		}
//
//		resp, err := http.Get(fmt.Sprintf("%s/test?access_token=%s", authServerURL, globalToken.AccessToken))
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		defer resp.Body.Close()
//
//		io.Copy(w, resp.Body)
//	})
//
//	engine.Any(  "/pwd", func(i *gin.Context) {
//		w := i.Writer
//		token, err := config.PasswordCredentialsToken(context.Background(), "test", "test")
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		globalToken = token
//		e := json.NewEncoder(w)
//		e.SetIndent("", "  ")
//		e.Encode(token)
//	})
//
//	engine.Any(  "/client", func(i *gin.Context) {
//		w := i.Writer
//		cfg := clientcredentials.Config{
//			ClientID:     config.ClientID,
//			ClientSecret: config.ClientSecret,
//			TokenURL:     config.Endpoint.TokenURL,
//		}
//
//		token, err := cfg.Token(context.Background())
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		e := json.NewEncoder(w)
//		e.SetIndent("", "  ")
//		e.Encode(token)
//	})
//}
//
