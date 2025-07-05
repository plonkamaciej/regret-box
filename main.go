package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var ctx = context.Background()
var rdb *redis.Client

type RegretData struct {
	ID      string
	Text    string
	Likes   int
	CanLike bool
}

type PopularRegret struct {
	ID    string
	Text  string
	Likes int
}

type LikeResponse struct {
	Success bool `json:"success"`
	Likes   int  `json:"likes"`
	Message string `json:"message"`
}

func main() {
	// Get Redis URL from environment, default to localhost
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // No password by default
		DB:       0,
	})

	// Test Redis connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")

	// Get port from environment, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/random", randomHandler)
	http.HandleFunc("/r/", viewHandler)
	http.HandleFunc("/like/", likeHandler)
	http.HandleFunc("/popular", popularHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Serving on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	text := r.FormValue("regret")
	if text == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	id := uuid.New().String()
	err := rdb.Set(ctx, "regret:"+id, text, 24*time.Hour).Err()
	if err != nil {
		http.Error(w, "Unable to save regret", 500)
		return
	}
	http.Redirect(w, r, "/r/"+id, http.StatusSeeOther)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/r/"):]
	text, err := rdb.Get(ctx, "regret:"+id).Result()
	if err == redis.Nil || err != nil {
		http.NotFound(w, r)
		return
	}

	likes, err := rdb.SCard(ctx, "likes:"+id).Result()
	if err != nil {
		http.Error(w, "Unable to retrieve likes", 500)
		return
	}

	canLike := true
	ip := r.RemoteAddr
	if strings.Contains(ip, ":") {
		ip, _, _ = net.SplitHostPort(ip)
	}
	liked, err := rdb.SIsMember(ctx, "likes:"+id, ip).Result()
	if err != nil {
		http.Error(w, "Unable to check like status", 500)
		return
	}
	if liked {
		canLike = false
	}

	regretData := RegretData{
		ID:      id,
		Text:    text,
		Likes:   int(likes),
		CanLike: canLike,
	}

	tmpl := template.Must(template.ParseFiles("templates/regret.html"))
	tmpl.Execute(w, regretData)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	keys, err := rdb.Keys(ctx, "regret:*").Result()
	if err != nil || len(keys) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	randomKey := keys[time.Now().UnixNano()%int64(len(keys))]
	id := randomKey[len("regret:"):]
	http.Redirect(w, r, "/r/"+id, http.StatusSeeOther)
}

func likeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/like/"):]
	ip := r.RemoteAddr
	if strings.Contains(ip, ":") {
		ip, _, _ = net.SplitHostPort(ip)
	}

	// Check if already liked
	liked, err := rdb.SIsMember(ctx, "likes:"+id, ip).Result()
	if err != nil {
		if r.Header.Get("Content-Type") == "application/json" || r.Header.Get("Accept") == "application/json" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(LikeResponse{Success: false, Message: "Error checking like status"})
			return
		}
		http.Error(w, "Unable to check like status", 500)
		return
	}

	if liked {
		if r.Header.Get("Content-Type") == "application/json" || r.Header.Get("Accept") == "application/json" {
			likes, _ := rdb.SCard(ctx, "likes:"+id).Result()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(LikeResponse{Success: false, Likes: int(likes), Message: "Already liked"})
			return
		}
		http.Redirect(w, r, "/r/"+id, http.StatusSeeOther)
		return
	}

	// Add the like
	err = rdb.SAdd(ctx, "likes:"+id, ip).Err()
	if err != nil {
		if r.Header.Get("Content-Type") == "application/json" || r.Header.Get("Accept") == "application/json" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(LikeResponse{Success: false, Message: "Unable to like regret"})
			return
		}
		http.Error(w, "Unable to like regret", 500)
		return
	}

	// Get updated like count
	likes, err := rdb.SCard(ctx, "likes:"+id).Result()
	if err != nil {
		likes = 1 // fallback
	}

	// Return JSON for AJAX requests, redirect for regular requests
	if r.Header.Get("Content-Type") == "application/json" || r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(LikeResponse{Success: true, Likes: int(likes), Message: "Liked successfully"})
	} else {
		http.Redirect(w, r, "/r/"+id, http.StatusSeeOther)
	}
}

func popularHandler(w http.ResponseWriter, r *http.Request) {
	keys, err := rdb.Keys(ctx, "regret:*").Result()
	if err != nil {
		http.Error(w, "Unable to retrieve regrets", 500)
		return
	}

	var popularRegrets []PopularRegret
	for _, key := range keys {
		id := key[len("regret:"):]
		likes, err := rdb.SCard(ctx, "likes:"+id).Result()
		if err != nil {
			continue
		}
		text, err := rdb.Get(ctx, key).Result()
		if err != nil {
			continue
		}
		popularRegrets = append(popularRegrets, PopularRegret{
			ID:    id,
			Text:  text,
			Likes: int(likes),
		})
	}

	sort.Slice(popularRegrets, func(i, j int) bool {
		return popularRegrets[i].Likes > popularRegrets[j].Likes
	})

	// Take only top 5
	if len(popularRegrets) > 5 {
		popularRegrets = popularRegrets[:5]
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}

	tmpl := template.Must(template.New("popular.html").Funcs(funcMap).ParseFiles("templates/popular.html"))
	tmpl.Execute(w, popularRegrets)
}
