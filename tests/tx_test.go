package tests_test

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// func setupDB(cfg config.Configuration) (DBI db.DB, cleanup func(), err error) {
// 	tmpfile, err := ioutil.TempFile("", "test.*.db")
// 	if err != nil {
// 		return
// 	}

// 	sqliteDb, err := gorm.Open("sqlite3", tmpfile.Name())
// 	if err != nil {
// 		return
// 	}
// 	logger := log.Logger.With("module", "tests")

// 	DBI = db.DB{Instance: sqliteDb, Logger: logger, Cfg: cfg}

// 	allMigrations := migrations.GetMigrations()
// 	m := migrations.NewGormigrate(DBI.Instance, migrations.DefaultOptions, allMigrations)
// 	err = m.Migrate()
// 	if err != nil {
// 		return
// 	}
// 	cleanup = func() {
// 		DBI.Close()
// 		os.Remove(tmpfile.Name())
// 	}

// 	return DBI, cleanup, nil
// }
