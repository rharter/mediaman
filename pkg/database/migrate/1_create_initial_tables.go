package migrate

type rev1 struct{}

var CreateInitialTables = &rev1{}

func (r *rev1) Revision() int64 {
	return 1
}

func (r *rev1) Up(mg *MigrationDriver) error {
	t := mg.T
	if _, err := mg.CreateTable("movies", []string{
		t.Integer("id", PRIMARYKEY, AUTOINCREMENT),
		t.String("title"),
		t.String("backdrop"),
		t.String("poster"),
		t.Timestamp("release_date"),
		t.Bool("adult"),
		t.String("genres"),
		t.String("homepage"),
		t.String("imdb_id"),
		t.String("overview"),
		t.Integer("runtime"),
		t.String("tagline"),
		t.Real("rating"),
		t.Timestamp("created"),
		t.Timestamp("updated"),
		t.String("filename", UNIQUE),
	}); err != nil {
		return err
	}

	if _, err := mg.CreateTable("libraries", []string{
		t.Integer("id", PRIMARYKEY, AUTOINCREMENT),
		t.String("name"),
		t.String("path"),
		t.Timestamp("created"),
		t.Timestamp("updated"),
		t.Timestamp("last_scan"),
	}); err != nil {
		return err
	}

	if _, err := mg.CreateTable("series", []string{
		t.Integer("id", PRIMARYKEY, AUTOINCREMENT),
		t.String("language"),
		t.String("title"),
		t.String("overview"),
		t.String("banner"),
		t.String("imdb_id"),
		t.Integer("series_id"),
		t.Timestamp("created"),
		t.Timestamp("updated"),
	}); err != nil {
		return err
	}
	return nil
}

func (r *rev1) Down(mg *MigrationDriver) error {
	if _, err := mg.DropTable("movies"); err != nil {
		return err
	}
	if _, err := mg.DropTable("libraries"); err != nil {
		return err
	}
	return nil
}
