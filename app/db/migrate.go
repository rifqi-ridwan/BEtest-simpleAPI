package main

import (
	"BEtest-simpleAPI/app/db/migrations"
	"BEtest-simpleAPI/app/db/mysql"
	"BEtest-simpleAPI/domain"
	"context"
	"log"
	"os"

	_ "BEtest-simpleAPI/app/lib/loadenv"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/migrator"
)

func main() {
	adapter := mysql.OpenMysqlConnection()
	defer adapter.Close()

	repo := rel.New(adapter)
	m := migrator.New(repo)
	ctx := context.Background()

	m.Register(20211007172000, migrations.MigrateCreateCategories, migrations.RollbackCreateCategories)
	m.Register(20211007172100, migrations.MigrateCreateProducts, migrations.RollbackCreateProducts)

	var op string
	if len(os.Args) > 1 {
		op = os.Args[1]
	}

	switch op {
	case "migrate", "up":
		m.Migrate(ctx)
	case "rollback", "down":
		m.Rollback(ctx)
	default:
		log.Fatal("Command not recognized")
	}

	if op == "migrate" || op == "up" {
		categories, _ := fetchCategories(repo)
		if len(categories) == 0 {
			addCategory(repo, "Category 1")
			addCategory(repo, "Category 2")
		}
	}
}

func addCategory(repo rel.Repository, name string) error {
	var category domain.Category
	category.Name = name
	err := repo.Insert(context.Background(), &category)
	return err
}

func fetchCategories(repo rel.Repository) ([]domain.Category, error) {
	var categories []domain.Category
	err := repo.FindAll(context.Background(), &categories, rel.Unscoped(true))
	return categories, err
}
