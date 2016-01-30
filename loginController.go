package main

import (
	"crypto/sha512"
	"fmt"
	"git.gate.sh/mdeheij/monitoring/configuration"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mdeheij/gotools"
	"strings"
)

func getLoginUsername(c *gin.Context) string {
	session := sessions.Default(c)
	username := session.Get("username")
	if username != nil {
		return strings.ToLower(username.(string))
	}
	return ""
}

//Authenticate User
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !gotools.StringInSlice(getLoginUsername(c), configuration.Config.AllowedUsers) {
			c.Redirect(302, "/login?pleaseloginfirst")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func loginInit(r *gin.Engine) {
	/*	authTokens := make(map[string]string)
		authTokens["henk"] = "Twaalf"
	*/
	store := sessions.NewCookieStore([]byte(configuration.Config.SecureCookie))
	store.Options(configuration.Config.CookieConfig)
	r.Use(sessions.Sessions("GateSHSession", store))

	//group := r.Group("/login")
	group := r.Group("/login")
	{
		group.GET("/", loginPage)
		group.GET("/logout", logout)
		group.GET("/session", getSession)
		group.POST("/", func(c *gin.Context) {
			username := strings.ToLower(c.PostForm("username"))
			password := c.PostForm("password")

			passwordFromDatabase, err := dbmap.SelectStr("SELECT password FROM user WHERE username = ?", username)
			if err != nil {
				fmt.Println(err)
			}

			//hasher := sha1.New()
			//hasher.Write(password)
			//sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
			encryptedPass := Encrypt(password)
			if encryptedPass == passwordFromDatabase && password != "" {
				fmt.Println("YES! " + encryptedPass + " == " + passwordFromDatabase)
				session := sessions.Default(c)
				fmt.Println("Setting username as " + username)
				session.Set("username", username)
				session.Save()
				fmt.Println("Getting username for test: " + getLoginUsername(c))
				c.Redirect(302, "/")
			} else {
				fmt.Println("No! " + encryptedPass + " == " + passwordFromDatabase)
				c.Redirect(302, "/login?invalidpassword")
			}

		})

	}
}

func Encrypt(password string) string {
	h := sha512.New()
	h.Write(append([]byte(password)))
	return fmt.Sprintf("%x", h.Sum([]byte{}))
}

func loginPage(c *gin.Context) {
	fmt.Println(getLoginUsername(c))
	c.HTML(200, "login.tmpl", gin.H{
		"user":     getLoginUsername(c),
		"subtitle": "Easy server statistics monitoring",
	})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("username", "")
	session.Save()
	session.Clear()

	c.Redirect(302, "/login?logout-succesful")
}
func getSession(c *gin.Context) {
	c.JSON(200, gin.H{
		"username": getLoginUsername(c),
	})
}
