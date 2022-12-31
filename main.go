package main

import (
	"context"
	"demo_goent/ent"
	"demo_goent/ent/user"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// -- STEP7: Delete
func main() {
	entOptions := []ent.Option{}
	entOptions = append(entOptions, ent.Debug())
	// client, err := ent.Open("mysql", "root:myrootpassword@tcp(localhost:3306)/mydb?parseTime=True", entOptions...)
	client, err := ent.Open("mysql", "root:myrootpassword@tcp(localhost:3306)/mydb?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	usr, err := client.Debug().User.
		Delete().
		Where(user.Age(10)).
		Exec(ctx)
	if err != nil {
		log.Fatalf("failed create user: %v", err)
	}
	log.Printf("user: %+v", usr)
	log.Print("ent sample done")
}

// -- STEP6: Read
// func main() {
// 	entOptions := []ent.Option{}
// 	entOptions = append(entOptions, ent.Debug())
// 	// client, err := ent.Open("mysql", "root:myrootpassword@tcp(localhost:3306)/mydb?parseTime=True", entOptions...)
// 	client, err := ent.Open("mysql", "root:myrootpassword@tcp(localhost:3306)/mydb?parseTime=True")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to mysql: %v", err)
// 	}
// 	defer client.Close()

// 	ctx := context.Background()
// 	usr, err := client.Debug().User.
// 		Query().
// 		Where(user.Age(10)).
// 		All(ctx)
// 	if err != nil {
// 		log.Fatalf("failed create user: %v", err)
// 	}
// 	log.Printf("user: %+v", usr)
// 	log.Print("ent sample done")
// }

// -- STEP5: Update
// func main() {
// 	entOptions := []ent.Option{}
// 	entOptions = append(entOptions, ent.Debug())
// 	// client, err := ent.Open("mysql", "root:myrootpassword@tcp(localhost:3306)/mydb?parseTime=True", entOptions...)
// 	client, err := ent.Open("mysql", "root:myrootpassword@tcp(localhost:3306)/mydb?parseTime=True")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to mysql: %v", err)
// 	}
// 	defer client.Close()

// 	ctx := context.Background()
// 	cmp, err := client.Debug().Company.
// 		Update().
// 		SetName("companyB").
// 		Where(company.Name("companyA")).
// 		Save(ctx)
// 	if err != nil {
// 		log.Fatalf("failed create company: %v", err)
// 	}
// 	log.Printf("cmp: %+v", cmp)

// 	usr, err := client.Debug().User.
// 		Update().
// 		SetAge(10).
// 		Where(user.Age(20)).
// 		Save(ctx)
// 	if err != nil {
// 		log.Fatalf("failed create user: %v", err)
// 	}
// 	log.Printf("user: %+v", usr)
// 	log.Print("ent sample done")
// }

// -- STEP4: Create
// func main() {
// 	entOptions := []ent.Option{}
// 	entOptions = append(entOptions, ent.Debug())
// 	// client, err := ent.Open("mysql", "root:myrootpassword@tcp(localhost:3306)/mydb?parseTime=True", entOptions...)
// 	client, err := ent.Open("mysql", "root:myrootpassword@tcp(localhost:3306)/mydb?parseTime=True")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to mysql: %v", err)
// 	}
// 	defer client.Close()

// 	ctx := context.Background()
// 	cmp, err := client.Debug().Company.
// 		Create().
// 		SetName("companyA").
// 		Save(ctx)
// 	if err != nil {
// 		log.Fatalf("failed create company: %v", err)
// 	}
// 	log.Printf("cmp: %+v", cmp)

// 	usr, err := client.Debug().User.
// 		Create().
// 		SetFirstName("first name").
// 		SetLastName("last name").
// 		SetAge(20).
// 		SetEmail("example@example.co.jp").
// 		SetCompany(cmp).
// 		Save(ctx)
// 	if err != nil {
// 		log.Fatalf("failed create user: %v", err)
// 	}
// 	log.Printf("user: %+v", usr)
// 	log.Print("ent sample done")
// }

// -- STEP3: テーブル定義反映
// func main() {
// 	client, err := ent.Open("mysql", "root:myrootpassword@tcp(localhost:3306)/mydb?parseTime=True")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to mysql: %v", err)
// 	}
// 	defer client.Close()
// 	ctx := context.Background()
// 	if err := client.Schema.Create(ctx, migrate.WithForeignKeys(false)); err != nil {
// 		log.Fatalf("failed printing schema changes: %v", err)
// 	}
// 	log.Print("ent sample done")
// }

// -- STEP2: スキーマ定義の確認
// func main() {
// 	client, err := ent.Open("mysql", "root:myrootpassword@tcp(localhost:3306)/mydb?parseTime=True")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to mysql: %v", err)
// 	}
// 	defer client.Close()
// 	ctx := context.Background()
// 	// dump migration changes to stdout
// 	if err := client.Schema.WriteTo(ctx, os.Stdout, migrate.WithForeignKeys(false)); err != nil {
// 		log.Fatalf("failed creating schema resources: %v", err)
// 	}
// 	log.Print("ent sample done")
// }

// -- STEP1: DB接続確認
// func main() {
// 	client, err := ent.Open("mysql", "root:myrootpassword@tcp(localhost:3306)/mydb?parseTime=True")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to mysql: %v", err)
// 	}
// 	defer client.Close()
// 	// Run the auto migration tool.
// 	if err := client.Schema.Create(context.Background()); err != nil {
// 		log.Fatalf("failed creating schema resources: %v", err)
// 	}
// 	log.Print("ent sample done")
// }
