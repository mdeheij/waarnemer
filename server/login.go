package server

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mdeheij/gotools"
	"github.com/mdeheij/monitoring/configuration"
	"golang.org/x/crypto/bcrypt"
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
	r.Use(sessions.Sessions(configuration.Config.SecureCookieName, store))

	//group := r.Group("/login")
	group := r.Group("/login")
	{
		group.GET("/", loginPage)
		group.GET("/logout", logout)
		group.GET("/session", getSession)
		group.POST("/", func(c *gin.Context) {
			username := strings.ToLower(c.PostForm("username"))
			password := c.PostForm("password")

			//hasher := sha1.New()
			//hasher.Write(password)
			//sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

			if CheckPassword(username, password) {
				session := sessions.Default(c)
				fmt.Println("Setting username " + username)
				session.Set("username", username)
				session.Save()
				fmt.Println("Creating session for " + getLoginUsername(c))
				c.Redirect(302, "/")
			} else {
				fmt.Println("Invalid credentials!")
				fmt.Println("If this is a new password, please set the following hash:", GenerateFromPassword(password))
				c.Redirect(302, "/login?invalidpassword")
			}

		})

	}
}

// func Encrypt(password string) string {
// 	h := sha512.New()
// 	h.Write(append([]byte(password)))
// 	return fmt.Sprintf("%x", h.Sum([]byte{}))
// }
func GenerateFromPassword(password string) string {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
func CheckPassword(username string, password string) bool {
	// Hashing the password with the default cost of 10

	hashedPassword, err := dbmap.SelectStr("SELECT password FROM user WHERE username = ?", username)
	if err != nil {
		fmt.Println(err)
	}
	errCompare := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if errCompare == nil {
		return true
	} else {
		return false
	}

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
