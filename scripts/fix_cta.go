package main

import (
	"log"
	"mla-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgresql://neondb_owner:npg_OHF4oC9iqPJE@ep-empty-brook-ak7sp9mg-pooler.c-3.us-west-2.aws.neon.tech/neondb?sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var cta models.CTA
	db.First(&cta)

	newCta := models.CTA{
		TitleHi:       "आइए, मिलकर वैशाली नगर को और बेहतर बनाएं",
		TitleEn:       "Let's make Vaishali Nagar better together",
		SubtitleHi:    "आपका साथ, हमारा प्रयास - विकास की नई उड़ान। कमल सेतु के माध्यम से सीधे जुड़ें।",
		SubtitleEn:    "Your support, our effort - A new flight of development. Connect directly via Kamal Setu.",
		Button1TextHi: "कमल सेतु से जुड़ें",
		Button1TextEn: "Connect with Kamal Setu",
		Button1Link:   "https://wa.me/919977891333",
		Button2TextHi: "संपर्क करें",
		Button2TextEn: "Contact Us",
		Button2Link:   "/contact",
	}

	if cta.ID == 0 {
		db.Create(&newCta)
		log.Println("Created new CTA record")
	} else {
		newCta.ID = cta.ID
		db.Save(&newCta)
		log.Println("Updated existing CTA record")
	}
}
