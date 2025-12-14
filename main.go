package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Simple in-memory user store (for demo purposes). Replace with a DB in prod.
type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	PasswordHash string `json:"-"`
	Provider     string `json:"provider"`
	ProviderID   string `json:"provider_id"`
}

type UserStore struct {
	mu    sync.Mutex
	users map[string]*User // key by email
}

func NewUserStore() *UserStore {
	return &UserStore{users: make(map[string]*User)}
}

func (s *UserStore) GetByEmail(email string) (*User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, ok := s.users[email]
	return u, ok
}

func (s *UserStore) Upsert(u *User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[u.Email] = u
}

func (s *UserStore) CreateLocal(email, name, password string) (*User, error) {
	if _, ok := s.GetByEmail(email); ok {
		return nil, errors.New("email already exists")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &User{ID: generateID(), Email: email, Name: name, PasswordHash: string(hash), Provider: "local"}
	s.Upsert(u)
	return u, nil
}

func generateID() string {
	b := make([]byte, 12)
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

var (
	userStore = NewUserStore()
	jwtKey    = []byte(getenv("APP_SECRET", "dev-secret"))
	oauthCfg  *oauth2.Config
)

func getenv(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}

func main() {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		log.Println("WARNING: GOOGLE_CLIENT_ID or GOOGLE_CLIENT_SECRET not set — Google OAuth will fail")
	}
	oauthCfg = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  getenv("GOOGLE_REDIRECT", "http://localhost:8080/auth/google/callback"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/signup", handleSignup)
		api.POST("/login", handleLogin)
		api.GET("/me", handleMe)
		api.POST("/logout", handleLogout)
	}

	r.GET("/auth/google/login", handleGoogleLogin)
	r.GET("/auth/google/callback", handleGoogleCallback)

	// Serve static assets for the built React app (ui/dist) and fallback to index.html for the SPA
	r.Static("/assets", "ui/dist/assets")
	r.StaticFile("/favicon.ico", "ui/dist/favicon.ico")
	r.GET("/", func(c *gin.Context) { c.File("ui/dist/index.html") })
	r.NoRoute(func(c *gin.Context) { c.File("ui/dist/index.html") })

	port := getenv("PORT", "8080")
	log.Printf("starting server on :%s", port)
	r.Run(":" + port)
}

// --- Handlers ---

type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

func handleSignup(c *gin.Context) {
	var req signupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := userStore.CreateLocal(req.Email, req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	setSessionCookie(c, u)
	c.JSON(http.StatusOK, gin.H{"user": u})
}

type loginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func handleLogin(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, ok := userStore.GetByEmail(req.Email)
	if !ok || u.PasswordHash == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	setSessionCookie(c, u)
	c.JSON(http.StatusOK, gin.H{"user": u})
}

func handleMe(c *gin.Context) {
	u, err := getUserFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": u})
}

func handleLogout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{Name: "session", Value: "", Path: "/", Expires: time.Unix(0, 0), HttpOnly: true})
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

// --- Session / JWT helpers ---

func setSessionCookie(c *gin.Context, u *User) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  u.Email,
		"name": u.Name,
		"exp":  time.Now().Add(7 * 24 * time.Hour).Unix(),
	})
	ss, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{Name: "session", Value: ss, Path: "/", HttpOnly: true})
}

func getUserFromRequest(c *gin.Context) (*User, error) {
	cookie, err := c.Request.Cookie("session")
	if err != nil {
		return nil, err
	}
	tkn, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil || !tkn.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := tkn.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	email, _ := claims["sub"].(string)
	if email == "" {
		return nil, errors.New("no sub")
	}
	u, ok := userStore.GetByEmail(email)
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

// --- Google OAuth2 ---

func handleGoogleLogin(c *gin.Context) {
	state := generateID()
	// store state in cookie
	http.SetCookie(c.Writer, &http.Cookie{Name: "oauthstate", Value: state, Path: "/", HttpOnly: true, Expires: time.Now().Add(10 * time.Minute)})
	url := oauthCfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

func handleGoogleCallback(c *gin.Context) {
	state := c.Query("state")
	cookie, err := c.Request.Cookie("oauthstate")
	if err != nil || cookie.Value != state {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid oauth state"})
		return
	}
	code := c.Query("code")
	tok, err := oauthCfg.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to exchange token"})
		return
	}
	// fetch user info
	client := oauthCfg.Client(context.Background(), tok)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed fetching userinfo"})
		return
	}
	defer resp.Body.Close()
	var info struct {
		Id            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		Picture       string `json:"picture"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed parsing userinfo"})
		return
	}
	// upsert user
	u, ok := userStore.GetByEmail(info.Email)
	if !ok {
		u = &User{ID: generateID(), Email: info.Email, Name: info.Name, Provider: "google", ProviderID: info.Id}
		userStore.Upsert(u)
	} else {
		u.Provider = "google"
		u.ProviderID = info.Id
		userStore.Upsert(u)
	}
	setSessionCookie(c, u)
	// Redirect into the app (dashboard)
	c.Redirect(http.StatusFound, "/dashboard")
}
