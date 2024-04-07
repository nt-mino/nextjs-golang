package main

import (
	"context"
	"fmt"
	"log"
	"next-turtrial/api/config"
)



func main() {

    ctx := context.Background()
    
    client, err := config.GetFirestoreClient(ctx)
    if err != nil {
        log.Fatalf("Firestoreクライアントの取得に失敗しました: %v", err)
    }
    fmt.Println("Firestoreクライアントの取得に成功しました")

    // * firestore.Client型のclientを使ってFirestoreの操作を行う
    dsnap, err := client.Collection("users").Doc("YMmC8X8zowCf7Dt6WXGK").Get(ctx)
    if err != nil {
        log.Fatalf("Documentの取得に失敗しました: %v", err)
    }
    fmt.Println(dsnap.Data())


    defer client.Close()
      
    // fmt.Printf("hello, world\n")
}



