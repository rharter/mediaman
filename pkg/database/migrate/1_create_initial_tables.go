package migrate

type rev1 struct{}

var CreateInitialTables = &rev1{}

func (r *rev1) Revision() int64 {
	return 1
}

func (r *rev1) Up(mg *MigrationDriver) error {
	t := mg.T

	if _, err := mg.CreateTable("libraries", []string{
		t.Integer("id", PRIMARYKEY, AUTOINCREMENT),
		t.String("type"),
		t.String("name"),
		t.Integer("root_id"),
		t.Timestamp("created"),
		t.Timestamp("updated"),
		t.Timestamp("last_scan"),
	}); err != nil {
		return err
	}

	if _, err := mg.CreateTable("elements", []string{
		t.Integer("id", PRIMARYKEY, AUTOINCREMENT),
		t.String("file", UNIQUE),
		t.String("type"),
		t.Integer("parent_id"),
		t.String("title"),
		t.String("description"),
		t.String("thumbnail"),
		t.String("background"),
		t.String("poster"),
		t.String("banner"),
		t.String("remote_id"),
		t.Timestamp("created"),
		t.Timestamp("updated"),
	}); err != nil {
		return err
	}
	return nil
}

func (r *rev1) Down(mg *MigrationDriver) error {
	if _, err := mg.DropTable("elements"); err != nil {
		return err
	}
	if _, err := mg.DropTable("directories"); err != nil {
		return err
	}
	if _, err := mg.DropTable("libraries"); err != nil {
		return err
	}
	return nil
}
