package controllers

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

// curl http://localhost:5000/login --include     --header "Content-Type: application/json" --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZnaXR5QHRoIiwiZXhwIjoxNjM1NTEzMTgwfQ.lqGApdcllILKQAL2D9qmmDkH5TceiBLUd_DtRLFSupw"    --request "POST" --data
// '{"email" : "fgity@th", "password" : "huy@34"}'

type usr struct {
	Email     string `json:"email"`
	UserName  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserName  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// type SetToken struct {
// 	Name     string
// 	Value    string
// 	MaxAge   int64
// 	Expires  time.Time
// 	HttpOnly bool
// }

// var t = &SetToken{}

var t = &http.Cookie{
	Value: " ",
}

var user_details = []User{}

var DB *sql.DB

var ACCESS_SECRET = "jdnfksdmfksd"

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var claims = Claims{
	Email: " ",
}

var tn time.Time

// var expirationTime time.Time

func InitDatabase() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "db1",
	}

	// db, err := sql.Open("mysql", "root:hiclass@12@/db1")

	db, err := sql.Open("mysql", cfg.FormatDSN())

	// Get a database handle.
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	DB = db
	// Close the database handle.
	// defer db.Close()

}

func Register(c *gin.Context) {
	var newuser User

	if err := c.BindJSON(&newuser); err != nil {
		return
	}
	// fmt.Println(newuser.Email, newuser.Password)

	hash := sha512.New()
	hash.Write([]byte(newuser.Password))
	pwd := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	_, err := DB.Exec("INSERT INTO user_db (Email, Password, UserName, FirstName, LastName) VALUES (?, ?, ?, ?, ?)", newuser.Email, pwd, newuser.UserName, newuser.FirstName, newuser.LastName)
	if err != nil {
		fmt.Printf("addUser: %v", err)
	}

	user_details = append(user_details, newuser)
	c.IndentedJSON(http.StatusCreated, "Successfully registered")
}

func ViewProfile(c *gin.Context) {

	claim := &Claims{}
	if t.Value != " " {
		token, err := jwt.ParseWithClaims(t.Value, claim, func(token *jwt.Token) (interface{}, error) {
			// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// 	return nil, fmt.Errorf("There was an error")
			// }
			return []byte(ACCESS_SECRET), nil
		})

		if err != nil {
			// c.JSON(http.StatusOK, gin.H{"Token": "Expired"})
			Logout(c)
			return
		}
		if _, ok := token.Claims.(*Claims); ok && token.Valid {
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"errorr": "Unauthorized"})
			return
		}
	}

	var user_details []User

	if claims.Email != " " {
		rows, err := DB.Query("SELECT * FROM user_db WHERE Email = ?", claims.Email)
		if err != nil {
			fmt.Printf("User_detailsbyname : %v", err)
		}
		defer rows.Close()

		user := User{}
		vp := usr{}

		for rows.Next() {

			if err := rows.Scan(&user.Email, &user.Password, &user.UserName, &user.FirstName, &user.LastName); err != nil {
				fmt.Printf("User_detailsbyname : %v", err)
			}
			user_details = append(user_details, user)
			vp = usr{
				Email:     user.Email,
				UserName:  user.UserName,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}
		}
		c.JSON(http.StatusOK, vp)
		// return
	}
	if claims.Email == " " {
		c.JSON(http.StatusNotFound, "User needs to be logged in")
		return
	}

	// timeNow := time.Now().Unix()
	// expT := time.After(10 * time.Second)

	// for {
	// 	select {
	// 	case <-expT:
	// 		Logout(c)
	// 		return
	// 	}
	// }

	// c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	// fmt.Println(user_details)

	// c.JSON(http.StatusOK, user_details)
}

func Login(c *gin.Context) {
	// user :[]User{}
	var user User

	if err := c.BindJSON(&user); err != nil {
		return
	}

	hash := sha512.New()
	hash.Write([]byte(user.Password))
	pwd := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	row := DB.QueryRow(`
		SELECT Email, Password FROM user_db WHERE
		Email = ? AND Password = ?`, user.Email, pwd)
	err := row.Scan(&user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, "Plz, enter correct details, User Not Found")
			return
		}
		c.JSON(http.StatusNotFound, "No such User details")
		return
	}

	expirationTime := time.Now().Add(15 * time.Second)

	claims = Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(ACCESS_SECRET))

	// c.SetCookie("gin_cookie", token, 60*60*24, "/", "localhost", false, true)

	// t = &SetToken{
	// 	Name:     "token",
	// 	Value:    token,
	// 	MaxAge:   60,
	// 	Expires:  expirationTime,
	// 	HttpOnly: true,
	// }
	t = &http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   60,
		Expires:  expirationTime,
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, t)

	// cookie, err := c.Request.Cookie(token)
	// fmt.Println(cookie, err)
	// if err == nil {
	// 	value := cookie.Value
	// 	fmt.Printf("Cookie value: %v", value)
	// }

	// fmt.Println(user.Email, user.Password, pwd)
	c.JSON(http.StatusOK, "Successfully logged in")

}

func Logout(c *gin.Context) {
	tn = time.Now()

	if claims.Email != " " {
		if t.Expires.Unix() < tn.Unix() {
			t = &http.Cookie{
				Name:     "token",
				Value:    " ",
				MaxAge:   -1,
				Expires:  time.Now().Add(-time.Minute),
				HttpOnly: true,
			}
			http.SetCookie(c.Writer, t)
			claims.Email = " "
			c.JSON(http.StatusOK, gin.H{"Token Expired": "Successfully logged out"})
		} else {
			t = &http.Cookie{
				Name:     "token",
				Value:    " ",
				MaxAge:   -1,
				Expires:  time.Now().Add(-time.Minute),
				HttpOnly: true,
			}
			http.SetCookie(c.Writer, t)
			claims.Email = " "
			c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
		}
		return
	}

	return
}
