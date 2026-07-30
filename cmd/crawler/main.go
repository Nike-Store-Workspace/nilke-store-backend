package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lib/pq" // برای مدیریت آرایه‌ها در پست‌گرس الزامی است
)

func main() {
	// ۱. اتصال به دیتابیس
	dsn := "host=127.0.0.1 user=postgres password=secret dbname=nikeStoreDb port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("خطا در اتصال به دیتابیس: %v", err)
	}
	defer db.Close()

	// مسیر پوشه عکس‌ها (مسیر دقیق بر اساس سیستم شما)
	baseDir := `E:/Go Projects/Nike_Strore_Api/assets/images`

	folders, err := os.ReadDir(baseDir)
	if err != nil {
		log.Fatalf("خطا در خواندن پوشه عکس‌ها: %v", err)
	}

	// ۲. حلقه روی پوشه‌ها (هر پوشه نامش معادل آیدی محصول است)
	for _, folder := range folders {
		if !folder.IsDir() {
			continue
		}

		productID, err := strconv.Atoi(folder.Name())
		if err != nil {
			continue
		}

		productDirPath := filepath.Join(baseDir, folder.Name())
		files, err := os.ReadDir(productDirPath)
		if err != nil {
			log.Printf("خطا در خواندن محتوای پوشه %d: %v", productID, err)
			continue
		}

		var imagePaths []string

		// ۳. جمع‌آوری تمام عکس‌های درون پوشه
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
				// ساخت مسیر نسبی عکس
				imagePath := fmt.Sprintf("/images/%d/%s", productID, file.Name())
				imagePaths = append(imagePaths, imagePath)
			}
		}

		if len(imagePaths) == 0 {
			continue
		}

		// ۴. آپدیت ستون images در جدول products با استفاده از پکیج pq برای تبدیل آرایه
		query := `UPDATE products SET images = $1 WHERE id = $2`
		_, err = db.Exec(query, pq.Array(imagePaths), productID)
		if err != nil {
			log.Printf("خطا در آپدیت عکس‌های محصول %d: %v", productID, err)
		} else {
			fmt.Printf("✅ تعداد %d عکس برای محصول %d با موفقیت ثبت شد.\n", len(imagePaths), productID)
		}
	}

	fmt.Println("🎉 همگام‌سازی و آپدیت عکس‌ها به پایان رسید!")
}
