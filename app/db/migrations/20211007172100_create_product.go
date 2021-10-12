package migrations

import "github.com/go-rel/rel"

func MigrateCreateProducts(schema *rel.Schema) {
	schema.CreateTable("products", func(t *rel.Table) {
		t.ID("id")
		t.String("name")
		t.Int("price")
		t.Int("category_id", rel.Unsigned(true), rel.Required(true))
		t.String("image")
		t.DateTime("created_at")
		t.DateTime("updated_at")

		t.ForeignKey("category_id", "categories", "id")
	})

	schema.CreateIndex("products", "name", []string{"name"})
	schema.CreateIndex("products", "category_id", []string{"category_id"})
}

func RollbackCreateProducts(schema *rel.Schema) {
	schema.DropTable("products")
}
