package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// اتصال به دیتابیس (حتما از 127.0.0.1 استفاده کن تا مشکل IPv6 پیش نیاید)
	connStr := "host=127.0.0.1 port=5432 user=postgres password=secret dbname=nikeStoreDb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	fmt.Println("Seeding started...")

	seedCategories(db)
	seedProducts(db)
	seedVariants(db)

	fmt.Println("Seeding completed successfully! 100 Nike shoes added. 🚀")
}

func seedCategories(db *sql.DB) {
	categories := []struct {
		slug, nameEn, nameFa string
	}{
		{"running", "Running", "رانینگ (دویدن)"},
		{"lifestyle", "Lifestyle", "روزمره (لایف‌استایل)"},
		{"basketball", "Basketball", "بسکتبال"},
	}

	for _, c := range categories {
		// با ON CONFLICT از ساخت دسته‌بندی تکراری در اجراهای مجدد جلوگیری می‌کنیم
		query := `INSERT INTO categories (slug, name_en, name_fa) VALUES ($1, $2, $3) ON CONFLICT (slug) DO NOTHING`
		_, err := db.Exec(query, c.slug, c.nameEn, c.nameFa)
		if err != nil {
			log.Printf("Error inserting category %s: %v\n", c.slug, err)
		}
	}
	fmt.Println("Categories seeded.")
}

func seedProducts(db *sql.DB) {
	// مدل‌های پایه برای تولید اسم‌های تصادفی
	baseModelsEn := []string{"Air Max 270", "Air Force 1", "Pegasus 40", "ZoomX Invincible", "Air Jordan 1 Retro", "Dunk Low"}
	baseModelsFa := []string{"ایر مکس ۲۷۰", "ایر فورس ۱", "پگاسوس ۴۰", "زوم ایکس اینوینسیبل", "ایر جردن ۱ رترو", "دانک لو"}

	colorsEn := []string{"Black/White", "Red/Black", "Triple White", "Blue/Grey", "Neon"}
	colorsFa := []string{"مشکی/سفید", "قرمز/مشکی", "تمام سفید", "آبی/خاکستری", "نئون"}

	rand.Seed(time.Now().UnixNano())

	// گرفتن آیدی دسته‌بندی‌ها از دیتابیس
	rows, err := db.Query("SELECT id FROM categories")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var categoryIDs []int
	for rows.Next() {
		var id int
		rows.Scan(&id)
		categoryIDs = append(categoryIDs, id)
	}

	// ساخت ۱۰۰ محصول
	for i := 1; i <= 100; i++ {
		catID := categoryIDs[rand.Intn(len(categoryIDs))] // انتخاب رندوم یک دسته‌بندی
		modelIndex := rand.Intn(len(baseModelsEn))
		colorIndex := rand.Intn(len(colorsEn))

		// تولید اسم محصول
		titleEn := fmt.Sprintf("Nike %s %s", baseModelsEn[modelIndex], colorsEn[colorIndex])
		titleFa := fmt.Sprintf("کفش نایک %s رنگ %s", baseModelsFa[modelIndex], colorsFa[colorIndex])

		descEn := fmt.Sprintf("Experience the ultimate comfort with %s. Perfect for your daily activities.", titleEn)
		descFa := fmt.Sprintf("نهایت راحتی را با %s تجربه کنید. انتخابی بی‌نظیر برای فعالیت‌های روزانه و استایل شما.", titleFa)

		// تولید قیمت رندوم (بین 80 تا 250 دلار)
		priceUSD := 80.0 + rand.Float64()*(250.0-80.0)

		// تبدیل حدودی به تومان (فرض: هر دلار 60,000 تومان) + کمی رندومایز برای طبیعی شدن قیمت
		priceToman := int64(priceUSD * 60000)
		// رند کردن به نزدیکترین هزار تومان (مثلا 5,234,000)
		priceToman = (priceToman / 1000) * 1000

		// چند عکس فیک
		images := []string{
			fmt.Sprintf("https://dummyimage.com/600x400/eeeeee/333333&text=Nike+%d-1", i),
			fmt.Sprintf("https://dummyimage.com/600x400/dddddd/333333&text=Nike+%d-2", i),
		}

		// برای Postgres آرایه‌ها باید به این فرمت استرینگ تبدیل بشن: {"url1","url2"}
		pgArray := fmt.Sprintf(`{"%s","%s"}`, images[0], images[1])

		query := `INSERT INTO products (category_id, title_en, title_fa, description_en, description_fa, price_usd, price_toman, images) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

		_, err := db.Exec(query, catID, titleEn, titleFa, descEn, descFa, priceUSD, priceToman, pgArray)
		if err != nil {
			log.Printf("Error inserting product %d: %v\n", i, err)
		}
	}
}

func seedVariants(db *sql.DB) {
	fmt.Println("Seeding product variants (colors & sizes)...")

	// گرفتن آیدی تمام محصولات
	rows, err := db.Query("SELECT id FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var productIDs []int
	for rows.Next() {
		var id int
		rows.Scan(&id)
		productIDs = append(productIDs, id)
	}

	colorsEn := []string{"Black", "White", "Red", "Blue"}
	colorsFa := []string{"مشکی", "سفید", "قرمز", "آبی"}
	sizes := []string{"40", "41", "42", "42.5", "43", "44"}

	rand.Seed(time.Now().UnixNano())

	// برای هر محصول، چند تا واریانت رندوم می‌سازیم
	for _, pID := range productIDs {
		// هر کفش بین ۲ تا ۴ رنگ/سایز مختلف داره
		numVariants := rand.Intn(3) + 2

		// برای جلوگیری از ساخت واریانت تکراری برای یک محصول
		createdVariants := make(map[string]bool)

		for i := 0; i < numVariants; i++ {
			cIndex := rand.Intn(len(colorsEn))
			sIndex := rand.Intn(len(sizes))

			// کلید یونیک برای چک کردن تکراری نبودن در مموری
			variantKey := fmt.Sprintf("%d-%s", cIndex, sIndex)
			if createdVariants[variantKey] {
				continue
			}
			createdVariants[variantKey] = true

			stock := rand.Intn(20) // موجودی رندوم بین 0 تا 19

			query := `INSERT INTO product_variants (product_id, color_en, color_fa, size, stock) VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`

			_, err := db.Exec(query, pID, colorsEn[cIndex], colorsFa[cIndex], sizes[sIndex], stock)
			if err != nil {
				log.Printf("Error inserting variant for product %d: %v\n", pID, err)
			}
		}
	}
	fmt.Println("Variants seeded.")
}
