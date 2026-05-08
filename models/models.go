package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"unique;not null" json:"username"`
	Password  string         `gorm:"not null" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type Hero struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Slogan1Hi string `json:"slogan1_hi"`
	Slogan1En string `json:"slogan1_en"`
	Slogan2Hi string `json:"slogan2_hi"`
	Slogan2En string `json:"slogan2_en"`
	DescHi    string `json:"desc_hi"`
	DescEn    string `json:"desc_en"`
	VideoURL  string `json:"video_url"`
	MLAImage  string `json:"mla_image"`
}

type News struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TitleHi       string    `json:"title_hi"`
	TitleEn       string    `json:"title_en"`
	Image         string    `json:"image"`
	DateHi        string    `json:"date_hi"`
	DateEn        string    `json:"date_en"`
	DescriptionHi string    `json:"description_hi"`
	DescriptionEn string    `json:"description_en"`
	ContentHi     string    `json:"content_hi"`
	ContentEn     string    `json:"content_en"`
	CategoryHi    string    `json:"category_hi"`
	CategoryEn    string    `json:"category_en"`
	CreatedAt     time.Time `json:"created_at"`
}

type Gallery struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Image     string    `gorm:"column:image_url" json:"image"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

type SocialLink struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Platform  string    `json:"platform"` // youtube, twitter, facebook, instagram
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	EmbedCode string    `json:"embed_code"` // for youtube videos
	CreatedAt time.Time `json:"created_at"`
}

type Stat struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Icon      string `json:"icon"`
	Label     string `json:"label"`
	Value     string `json:"value"`
	Sub       string `json:"sub"`
}

type Scheme struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Icon       string `json:"icon"`
	TitleHi    string `json:"title_hi"`
	TitleEn    string `json:"title_en"`
	Color      string `json:"color"`
	DescHi     string `json:"desc_hi"`
	DescEn     string `json:"desc_en"`
	CategoryHi string `json:"category_hi"`
	CategoryEn string `json:"category_en"`
}

type Janasamvad struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Email     string    `json:"email"`
	Type      string    `json:"type"`
	Ward      string    `json:"ward"`
	Message   string    `json:"message"`
	Status    string    `gorm:"default:'pending'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Settings struct {
	ID            uint `gorm:"primaryKey" json:"id"`
	ShowTopbar    bool `gorm:"default:true" json:"show_topbar"`
	ShowNavbar    bool `gorm:"default:true" json:"show_navbar"`
	ShowHero      bool `gorm:"default:true" json:"show_hero"`
	ShowHeroSlogan bool `gorm:"default:true" json:"show_hero_slogan"`
	ShowHeroButtons bool `gorm:"default:true" json:"show_hero_buttons"`
	ShowHeroMLA   bool `gorm:"default:true" json:"show_hero_mla"`
	ShowHeroServices bool `gorm:"default:true" json:"show_hero_services"`
	ShowStats     bool `gorm:"default:true" json:"show_stats"`
	ShowNews      bool `gorm:"default:true" json:"show_news"`
	ShowSchemes   bool `gorm:"default:true" json:"show_schemes"`
	ShowGallery   bool `gorm:"default:true" json:"show_gallery"`
	ShowAbout     bool `gorm:"default:true" json:"show_about"`
	ShowTimeline  bool `gorm:"default:true" json:"show_timeline"`
	ShowSocial    bool `gorm:"default:true" json:"show_social"`
	ShowFooter    bool `gorm:"default:true" json:"show_footer"`
	ShowCTA       bool `gorm:"default:true" json:"show_cta"`
}

type CTA struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	TitleHi       string `json:"title_hi"`
	TitleEn       string `json:"title_en"`
	SubtitleHi    string `json:"subtitle_hi"`
	SubtitleEn    string `json:"subtitle_en"`
	Button1TextHi string `json:"button1_text_hi"`
	Button1TextEn string `json:"button1_text_en"`
	Button1Link   string `json:"button1_link"`
	Button2TextHi string `json:"button2_text_hi"`
	Button2TextEn string `json:"button2_text_en"`
	Button2Link   string `json:"button2_link"`
}

type About struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	TitleHi       string `json:"title_hi"`
	TitleEn       string `json:"title_en"`
	SubtitleHi    string `json:"subtitle_hi"`
	SubtitleEn    string `json:"subtitle_en"`
	DescriptionHi string `json:"description_hi"`
	DescriptionEn string `json:"description_en"`
	Image         string `json:"image"`
	ExperienceHi  string `json:"experience_hi"`
	ExperienceEn  string `json:"experience_en"`
	Point1Hi      string `json:"point1_hi"`
	Point1En      string `json:"point1_en"`
	Point2Hi      string `json:"point2_hi"`
	Point2En      string `json:"point2_en"`
	Point3Hi      string `json:"point3_hi"`
	Point3En      string `json:"point3_en"`
	Point4Hi      string `json:"point4_hi"`
	Point4En      string `json:"point4_en"`
	BirthplaceHi  string `json:"birthplace_hi"`
	BirthplaceEn  string `json:"birthplace_en"`
	EducationHi   string `json:"education_hi"`
	EducationEn   string `json:"education_en"`
}

type Milestone struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Year          string `json:"year"`
	TitleHi       string `json:"title_hi"`
	TitleEn       string `json:"title_en"`
	DescriptionHi string `json:"description_hi"`
	DescriptionEn string `json:"description_en"`
	Image         string `json:"image"`
}

type Value struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	TitleHi string `json:"title_hi"`
	TitleEn string `json:"title_en"`
	TextHi  string `json:"text_hi"`
	TextEn  string `json:"text_en"`
	Icon    string `json:"icon"` // icon name from lucide
}

type AboutPage struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	PageTitleHi    string `json:"page_title_hi"`
	PageTitleEn    string `json:"page_title_en"`
	PageSubtitleHi string `json:"page_subtitle_hi"`
	PageSubtitleEn string `json:"page_subtitle_en"`
	IntroText1Hi   string `json:"intro_text_1_hi"`
	IntroText1En   string `json:"intro_text_1_en"`
	IntroText2Hi   string `json:"intro_text_2_hi"`
	IntroText2En   string `json:"intro_text_2_en"`
	BirthplaceHi   string `json:"birthplace_hi"`
	BirthplaceEn   string `json:"birthplace_en"`
	EducationHi    string `json:"education_hi"`
	EducationEn    string `json:"education_en"`
	PartyHi        string `json:"party_hi"`
	PartyEn        string `json:"party_en"`
	CurrentPosHi   string `json:"current_pos_hi"`
	CurrentPosEn   string `json:"current_pos_en"`
	MainImage      string `json:"main_image"`
}

type Development struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	TitleHi string `json:"title_hi"`
	TitleEn string `json:"title_en"`
	DescHi  string `json:"desc_hi"`
	DescEn  string `json:"desc_en"`
	Count   string `json:"count"`
	Icon    string `json:"icon"`
	Color   string `json:"color"`
}

type WardWork struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	WardNum      int    `json:"ward_num"`
	ProjectCount int    `json:"project_count"`
	Details      string `json:"details"`
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Hero{}, &News{}, &Stat{}, &Scheme{}, &Gallery{}, &Janasamvad{}, &Settings{}, &About{}, &Milestone{}, &Value{}, &AboutPage{}, &Development{}, &WardWork{}, &SocialLink{}, &CTA{})
}
