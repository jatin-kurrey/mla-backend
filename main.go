package main

import (
	"log"
	"os"
	"fmt"
	"mla-backend/handlers"
	"mla-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system env")
	}

	dbURL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	models.AutoMigrate(db)

	// Seed Admin User
	var userCount int64
	db.Model(&models.User{}).Count(&userCount)
	if userCount == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		admin := models.User{Username: "admin", Password: string(hashedPassword)}
		db.Create(&admin)
		log.Println("Admin user seeded: admin / admin123")
	}

	// Seed Settings
	var settingsCount int64
	db.Model(&models.Settings{}).Count(&settingsCount)
	if settingsCount == 0 {
		settings := models.Settings{
			ShowTopbar: true, ShowNavbar: true, ShowFooter: true,
			ShowHero: true, ShowHeroSlogan: true, ShowHeroButtons: true, ShowHeroMLA: true, ShowHeroServices: true,
			ShowStats: true, ShowNews: true, ShowSchemes: true,
			ShowGallery: true, ShowAbout: true, ShowTimeline: true, ShowSocial: true, ShowCTA: true,
		}
		db.Create(&settings)
		log.Println("Default settings seeded")
	}

	// Seed CTA
	var ctaCount int64
	db.Model(&models.CTA{}).Count(&ctaCount)
	if ctaCount == 0 {
		cta := models.CTA{
			TitleHi:       "वैशाली नगर को मिलकर बेहतर बनाएं",
			TitleEn:       "Let's make Vaishali Nagar better together",
			SubtitleHi:    "आपका साथ, हमारा प्रयास - विकास की नई उड़ान",
			SubtitleEn:    "Your support, our effort - A new flight of development",
			Button1TextHi: "स्वयंसेवक बनें",
			Button1TextEn: "Become a Volunteer",
			Button1Link:   "/janasamvad",
			Button2TextHi: "संपर्क करें",
			Button2TextEn: "Contact Us",
			Button2Link:   "/contact",
		}
		db.Create(&cta)
		log.Println("CTA data seeded")
	}

	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // 100MB
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	app.Use(logger.New())

	var aboutCount int64
	db.Model(&models.About{}).Count(&aboutCount)
	if aboutCount == 0 {
		about := models.About{
			TitleHi:       "दशकों का संघर्ष,",
			TitleEn:       "Decades of Struggle,",
			SubtitleHi:    "जनता का अटूट विश्वास",
			SubtitleEn:    "Unwavering Trust of the People",
			DescriptionHi: "रिकेश सेन, वैशाली नगर विधानसभा के एक समर्पित जनसेवक हैं। पिछले 20 वर्षों से अधिक समय से वे क्षेत्र के विकास और नागरिकों की समस्याओं के समाधान के लिए निरंतर सक्रिय रहे हैं। 5 बार के पार्षद रहने के दौरान उन्होंने जमीनी स्तर पर विकास की जो नींव रखी, आज वह विधायक के रूप में वैशाली नगर को एक 'आदर्श विधानसभा' बनाने के संकल्प के साथ आगे बढ़ रही है।",
			DescriptionEn: "Rikesh Sen is a dedicated public servant of Vaishali Nagar constituency. For over 20 years, he has been constantly active in the development of the area and solving citizens' problems. During his 5 terms as a Corporator, the foundation of development he laid at the grassroots level is now moving forward with the resolve to make Vaishali Nagar an 'Ideal Assembly' as an MLA.",
			Image:         "/images/DSC_7316.JPG",
			ExperienceHi:  "20+",
			ExperienceEn:  "20+",
			Point1Hi:      "5 बार पार्षद - अटूट जनविश्वास",
			Point1En:      "5-time Corporator - Unwavering Trust",
			Point2Hi:      "सक्रिय समाजसेवी एवं जनसेवक",
			Point2En:      "Active Social Worker & Public Servant",
			Point3Hi:      "वैशाली नगर विकास हेतु संकल्पित",
			Point3En:      "Committed to Vaishali Nagar Development",
			Point4Hi:      "BJP नेतृत्व का भरोसा एवं साथ",
			Point4En:      "Trust and Support of BJP Leadership",
			BirthplaceHi:  "भिलाई, छत्तीसगढ़",
			BirthplaceEn:  "Bhilai, Chhattisgarh",
			EducationHi:   "स्नातक",
			EducationEn:   "Graduate",
		}
		db.Create(&about)
		log.Println("Default about data seeded")
	}

	var ap models.AboutPage
	db.First(&ap)
	if ap.ID == 0 {
		ap = models.AboutPage{
			PageTitleHi:    "मेरे बारे में",
			PageTitleEn:    "About Me",
			PageSubtitleHi: "जनसेवा और समर्पण की एक गौरवशाली यात्रा",
			PageSubtitleEn: "A glorious journey of public service and dedication",
			IntroText1Hi:   "रिकेश सेन एक समर्पित जनप्रतिनिधि हैं जिन्होंने अपना जीवन जनसेवा को समर्पित कर दिया है। पाँच बार पार्षद रहने के बाद वैशाली नगर की जनता ने उन्हें विधायक के रूप में चुनकर अपना विश्वास व्यक्त किया।",
			IntroText1En:   "Rikesh Sen is a dedicated public representative who has dedicated his life to public service. After being a corporator five times, the people of Vaishali Nagar expressed their trust by electing him as an MLA.",
			IntroText2Hi:   "भारतीय जनता पार्टी के निष्ठावान कार्यकर्ता के रूप में, उनका मानना है कि सच्चा विकास तभी संभव है जब प्रतिनिधि जनता के बीच रहकर उनकी समस्याओं को समझे और हल करे।",
			IntroText2En:   "As a loyal worker of the Bharatiya Janata Party, he believes that true development is possible only when the representative stays among the people and understands and solves their problems.",
			BirthplaceHi:   "भिलाई, छत्तीसगढ़",
			BirthplaceEn:   "Bhilai, Chhattisgarh",
			EducationHi:    "स्नातक",
			EducationEn:    "Graduate",
			PartyHi:        "भारतीय जनता पार्टी",
			PartyEn:        "BJP",
			CurrentPosHi:   "विधायक, वैशाली नगर",
			CurrentPosEn:   "MLA, Vaishali Nagar",
			MainImage:      "/images/DSC_7316.JPG",
		}
		db.Create(&ap)
	}
	
	var milestoneCount int64
	db.Model(&models.Milestone{}).Count(&milestoneCount)
	if milestoneCount == 0 {
		milestones := []models.Milestone{
			{Year: "2000", TitleHi: "राजनीतिक शुरुआत", TitleEn: "Political Inception", DescriptionHi: "एक प्राथमिक सदस्य के रूप में भाजपा में शामिल हुए, समर्पण और विचारधारा की यात्रा शुरू की।", DescriptionEn: "Joined the BJP as a primary member, starting a journey of dedication and ideology.", Image: "/images/FB_IMG_1772985462804.jpg"},
			{Year: "2004", TitleHi: "शासन में पहला कदम", TitleEn: "First Step in Governance", DescriptionHi: "पहली बार पार्षद चुने गए, वार्ड स्तर के विकास पर ध्यान केंद्रित किया।", DescriptionEn: "Elected as Corporator (Parshad) for the first time, focused on ward-level development.", Image: "/images/DSC_7316.JPG"},
			{Year: "2009", TitleHi: "विश्वास के साथ पुनः निर्वाचित", TitleEn: "Re-elected with Trust", DescriptionHi: "दूसरी बार पार्षद का चुनाव जीता, पहले कार्यकाल में किए गए कार्यों की पुष्टि की।", DescriptionEn: "Won the corporator seat for the second time, validating the work done in the first term.", Image: "/images/DSC_8478.JPG"},
			{Year: "2014", TitleHi: "जीत की हैट्रिक", TitleEn: "Hat-trick Victory", DescriptionHi: "लगातार तीसरी बार पार्षद के रूप में जीत हासिल की, स्थानीय शासन में एक मजबूत आधार स्थापित किया।", DescriptionEn: "Third consecutive win as Corporator, establishing a strong base in local governance.", Image: "/images/DSC_8711.JPG"},
			{Year: "2017", TitleHi: "जनता की आवाज", TitleEn: "Voice of the People", DescriptionHi: "नगर निगम में विपक्ष के नेता (नेता प्रतिपक्ष) के रूप में नियुक्त।", DescriptionEn: "Appointed as the Leader of Opposition (Neta Pratipaksh) in the Municipal Corporation.", Image: "/images/DSC_7115.JPG"},
			{Year: "2019", TitleHi: "ऐतिहासिक 5वां कार्यकाल", TitleEn: "Historic 5th Term", DescriptionHi: "लगातार 5वीं बार पार्षद चुने जाकर एक दुर्लभ उपलब्धि हासिल की।", DescriptionEn: "Achieved a rare milestone by being elected as Corporator for the 5th consecutive time.", Image: "/images/IMG-20250718-WA0279.jpg"},
			{Year: "2021", TitleHi: "संकट के दौरान सेवा", TitleEn: "Service During Crisis", DescriptionHi: "कोविड-19 महामारी के दौरान व्यापक राहत कार्य और स्वास्थ्य पहल का नेतृत्व किया।", DescriptionEn: "Led extensive relief work and health initiatives during the COVID-19 pandemic.", Image: "/images/timeline_community_service_1777972662314.png"},
			{Year: "2023", TitleHi: "वैशाली नगर के लिए जनादेश", TitleEn: "Mandate for Vaishali Nagar", DescriptionHi: "भारी जनमत के साथ विधायक (MLA) के रूप में निर्वाचित।", DescriptionEn: "Elected as Member of Legislative Assembly (MLA) with a massive public mandate.", Image: "/images/IMG-20260209-WA0468.jpg"},
			{Year: "2024", TitleHi: "मिशन आदर्श विधानसभा", TitleEn: "Mission Ideal Constituency", DescriptionHi: "वैशाली नगर को राज्य की एक आदर्श विधानसभा बनाने के लिए एक व्यापक रोडमैप लॉन्च किया।", DescriptionEn: "Launched a comprehensive roadmap to make Vaishali Nagar a model assembly of the state.", Image: "/images/IMG-20260315-WA0164.jpg"},
		}
		for _, m := range milestones { db.Create(&m) }
	}


	var valuesCount int64
	db.Model(&models.Value{}).Count(&valuesCount)
	if valuesCount == 0 {
		values := []models.Value{
			{TitleHi: "पारदर्शी शासन", TitleEn: "Transparent Governance", TextHi: "हर नागरिक के लिए भ्रष्टाचार मुक्त और खुला प्रशासन सुनिश्चित करने की प्रतिबद्धता।", TextEn: "Commitment to corruption-free and open administration for every citizen.", Icon: "ShieldCheck"},
			{TitleHi: "समावेशी विकास", TitleEn: "Inclusive Development", TextHi: "यह सुनिश्चित करना कि प्रगति का लाभ सामाजिक सीढ़ी के अंतिम व्यक्ति तक पहुंचे।", TextEn: "Ensuring that the benefits of progress reach the last person in the social hierarchy.", Icon: "Users"},
			{TitleHi: "युवा सशक्तिकरण", TitleEn: "Youth Empowerment", TextHi: "युवा पीढ़ी के लिए शिक्षा, कौशल और रोजगार के अवसर पैदा करना।", TextEn: "Creating opportunities for education, skills, and employment for the younger generation.", Icon: "Zap"},
			{TitleHi: "सार्वजनिक पहुंच", TitleEn: "Public Accessibility", TextHi: "लोगों की समस्याओं को सुनने और हल करने के लिए 24/7 खुले दरवाजे की नीति बनाए रखना।", TextEn: "Maintaining a 24/7 open-door policy to listen and solve the problems of the people.", Icon: "MessageSquare"},
		}
		for _, v := range values { db.Create(&v) }
	}

	// Detailed Seeding for News
	var newsCount int64
	db.Model(&models.News{}).Count(&newsCount)
	if newsCount == 0 {
		news := []models.News{
			{TitleHi: "मेधावी छात्रों का सम्मान समारोह", TitleEn: "Meritorious Students Honor Ceremony", Image: "https://pub-e1f8fc9a8eab4bc8a2dcece0862b2ba0.r2.dev/news1.png", DateHi: "05 मई, 2024", DateEn: "May 05, 2024", CategoryHi: "शिक्षा", CategoryEn: "Education", DescriptionHi: "वैशाली नगर के प्रतिभावान छात्रों को विधायक द्वारा सम्मानित किया गया।", DescriptionEn: "Talented students of Vaishali Nagar were honored by the MLA.", ContentHi: "वैशाली नगर विधानसभा के मेधावी छात्र-छात्राओं के लिए आयोजित इस समारोह में विधायक रिकेश सेन ने 500 से अधिक छात्रों को प्रशस्ति पत्र और पुरस्कार प्रदान किए। उन्होंने छात्रों के उज्ज्वल भविष्य की कामना की।", ContentEn: "In this ceremony organized for meritorious students of Vaishali Nagar assembly, MLA Rikesh Sen presented certificates and awards to more than 500 students. He wished for the bright future of the students."},
			{TitleHi: "वैशाली नगर में नया स्वास्थ्य केंद्र जल्द", TitleEn: "New Health Center in Vaishali Nagar Soon", Image: "https://pub-e1f8fc9a8eab4bc8a2dcece0862b2ba0.r2.dev/news2.png", DateHi: "10 मई, 2024", DateEn: "May 10, 2024", CategoryHi: "स्वास्थ्य", CategoryEn: "Health", DescriptionHi: "क्षेत्रवासियों को बेहतर स्वास्थ्य सुविधाएं प्रदान करने के लिए नए केंद्र का शिलान्यास।", DescriptionEn: "Foundation stone of new center laid to provide better health facilities to the residents.", ContentHi: "वैशाली नगर वासियों को आधुनिक स्वास्थ्य सुविधाएं प्रदान करने के उद्देश्य से नए प्राथमिक स्वास्थ्य केंद्र का निर्माण कार्य युद्ध स्तर पर जारी है। विधायक ने बताया कि आगामी 3 महीनों में यह केंद्र जनता के लिए समर्पित कर दिया जाएगा।", ContentEn: "The construction work of the new primary health center is going on at war footing with the aim of providing modern health facilities to the residents of Vaishali Nagar. The MLA said that in the next 3 months this center will be dedicated to the public."},
			{TitleHi: "विधायक रिकेश सेन ने किया विकास कार्यों का निरीक्षण", TitleEn: "MLA Rikesh Sen Inspected Development Works", Image: "https://pub-e1f8fc9a8eab4bc8a2dcece0862b2ba0.r2.dev/news3.png", DateHi: "15 मई, 2024", DateEn: "May 15, 2024", CategoryHi: "विकास", CategoryEn: "Development", DescriptionHi: "सड़क और नाली निर्माण कार्यों की गुणवत्ता की जांच की गई।", DescriptionEn: "The quality of road and drain construction works was checked.", ContentHi: "विधायक रिकेश सेन ने आज वैशाली नगर के विभिन्न वार्डों का दौरा कर चल रहे सड़क डामरीकरण और नाली निर्माण कार्यों का निरीक्षण किया। उन्होंने निर्माण की गुणवत्ता पर विशेष ध्यान देने और समय सीमा में कार्य पूरा करने के निर्देश दिए।", ContentEn: "MLA Rikesh Sen today visited various wards of Vaishali Nagar and inspected the ongoing road tarring and drain construction works. He directed to pay special attention to the quality of construction and complete the work within the time limit."},
		}
		for _, n := range news { db.Create(&n) }
	}

	// Detailed Seeding for Stats
	var statsCount int64
	db.Model(&models.Stat{}).Count(&statsCount)
	if statsCount == 0 {
		stats := []models.Stat{
			{Label: "Development Budget", Value: "₹250Cr+", Sub: "Allocated for infrastructure projects", Icon: "Construction"},
			{Label: "Beneficiaries", Value: "10K+", Sub: "Families touched by schemes", Icon: "Users"},
			{Label: "Drinking Water", Value: "3500+", Sub: "New tap connections installed", Icon: "Droplets"},
			{Label: "Street Lights", Value: "5000+", Sub: "LED lights across Vaishali Nagar", Icon: "Zap"},
			{Label: "Schools Renovated", Value: "15+", Sub: "Smart classrooms & facilities", Icon: "GraduationCap"},
		}
		for _, s := range stats { db.Create(&s) }
	}

	// Detailed Seeding for Schemes
	var schemeCount int64
	db.Model(&models.Scheme{}).Count(&schemeCount)
	if schemeCount == 0 {
		schemes := []models.Scheme{
			{Icon: "HeartPulse", TitleHi: "आयुष्मान भारत", TitleEn: "Ayushman Bharat", DescHi: "₹5 लाख तक का निःशुल्क स्वास्थ्य बीमा कवर।", DescEn: "Free medical treatment up to 5 Lakhs for eligible families.", CategoryHi: "स्वास्थ्य", CategoryEn: "Health", Color: "from-orange-500 to-amber-500"},
			{Icon: "Home", TitleHi: "PM आवास योजना", TitleEn: "PM Awas Yojana", DescHi: "गरीब परिवारों के लिए पक्के मकान का सपना।", DescEn: "Providing permanent housing to every needy family in the constituency.", CategoryHi: "आवास", CategoryEn: "Housing", Color: "from-orange-500 to-amber-500"},
			{Icon: "Baby", TitleHi: "कन्या विवाह योजना", TitleEn: "Kanya Vivah Yojana", DescHi: "गरीब परिवारों की बेटियों के विवाह के लिए वित्तीय सहायता।", DescEn: "Financial assistance for the marriage of daughters from poor families.", CategoryHi: "सामाजिक", CategoryEn: "Social", Color: "from-orange-500 to-amber-500"},
			{Icon: "Users", TitleHi: "उज्ज्वला योजना", TitleEn: "Ujjwala Yojana", DescHi: "महिलाओं को मुफ्त गैस कनेक्शन और धुआं मुक्त रसोई।", DescEn: "Free gas connections to empower women and ensure smoke-free kitchens.", CategoryHi: "महिलाएं", CategoryEn: "Women", Color: "from-orange-500 to-amber-500"},
			{Icon: "GraduationCap", TitleHi: "शिक्षा छात्रवृत्ति", TitleEn: "Education Scholarship", DescHi: "मेधावी छात्रों के लिए छात्रवृत्ति एवं निःशुल्क पुस्तकें।", DescEn: "Scholarships and free books for meritorious students.", CategoryHi: "युवा", CategoryEn: "Youth", Color: "from-orange-500 to-amber-500"},
			{Icon: "Briefcase", TitleHi: "मुद्रा लोन", TitleEn: "Mudra Loan", DescHi: "स्वरोजगार के लिए ₹10 लाख तक का ऋण।", DescEn: "Loans up to ₹10 Lakhs for self-employment.", CategoryHi: "रोजगार", CategoryEn: "Employment", Color: "from-orange-500 to-amber-500"},
		}
		for _, s := range schemes { db.Create(&s) }
	}

	var galleryCount int64
	db.Model(&models.Gallery{}).Count(&galleryCount)
	if galleryCount == 0 {
		gallery := []models.Gallery{
			{Title: "गणेश चतुर्थी उत्सव", Image: "https://images.unsplash.com/photo-1567591974574-e852631b1813?q=80&w=2070", Category: "Events"},
			{Title: "विकास कार्यों का निरीक्षण", Image: "https://images.unsplash.com/photo-1541888946425-d81bb19480c5?q=80&w=2070", Category: "Work"},
			{Title: "जनता से संवाद", Image: "https://images.unsplash.com/photo-1517048676732-d65bc937f952?q=80&w=2070", Category: "Public"},
		}
		for _, g := range gallery { db.Create(&g) }
	}

	var socialCount int64
	db.Model(&models.SocialLink{}).Count(&socialCount)
	if socialCount == 0 {
		socials := []models.SocialLink{
			{Platform: "youtube", Title: "YouTube Live", URL: "https://www.youtube.com/@RIKESH_SEN_BJP", EmbedCode: "https://www.youtube.com/embed/videoseries?list=PL..." },
			{Platform: "twitter", Title: "Twitter Updates", URL: "https://twitter.com/rikeshsen", EmbedCode: "" },
			{Platform: "facebook", Title: "Facebook Feed", URL: "https://www.facebook.com/SenRikesh/", EmbedCode: "" },
			{Platform: "instagram", Title: "@rikeshsenbjp", URL: "https://www.instagram.com/rikeshsenbjp/", EmbedCode: "" },
		}
		for _, s := range socials { db.Create(&s) }
	}

	// Detailed Seeding for Hero
	var h models.Hero
	db.First(&h)
	if h.ID == 0 {
		h = models.Hero{
			Slogan1Hi: "विकास की नई राह,",
			Slogan1En: "New Path of Development,",
			Slogan2Hi: "वैशाली नगर का नया उत्साह।",
			Slogan2En: "New Enthusiasm of Vaishali Nagar.",
			DescHi:    "रिकेश सेन के नेतृत्व में वैशाली नगर विधानसभा क्षेत्र अब प्रगति के नए आयाम छू रहा है। हमारा संकल्प - सबका साथ, सबका विकास।",
			DescEn:    "Under the leadership of Rikesh Sen, Vaishali Nagar assembly constituency is now reaching new heights of progress. Our resolve - Sabka Saath, Sabka Vikas.",
			MLAImage:  "/images/DSC_7316.JPG",
			VideoURL:  "https://pub-e1f8fc9a8eab4bc8a2dcece0862b2ba0.r2.dev/rikesh_sen_hero.mp4",
		}
		db.Create(&h)
	}

	// Detailed Seeding for Development
	var devCount int64
	db.Model(&models.Development{}).Count(&devCount)
	if devCount == 0 {
		devs := []models.Development{
			{Icon: "Construction", TitleHi: "सड़क निर्माण", TitleEn: "Road Construction", Count: "125+", DescHi: "वैशाली नगर में सड़कों का चौड़ीकरण और मरम्मत कार्य पूर्ण किया गया।", DescEn: "Road widening and repair work completed in Vaishali Nagar.", Color: "from-orange-500 to-amber-500"},
			{Icon: "Droplets", TitleHi: "जल आपूर्ति", TitleEn: "Water Supply", Count: "85+", DescHi: "नई पाइपलाइन और जल वितरण प्रणाली स्थापित की गई।", DescEn: "New pipeline and water distribution system established.", Color: "from-blue-500 to-cyan-500"},
			{Icon: "GraduationCap", TitleHi: "शिक्षा", TitleEn: "Education", Count: "60+", DescHi: "स्कूलों का जीर्णोद्धार और स्मार्ट कक्षाओं की स्थापना।", DescEn: "Renovation of schools and establishment of smart classrooms.", Color: "from-emerald-500 to-teal-500"},
			{Icon: "HeartPulse", TitleHi: "स्वास्थ्य", TitleEn: "Health", Count: "40+", DescHi: "स्वास्थ्य केंद्रों का आधुनिकीकरण और निःशुल्क शिविर।", DescEn: "Modernization of health centers and free camps.", Color: "from-rose-500 to-pink-500"},
			{Icon: "Zap", TitleHi: "विद्युत व्यवस्था", TitleEn: "Electricity", Count: "70+", DescHi: "स्ट्रीट लाइट, ट्रांसफार्मर और नई विद्युत लाइनें।", DescEn: "Street lights, transformers and new power lines.", Color: "from-yellow-500 to-orange-500"},
			{Icon: "Briefcase", TitleHi: "युवा रोजगार", TitleEn: "Youth Employment", Count: "30+", DescHi: "रोजगार मेले और कौशल विकास कार्यक्रम आयोजित।", DescEn: "Employment fairs and skill development programs organized.", Color: "from-violet-500 to-purple-500"},
			{Icon: "TreePine", TitleHi: "हरित क्षेत्र", TitleEn: "Green Zone", Count: "20+", DescHi: "पार्कों का विकास और वृक्षारोपण अभियान।", DescEn: "Development of parks and plantation campaign.", Color: "from-green-500 to-emerald-500"},
			{Icon: "Building2", TitleHi: "सामुदायिक भवन", TitleEn: "Community Building", Count: "15+", DescHi: "नए सामुदायिक केंद्र और भवनों का निर्माण।", DescEn: "Construction of new community centers and buildings.", Color: "from-indigo-500 to-blue-500"},
			{Icon: "Bus", TitleHi: "परिवहन", TitleEn: "Transport", Count: "12+", DescHi: "बस स्टॉप, यातायात व्यवस्था में सुधार।", DescEn: "Improvement in bus stops and traffic system.", Color: "from-fuchsia-500 to-pink-500"},
		}
		for _, d := range devs { db.Create(&d) }
	}

	// Ward Work Seeding
	var wardCount int64
	db.Model(&models.WardWork{}).Count(&wardCount)
	if wardCount < 37 {
		for i := 1; i <= 37; i++ {
			var existing models.WardWork
			db.Where("ward_num = ?", i).First(&existing)
			if existing.ID == 0 {
				db.Create(&models.WardWork{
					WardNum: i, 
					ProjectCount: 15 + (i % 10) * 5, 
					Details: fmt.Sprintf("वार्ड %d में सड़क डामरीकरण, नाली निर्माण एवं स्ट्रीट लाइट की व्यवस्था सुचारू रूप से की गई है।", i),
				})
			}
		}
	}

	api := app.Group("/api")

	// Auth & Upload
	api.Post("/login", handlers.Login(db))
	api.Post("/upload", handlers.UploadFile())
	api.Post("/delete-file", handlers.DeleteFile())

	// News
	api.Get("/news", func(c *fiber.Ctx) error {
		var news []models.News
		db.Order("created_at desc").Find(&news)
		return c.JSON(news)
	})
	api.Post("/news", handlers.CreateItem(db, &models.News{}))
	api.Put("/news/:id", handlers.UpdateItem(db, &models.News{}))
	api.Delete("/news/:id", handlers.DeleteItem(db, &models.News{}))

	// Stats
	api.Get("/stats", handlers.GetItems(db, &[]models.Stat{}))
	api.Post("/stats", handlers.CreateItem(db, &models.Stat{}))
	api.Put("/stats/:id", handlers.UpdateItem(db, &models.Stat{}))

	// Schemes
	api.Get("/schemes", handlers.GetItems(db, &[]models.Scheme{}))
	api.Post("/schemes", handlers.CreateItem(db, &models.Scheme{}))
	api.Put("/schemes/:id", handlers.UpdateItem(db, &models.Scheme{}))

	// Gallery
	api.Get("/gallery", handlers.GetItems(db, &[]models.Gallery{}))
	api.Post("/gallery", handlers.CreateItem(db, &models.Gallery{}))
	api.Delete("/gallery/:id", handlers.DeleteItem(db, &models.Gallery{}))

	// Social Media
	api.Get("/socials", handlers.GetItems(db, &[]models.SocialLink{}))
	api.Post("/socials", handlers.CreateItem(db, &models.SocialLink{}))
	api.Put("/socials/:id", handlers.UpdateItem(db, &models.SocialLink{}))
	api.Delete("/socials/:id", handlers.DeleteItem(db, &models.SocialLink{}))

	// Janasamvad (Complaints)
	api.Get("/complaints", handlers.GetItems(db, &[]models.Janasamvad{}))
	api.Post("/complaints", handlers.CreateItem(db, &models.Janasamvad{}))
	api.Put("/complaints/:id", handlers.UpdateItem(db, &models.Janasamvad{}))

	// Hero
	api.Get("/hero", func(c *fiber.Ctx) error {
		var hero models.Hero
		if err := db.First(&hero).Error; err != nil && err != gorm.ErrRecordNotFound {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(hero)
	})
	api.Put("/hero", func(c *fiber.Ctx) error {
		var hero models.Hero
		if err := c.BodyParser(&hero); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		
		var existing models.Hero
		db.First(&existing)
		if existing.ID == 0 {
			db.Create(&hero)
		} else {
			hero.ID = existing.ID
			db.Save(&hero)
		}
		return c.JSON(hero)
	})

	// Settings
	api.Get("/settings", func(c *fiber.Ctx) error {
		var settings models.Settings
		if err := db.First(&settings).Error; err != nil && err != gorm.ErrRecordNotFound {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(settings)
	})
	api.Put("/settings", func(c *fiber.Ctx) error {
		var settings models.Settings
		if err := c.BodyParser(&settings); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		db.Save(&settings)
		return c.JSON(settings)
	})

	// About
	api.Get("/about", func(c *fiber.Ctx) error {
		var about models.About
		if err := db.First(&about).Error; err != nil && err != gorm.ErrRecordNotFound {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(about)
	})
	api.Put("/about", func(c *fiber.Ctx) error {
		var about models.About
		if err := c.BodyParser(&about); err != nil { return c.Status(400).JSON(fiber.Map{"error": err.Error()}) }
		db.Save(&about)
		return c.JSON(about)
	})

	// About Page
	api.Get("/about-page", func(c *fiber.Ctx) error {
		var ap models.AboutPage
		if err := db.First(&ap).Error; err != nil && err != gorm.ErrRecordNotFound {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(ap)
	})
	api.Put("/about-page", func(c *fiber.Ctx) error {
		var ap models.AboutPage
		if err := c.BodyParser(&ap); err != nil { return c.Status(400).JSON(fiber.Map{"error": err.Error()}) }
		db.Save(&ap)
		return c.JSON(ap)
	})
	// CTA
	api.Get("/cta", func(c *fiber.Ctx) error {
		var cta models.CTA
		if err := db.First(&cta).Error; err != nil && err != gorm.ErrRecordNotFound {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(cta)
	})
	api.Put("/cta", func(c *fiber.Ctx) error {
		var cta models.CTA
		if err := c.BodyParser(&cta); err != nil { return c.Status(400).JSON(fiber.Map{"error": err.Error()}) }
		db.Save(&cta)
		return c.JSON(cta)
	})

	// Development
	api.Get("/development", func(c *fiber.Ctx) error {
		var devs []models.Development
		db.Find(&devs)
		return c.JSON(devs)
	})

	api.Post("/development", func(c *fiber.Ctx) error {
		var dev models.Development
		if err := c.BodyParser(&dev); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		db.Create(&dev)
		return c.JSON(dev)
	})

	api.Put("/development/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var dev models.Development
		if err := c.BodyParser(&dev); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		db.Model(&models.Development{}).Where("id = ?", id).Updates(dev)
		return c.JSON(dev)
	})

	api.Delete("/development/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		db.Delete(&models.Development{}, id)
		return c.SendStatus(204)
	})

	// Ward Work
	api.Get("/ward-work", func(c *fiber.Ctx) error {
		var wards []models.WardWork
		db.Order("ward_num asc").Find(&wards)
		return c.JSON(wards)
	})

	api.Post("/ward-work", func(c *fiber.Ctx) error {
		var w models.WardWork
		if err := c.BodyParser(&w); err != nil { return c.Status(400).JSON(fiber.Map{"error": err.Error()}) }
		db.Create(&w)
		return c.JSON(w)
	})

	api.Put("/ward-work/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var w models.WardWork
		if err := c.BodyParser(&w); err != nil { return c.Status(400).JSON(fiber.Map{"error": err.Error()}) }
		db.Model(&models.WardWork{}).Where("id = ?", id).Updates(w)
		return c.JSON(w)
	})

	api.Delete("/ward-work/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		db.Delete(&models.WardWork{}, id)
		return c.SendStatus(204)
	})

	// Milestones
	api.Get("/milestones", func(c *fiber.Ctx) error {
		var m []models.Milestone
		db.Order("year asc").Find(&m)
		return c.JSON(m)
	})
	api.Post("/milestones", func(c *fiber.Ctx) error {
		var m models.Milestone
		if err := c.BodyParser(&m); err != nil { return c.Status(400).JSON(fiber.Map{"error": err.Error()}) }
		db.Create(&m)
		return c.JSON(m)
	})
	api.Put("/milestones/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var m models.Milestone
		db.First(&m, id)
		if err := c.BodyParser(&m); err != nil { return c.Status(400).JSON(fiber.Map{"error": err.Error()}) }
		db.Save(&m)
		return c.JSON(m)
	})
	api.Delete("/milestones/:id", func(c *fiber.Ctx) error {
		db.Delete(&models.Milestone{}, c.Params("id"))
		return c.SendStatus(204)
	})

	// Values
	api.Get("/values", func(c *fiber.Ctx) error {
		var v []models.Value
		db.Find(&v)
		return c.JSON(v)
	})
	api.Put("/values", func(c *fiber.Ctx) error {
		var v []models.Value
		if err := c.BodyParser(&v); err != nil { return c.Status(400).JSON(fiber.Map{"error": err.Error()}) }
		db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Value{})
		for _, val := range v { db.Create(&val) }
		return c.JSON(v)
	})


	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}
	log.Fatal(app.Listen(":" + port))
}
