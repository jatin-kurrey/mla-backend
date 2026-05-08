package main

import (
	"log"
	"os"
	"mla-backend/models"
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

	var all []models.Milestone
	db.Find(&all)
	log.Printf("Found %d milestones in DB\n", len(all))
	for _, m := range all {
		log.Printf("ID: %d, Year: %s, TitleHi: '%s', TitleEn: '%s'\n", m.ID, m.Year, m.TitleHi, m.TitleEn)
	}

	// Correct Bilingual Data Map
	data := map[string]struct{hiTitle, enTitle, hiDesc, enDesc string}{
		"2000": {"राजनीतिक शुरुआत", "Political Inception", "एक प्राथमिक सदस्य के रूप में भाजपा में शामिल हुए, समर्पण और विचारधारा की यात्रा शुरू की।", "Joined the BJP as a primary member, starting a journey of dedication and ideology."},
		"2004": {"शासन में पहला कदम", "First Step in Governance", "पहली बार पार्षद चुने गए, वार्ड स्तर के विकास पर ध्यान केंद्रित किया।", "Elected as Corporator (Parshad) for the first time, focused on ward-level development."},
		"2009": {"विश्वास के साथ पुनः निर्वाचित", "Re-elected with Trust", "दूसरी बार पार्षद का चुनाव जीता, पहले कार्यकाल में किए गए कार्यों की पुष्टि की।", "Won the corporator seat for the second time, validating the work done in the first term."},
		"2014": {"जीत की हैट्रिक", "Hat-trick Victory", "लगातार तीसरी बार पार्षद के रूप में जीत हासिल की, स्थानीय शासन में एक मजबूत आधार स्थापित किया।", "Third consecutive win as Corporator, establishing a strong base in local governance."},
		"2017": {"जनता की आवाज", "Voice of the People", "नगर निगम में विपक्ष के नेता (नेता प्रतिपक्ष) के रूप में नियुक्त।", "Appointed as the Leader of Opposition (Neta Pratipaksh) in the Municipal Corporation."},
		"2019": {"ऐतिहासिक 5वां कार्यकाल", "Historic 5th Term", "लगातार 5वीं बार पार्षद चुने जाकर एक दुर्लभ उपलब्धि हासिल की।", "Achieved a rare milestone by being elected as Corporator for the 5th consecutive time."},
		"2021": {"संकट के दौरान सेवा", "Service During Crisis", "कोविड-19 महामारी के दौरान व्यापक राहत कार्य और स्वास्थ्य पहल का नेतृत्व किया।", "Led extensive relief work and health initiatives during the COVID-19 pandemic."},
		"2023": {"वैशाली नगर के लिए जनादेश", "Mandate for Vaishali Nagar", "भारी जनमत के साथ विधायक (MLA) के रूप में निर्वाचित।", "Elected as Member of Legislative Assembly (MLA) with a massive public mandate."},
		"2024": {"मिशन आदर्श विधानसभा", "Mission Ideal Constituency", "वैशाली नगर को राज्य की एक आदर्श विधानसभा बनाने के लिए एक व्यापक रोडमैप लॉन्च किया।", "Launched a comprehensive roadmap to make Vaishali Nagar a model assembly of the state."},
	}

	for _, m := range all {
		if correct, ok := data[m.Year]; ok {
			db.Model(&m).Updates(map[string]interface{}{
				"title_hi":       correct.hiTitle,
				"title_en":       correct.enTitle,
				"description_hi": correct.hiDesc,
				"description_en": correct.enDesc,
			})
			log.Printf("Updated ID %d (%s)\n", m.ID, m.Year)
		} else {
			// If year not in map, maybe it has English in both?
			log.Printf("Year %s not in fix map, skipping\n", m.Year)
		}
	}

	log.Println("Data fix complete!")
}
