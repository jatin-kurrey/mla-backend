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

	// Update Hero
	var hero models.Hero
	db.First(&hero)
	if hero.ID != 0 {
		db.Model(&hero).Updates(map[string]interface{}{
			"slogan1_hi": "विकास की नई राह,",
			"slogan2_hi": "वैशाली नगर का नया उत्साह।",
			"desc_hi":    "रिकेश सेन के नेतृत्व में वैशाली नगर विधानसभा क्षेत्र अब प्रगति के नए आयाम छू रहा है। हमारा संकल्प - सबका साथ, सबका विकास।",
		})
		log.Println("Updated Hero Hindi data")
	}

	// Update About (Home Page)
	var about models.About
	db.First(&about)
	if about.ID != 0 {
		db.Model(&about).Updates(map[string]interface{}{
			"title_hi":       "दशकों का संघर्ष,",
			"subtitle_hi":    "जनता का अटूट विश्वास",
			"description_hi": "रिकेश सेन, वैशाली नगर विधानसभा के एक समर्पित जनसेवक हैं। पिछले 20 वर्षों से अधिक समय से वे क्षेत्र के विकास और नागरिकों की समस्याओं के समाधान के लिए निरंतर सक्रिय रहे हैं।",
			"experience_hi":  "20+",
			"point1_hi":      "5 बार पार्षद - अटूट जनविश्वास",
			"point2_hi":      "सक्रिय समाजसेवी एवं जनसेवक",
			"point3_hi":      "वैशाली नगर विकास हेतु संकल्पित",
			"point4_hi":      "BJP नेतृत्व का भरोसा एवं साथ",
			"birthplace_hi":  "भिलाई, छत्तीसगढ़",
			"education_hi":   "स्नातक",
		})
		log.Println("Updated About (Home) Hindi data")
	}

	// Update About Page
	var ap models.AboutPage
	db.First(&ap)
	if ap.ID != 0 {
		db.Model(&ap).Updates(map[string]interface{}{
			"page_title_hi":    "मेरे बारे में",
			"page_subtitle_hi": "जनसेवा और समर्पण की एक गौरवशाली यात्रा",
			"intro_text_1_hi":   "रिकेश सेन एक समर्पित जनप्रतिनिधि हैं जिन्होंने अपना जीवन जनसेवा को समर्पित कर दिया है। पाँच बार पार्षद रहने के बाद वैशाली नगर की जनता ने उन्हें विधायक के रूप में चुनकर अपना विश्वास व्यक्त किया।",
			"intro_text_2_hi":   "भारतीय जनता पार्टी के निष्ठावान कार्यकर्ता के रूप में, उनका मानना है कि सच्चा विकास तभी संभव है जब प्रतिनिधि जनता के बीच रहकर उनकी समस्याओं को समझे और हल करे।",
			"birthplace_hi":   "भिलाई, छत्तीसगढ़",
			"education_hi":    "स्नातक",
			"party_hi":        "भारतीय जनता पार्टी",
			"current_pos_hi":   "विधायक, वैशाली नगर",
		})
		log.Println("Updated About Page Hindi data")
	}

	// Update CTA
	var cta models.CTA
	db.First(&cta)
	if cta.ID != 0 {
		db.Model(&cta).Updates(map[string]interface{}{
			"title_hi":       "वैशाली नगर को मिलकर बेहतर बनाएं",
			"subtitle_hi":    "आपका साथ, हमारा प्रयास - विकास की नई उड़ान",
			"button1_text_hi": "स्वयंसेवक बनें",
			"button2_text_hi": "संपर्क करें",
		})
		log.Println("Updated CTA Hindi data")
	}

	// Update Development Items
	var devs []models.Development
	db.Find(&devs)
	for _, d := range devs {
		if d.TitleEn == "Road Construction" {
			db.Model(&d).Updates(map[string]interface{}{"title_hi": "सड़क निर्माण", "desc_hi": "वैशाली नगर में सड़कों का चौड़ीकरण और मरम्मत कार्य पूर्ण किया गया।"})
		} else if d.TitleEn == "Water Supply" {
			db.Model(&d).Updates(map[string]interface{}{"title_hi": "जल आपूर्ति", "desc_hi": "नई पाइपलाइन और जल वितरण प्रणाली स्थापित की गई।"})
		} else if d.TitleEn == "Education" {
			db.Model(&d).Updates(map[string]interface{}{"title_hi": "शिक्षा", "desc_hi": "स्कूलों का जीर्णोद्धार और स्मार्ट कक्षाओं की स्थापना।"})
		} else if d.TitleEn == "Health" {
			db.Model(&d).Updates(map[string]interface{}{"title_hi": "स्वास्थ्य", "desc_hi": "स्वास्थ्य केंद्रों का आधुनिकीकरण और निःशुल्क शिविर।"})
		}
	}
	log.Println("Updated Development Hindi data")

	// Update Schemes
	var schemes []models.Scheme
	db.Find(&schemes)
	for _, s := range schemes {
		if s.TitleEn == "Ayushman Bharat" {
			db.Model(&s).Updates(map[string]interface{}{"title_hi": "आयुष्मान भारत", "desc_hi": "₹5 लाख तक का निःशुल्क स्वास्थ्य बीमा कवर।", "category_hi": "स्वास्थ्य"})
		} else if s.TitleEn == "PM Awas Yojana" {
			db.Model(&s).Updates(map[string]interface{}{"title_hi": "PM आवास योजना", "desc_hi": "गरीब परिवारों के लिए पक्के मकान का सपना।", "category_hi": "आवास"})
		}
	}
	log.Println("Updated Schemes Hindi data")

	log.Println("Comprehensive Hindi Data update complete!")
}
