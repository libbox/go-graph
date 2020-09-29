package main

import (
	"context"
	"fmt"
	"graph/ent/post"
	"log"

	"graph/ent"

	_ "github.com/go-sql-driver/mysql"
	"graph/ent/user"
)

const (
	pass = "Tuan_10021997"
	user1 = "kelvin"
)

type UserRepo interface {
	Create(name string, age int) (err error)
	GetByName(name string) (res *ent.User, err error)
}

func main() {
	client, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/graph?parseTime=True", user1, pass))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	ctx := context.Background()
	_, err = client.User.Create().SetName("tuan").Save(ctx)
	_, err = client.User.Create().SetName("toan").Save(ctx)
	client.User.Create().SetName("hieu").Save(ctx)
	client.User.Create().SetName("triet").Save(ctx)

	tuan, _ := client.User.Query().Where(user.Name("tuan")).First(ctx)
	toan, _ := client.User.Query().Where(user.Name("toan")).First(ctx)
	hieu, _ := client.User.Query().Where(user.Name("hieu")).First(ctx)
	triet, _ := client.User.Query().Where(user.Name("triet")).First(ctx)
	tuan.Update().AddFollowers(toan).Save(ctx)
	hieu.Update().AddFollowing(tuan).Save(ctx)
	toan.Update().AddFriend(hieu).Save(ctx)
	tuan.Update().AddFriend(hieu).Save(ctx)
	hieu.Update().AddFriend(triet).Save(ctx)
	fmt.Println(tuan.QueryFollowers().All(ctx))
	fmt.Println(hieu.QueryFollowing().All(ctx))
	fmt.Println(toan.QueryFriend().All(ctx))
	fmt.Println(hieu.QueryFriend().All(ctx))
	fmt.Println(tuan.QueryFriend().All(ctx))
	//sg, _ := tuan.QueryFriend().Where(user.HasFriendWith(user.ID(hieu.ID))).All(ctx)
	sg, _ := client.User.Query().
		Where(user.HasFriendWith(user.HasFriendWith(user.ID(toan.ID)))).All(ctx)

	fmt.Println(sg)


	client.Post.Create().SetContent("post 1").Save(ctx)

	post1, _ := client.Post.Query().Where(post.Content("post 1")).First(ctx)
	tuan.Update().AddLikes(post1)
	post1.Update().AddLikes(hieu)
	//post1.Update()
	fmt.Println(tuan.QueryLikes().All(ctx))
	fmt.Println(post1.QueryLikes().All(ctx))
	defer client.Close()
}
