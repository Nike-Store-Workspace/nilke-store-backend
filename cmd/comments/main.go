package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt" // برای هش کردن پسورد کاربران
)

type CommentTemplate struct {
	TitleFa string
	TitleEn string
	BodyFa  string
	BodyEn  string
	Rating  int
}

var commentTemplates = []CommentTemplate{
	{"کیفیت عالی", "Excellent quality", "جنسش خیلی خوبه و کاملاً ارزش خرید داره. پیشنهاد می‌کنم.", "The material is very good and totally worth buying. I recommend it.", 5},
	{"معمولی بود", "It was average", "نسبت به قیمتش بد نیست ولی شگفت‌زده‌تون نمی‌کنه.", "Not bad for the price, but it won't amaze you.", 3},
	{"اصلاً راضی نبودم", "Not satisfied at all", "رنگ محصول با عکس فرق داشت و کیفیت دوخت پایین بود.", "The product color was different from the picture and stitching quality was low.", 1},
	{"ارسال سریع و عالی", "Fast and great shipping", "بسته‌بندی عالی بود و خیلی سریع به دستم رسید. ممنون.", "The packaging was great and it arrived very fast. Thanks.", 5},
	{"کیفیت ساخت متوسط", "Medium build quality", "برای استفاده روزمره خوبه ولی انتظار کار خیلی خاصی نداشته باشید.", "Good for daily use, but don't expect anything extraordinary.", 4},
	{"بسیار شیک و راحت", "Very stylish and comfortable", "دقیقاً همونی بود که تو عکس دیدم. پام توش کاملاً راحته.", "It was exactly what I saw in the picture. My feet are very comfortable in it.", 5},
	{"ارزش خرید پایین", "Low value for money", "با این قیمت می‌تونید گزینه‌های خیلی بهتری پیدا کنید.", "You can find much better options at this price.", 2},
	{"سایز استاندارد", "Standard size", "سایزش کاملاً اندازه بود و مشکلی نداشت.", "The size fit perfectly and there were no issues.", 4},
	{"پشیمون شدم از خرید", "Regret buying", "بعد از دو هفته استفاده شروع کرد به خراب شدن.", "It started falling apart after two weeks of use.", 1},
	{"طراحی فوق‌العاده", "Gorgeous design", "طراحیش خیلی خاصه و همه ازم پرسیدن از کجا خریدم.", "The design is very unique and everyone asked me where I bought it.", 5},
}

// لیست کاربران واقعی برای ثبت‌نام در جدول اصلی users
var sampleUsers = []struct {
	FullName string
	Email    string
	Password string
}{
	{"علی رضایی", "ali.rezaei@gmail.com", "Password123"},
	{"سارا احمدی", "sara.ahmadi@yahoo.com", "Password123"},
	{"محمد کریمی", "mohammad.karimi@gmail.com", "Password123"},
	{"زهرا مرادی", "zahra.moradi@gmail.com", "Password123"},
	{"رضا تهرانی", "reza.tehrani@gmail.com", "Password123"},
}

func main() {
	dsn := "host=127.0.0.1 port=5432 user=postgres password=secret dbname=nikeStoreDb sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("خطا در اتصال به دیتابیس: %v", err)
	}
	defer db.Close()

	// ۱. وارد کردن آیدی محصولات دلخواه شما (آیدی‌ها را اینجا دستی وارد کنید)
	productIDs := []int{1, 2, 3, 4, 5} // آیدی محصولات خود را اینجا جایگزین کنید

	// ۲. ساخت کاربران واقعی در جدول اصلی users (اگر از قبل نباشند)
	var userIDs []int64
	for _, u := range sampleUsers {
		// بررسی اینکه آیا کاربر از قبل وجود دارد یا خیر
		var existingID int64
		err := db.QueryRow("SELECT id FROM users WHERE email = $1", u.Email).Scan(&existingID)

		if err == sql.ErrNoRows {
			// هش کردن پسورد دقیقاً مطابق استانداردهای احراز هویت
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
			if err != nil {
				log.Printf("خطا در هش کردن پسورد: %v", err)
				continue
			}

			// درج کاربر جدید در جدول اصلی users
			var newID int64
			err = db.QueryRow(
				`INSERT INTO users (email, password_hash, full_name, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id`,
				u.Email, string(hashedPassword), u.FullName,
			).Scan(&newID)

			if err != nil {
				log.Printf("خطا در ثبت کاربر %s: %v", u.Email, err)
				continue
			}
			userIDs = append(userIDs, newID)
		} else {
			userIDs = append(userIDs, existingID)
		}
	}

	if len(userIDs) == 0 {
		log.Fatal("هیچ کاربری برای ثبت کامنت یافت نشد.")
	}

	rand.Seed(time.Now().UnixNano())

	// ۳. درج کامنت برای محصولات مشخص شده
	for _, prodID := range productIDs {
		for c := 0; c < 50; c++ {
			userID := userIDs[rand.Intn(len(userIDs))]
			template := commentTemplates[rand.Intn(len(commentTemplates))]

			titleFa := fmt.Sprintf("%s (%d)", template.TitleFa, c+1)
			titleEn := fmt.Sprintf("%s #%d", template.TitleEn, c+1)

			// تصحیح کوئری درج و تطابق کامل تعداد ستون‌ها با پارامترها ($1 تا $6)
			_, err := db.Exec(`
				INSERT INTO product_comments (product_id, user_id, title_fa,title_en, body_fa, body_en, rating) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`,
				prodID, userID, titleFa, titleEn, template.BodyFa, template.BodyEn, template.Rating,
			)
			if err != nil {
				log.Printf("خطا در ثبت کامنت برای محصول %d: %v", prodID, err)
			}
		}
		fmt.Printf("✅ تعداد ۵۰ کامنت برای محصول با آیدی %d ثبت شد.\n", prodID)
	}

	fmt.Println("🎉 عملیات با موفقیت به پایان رسید!")
}
