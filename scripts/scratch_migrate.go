package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migrating Hero...")
	db.Exec("UPDATE heros SET slogan1_hi = slogan1, slogan1_en = slogan1, slogan2_hi = slogan2, slogan2_en = slogan2, desc_hi = \"desc\", desc_en = \"desc\" WHERE slogan1_hi IS NULL OR slogan1_hi = ''")

	log.Println("Migrating News...")
	db.Exec("UPDATE news SET title_hi = title, title_en = title, date_hi = date, date_en = date, description_hi = description, description_en = description, content_hi = content, content_en = content WHERE title_hi IS NULL OR title_hi = ''")

	log.Println("Migrating Schemes...")
	db.Exec("UPDATE schemes SET title_hi = title, title_en = title, desc_hi = \"desc\", desc_en = \"desc\", category_hi = category, category_en = category WHERE title_hi IS NULL OR title_hi = ''")

	log.Println("Migrating About...")
	db.Exec("UPDATE abouts SET title_hi = title, title_en = title, subtitle_hi = subtitle, subtitle_en = subtitle, description_hi = description, description_en = description, experience_hi = experience, experience_en = experience, point1_hi = point1, point1_en = point1, point2_hi = point2, point2_en = point2, point3_hi = point3, point3_en = point3, point4_hi = point4, point4_en = point4, birthplace_hi = birthplace, birthplace_en = birthplace, education_hi = education, education_en = education WHERE title_hi IS NULL OR title_hi = ''")

	log.Println("Migrating AboutPage...")
	db.Exec("UPDATE about_pages SET page_title_hi = page_title, page_title_en = page_title, page_subtitle_hi = page_subtitle, page_subtitle_en = page_subtitle, intro_text1_hi = intro_text1, intro_text1_en = intro_text1, intro_text2_hi = intro_text2, intro_text2_en = intro_text2, birthplace_hi = birthplace, birthplace_en = birthplace, education_hi = education, education_en = education, party_hi = party, party_en = party, current_pos_hi = current_pos, current_pos_en = current_pos WHERE page_title_hi IS NULL OR page_title_hi = ''")

	log.Println("Migrating Milestones...")
	db.Exec("UPDATE milestones SET title_hi = title, title_en = title, description_hi = description, description_en = description WHERE title_hi IS NULL OR title_hi = ''")

	log.Println("Migrating Values...")
	db.Exec("UPDATE values SET title_hi = title, title_en = title, text_hi = text, text_en = text WHERE title_hi IS NULL OR title_hi = ''")

	log.Println("Migrating Development...")
	db.Exec("UPDATE developments SET title_hi = title, title_en = title, desc_hi = \"desc\", desc_en = \"desc\" WHERE title_hi IS NULL OR title_hi = ''")

	log.Println("Migrating CTA...")
	db.Exec("UPDATE cta SET title_hi = title, title_en = title, subtitle_hi = subtitle, subtitle_en = subtitle, button1_text_hi = button1_text, button1_text_en = button1_text, button2_text_hi = button2_text, button2_text_en = button2_text WHERE title_hi IS NULL OR title_hi = ''")

	log.Println("Done!")
}
