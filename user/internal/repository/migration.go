package repository

func migration() {
	err := DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&User{},
		)

	if err != nil {
		panic(err)
	}
}
