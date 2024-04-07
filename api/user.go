package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"next-turtrial/api/config"
	"time"

	"github.com/julienschmidt/httprouter"
)

// ユーザー情報を取得する
func getUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Get User\n")

	// デモ用のユーザーデータ
	userData := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "Taro Go2",
	}

	// JSONとしてレスポンスを設定
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(userData); err != nil {
		fmt.Printf("エラー: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// ユーザーの作成
func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()
	client, err := config.GetFirestoreClient(ctx)
	if err != nil {
		log.Printf("Failed to get Firestore client: %s", err)
		return
	}

	_, _, err = client.Collection("users").Add(ctx, NewUser{
		Name:      "Taro Go",
		Email:     "taro.go@gmail.com",
		Plan:      "premium",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	fmt.Printf("user created\n")
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
}

type NewUser struct {
	Name      string    `firestore:"name"`
	Email     string    `firestore:"email"`
	Plan      string    `firestore:"plan"`
	CreatedAt time.Time `firestore:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at"`
}
