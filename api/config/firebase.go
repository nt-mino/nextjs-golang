package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func init() {
    // Firebase Admin SDKを初期化
    ctx := context.Background()
    
    saPath := "./firebase-adminsdk.json"
    opt := option.WithCredentialsFile(saPath)
    
    // Firebaseアプリケーションの初期化
    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
        log.Fatalf("Firebaseの初期化に失敗しました: %v", err)
    }
    FirebaseApp = app
}


func GetFirestoreClient(ctx context.Context) (*firestore.Client, error) {
    return FirebaseApp.Firestore(ctx)
}
