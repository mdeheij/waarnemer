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

// CsrfOptions stores the options to use for CSRF protection.
var CsrfOptions csrf.Options

//AuthRequired is authentication middleware for user authenticaton.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := getLoginUsername(c)

		_, err := getUserByUsername(username)

		if err == nil {
			c.Next()
		} else {
			c.Redirect(302, "/login?pleaseloginfirst")
			c.Abort()
		}
	}
}

// HashPassword hashes the provided password.
// The hashing uses bcrypt with a default cost of 10.
func HashPassword(password string) string {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

// CheckPassword compares whether or not the provided password matches the saved password hash for the provided username.
func CheckPassword(username string, password string) bool {
	user, err := getUserByUsername(username)

	if err == nil {
		compareErr := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))

		if err != nil {
			fmt.Println(compareErr)
			return false
		}

		return true
	}

	fmt.Println(err)
	return false
}

func getLoginUsername(c *gin.Context) string {
	session := sessions.Default(c)
	username := session.Get("username")

	if username != nil {
		return strings.ToLower(username.(string))
	}

	return ""
}

func loginInit(r *gin.Engine) {
	CsrfOptions = csrf.Options{
		Secret: configuration.Config.SecureCookie,
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "Please try again later")
			c.Abort()
		},
	}

	store := sessions.NewCookieStore([]byte(configuration.Config.SecureCookie))
	store.Options(configuration.Config.CookieConfig)

	r.Use(sessions.Sessions(configuration.Config.SecureCookieName, store))
	r.Use(csrf.Middleware(CsrfOptions))

	group := r.Group("/login")
	{
		group.GET("/", loginPage)
		group.GET("/logout", logout)
		group.GET("/session", getSession)
		group.POST("/", func(c *gin.Context) {
			username := strings.ToLower(c.PostForm("username"))
			password := c.PostForm("password")

			if CheckPassword(username, password) {
				session := sessions.Default(c)
				fmt.Println("Setting username " + username)
				session.Set("username", username)
				session.Save()
				fmt.Println("Creating session for " + getLoginUsername(c))
				c.Redirect(302, "/")
			} else {
				fmt.Println("Invalid credentials!")
				fmt.Println("If this is a new password, please set the following hash:", HashPassword(password))
				c.Redirect(302, "/login?invalidpassword")
			}
		})
	}
}

// getUserByUsername returns the user details if the username matches with the known usernames.
// An empty result and an error will be returned if the user is not found.
func getUserByUsername(username string) (configuration.User, error) {
	for _, v := range configuration.Config.Users {
		if v.Username == username {
			return v, nil
		}
	}

	return configuration.User{}, fmt.Errorf("Cannot find user %s", username)
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
