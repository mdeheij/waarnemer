package server

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mdeheij/gin-csrf"
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

//AuthRequired is authentication middleware for user
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := getLoginUsername(c)

		if len(getUserByUsername(username).Username) > 0 { //TODO fix this
			//c.Next()
			c.Next()
		} else {
			c.Redirect(302, "/login?pleaseloginfirst")
			c.Abort()
		}
	}
}

var CsrfOpties csrf.Options

func loginInit(r *gin.Engine) {

	CsrfOpties = csrf.Options{
		Secret: configuration.Config.SecureCookie,
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "Please try again later")
			c.Abort()
		},
	}

	store := sessions.NewCookieStore([]byte(configuration.Config.SecureCookie))
	store.Options(configuration.Config.CookieConfig)
	r.Use(sessions.Sessions(configuration.Config.SecureCookieName, store))

	r.Use(csrf.Middleware(CsrfOpties))

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

func GenerateFromPassword(password string) string {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func getUserByUsername(username string) configuration.User {
	for _, v := range configuration.Config.Users {
		if v.Username == username {
			return v
		}
	}
	return configuration.User{}
}

//CheckPassword compares input to saved hash of password
func CheckPassword(username string, password string) bool {

	hashedPassword := getUserByUsername(username).Hash
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//TODO implement error handling

	errCompare := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if errCompare == nil {
		return true
	}

	return false
}

func loginPage(c *gin.Context) {
	fmt.Println(getLoginUsername(c))
	c.HTML(200, "login.tmpl", gin.H{
		"user":     getLoginUsername(c),
		"csrf":     csrf.GetToken(c),
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
