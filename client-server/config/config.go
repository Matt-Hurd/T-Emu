package config

import (
	"client-server/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Chicago",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = database
}

func MigrateDatabase() {

	tx := DB.Exec(`do $$ declare
    r record;
begin
    for r in (select tablename from pg_tables where schemaname = 'public') loop
        execute 'drop table if exists ' || quote_ident(r.tablename) || ' cascade';
    end loop;
end $$;`,
	)

	if tx.Error != nil {
		log.Fatalf("Failed to drop tables: %v", tx.Error)
	}
	err := DB.AutoMigrate(
		&models.Account{},
		&models.Profile{},
		&models.ProfileInfo{},
		&models.ProfileInfoSettings{},
		&models.ProfileCustomization{},
		&models.ProfileHealth{},
		&models.ProfileHealthAttribute{},
		&models.ProfileBodyPart{},
		&models.ProfileInventory{},
		&models.ProfileAchievement{},
		&models.ProfileQuest{},
		&models.ProfileWishListItem{},
		&models.ProfileEncyclopediaEntry{},
		&models.ProfileItem{},
		&models.ProfileHideout{},
		&models.ProfileHideoutArea{},
		&models.ProfileBonus{},
		&models.ProfileNotes{},
		&models.ProfileRagfairInfo{},
		&models.ProfileTraderInfo{},
		&models.ProfileUnlockedInfo{},
		&models.ProfileSkillsGroup{},
		&models.ProfileSkill{},
		&models.ProfileStats{},
		&models.ProfileEftStats{},
		&models.ProfileCounter{},
		&models.ProfileCounterItem{},
		&models.ProfileDamageHistory{},
		&models.ProfileBodyPartDamage{},
		&models.ProfileStatus{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	tx2 := DB.Exec("INSERT INTO accounts (id, name, email) VALUES ('1', 'abahbob', 'test@gmail.com');")
	if tx2.Error != nil {
		log.Fatalf("Failed to drop tables: %v", tx2.Error)
	}

	createSession("1")
}
