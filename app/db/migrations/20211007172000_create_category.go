package migrations

import "github.com/go-rel/rel"

func MigrateCreateCategories(schema *rel.Schema) {
	schema.CreateTable("categories", func(t *rel.Table) {
		t.ID("id")
		t.String("name")
		t.DateTime("created_at")
		t.DateTime("updated_at")
	})

	schema.CreateIndex("categories", "name", []string{"name"})
}

func RollbackCreateCategories(schema *rel.Schema) {
	schema.DropTable("categories")
}
