package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"log"
	"text/template"
	"strconv"
)

// بنية (Struct) لبيانات الكتب
type Book struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	PublishedYear int    `json:"published_year"`
	Description  string `json:"description"`
}

type Library struct {
	Books []Book `json:"books"`
}

// تحميل البيانات من ملف JSON
func loadBooks() ([]Book, error) {
	file, err := os.Open("books.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var library Library
	err = json.NewDecoder(file).Decode(&library)
	if err != nil {
		return nil, err
	}

	return library.Books, nil
}

// عرض قائمة الكتب
func booksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := loadBooks()
	if err != nil {
		http.Error(w, "Unable to load books", http.StatusInternalServerError)
		return
	}

	tmpl , err := template.ParseFiles("books.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, books); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// عرض تفاصيل كتاب معين
func bookDetailsHandler(w http.ResponseWriter, r *http.Request) {
	books, err := loadBooks()
	if err != nil {
		http.Error(w, "Unable to load books", http.StatusInternalServerError)
		return
	}

	// الحصول على معرف الكتاب من الرابط
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 || id > len(books) {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// البحث عن الكتاب المطلوب
	var selectedBook Book
	for _, book := range books {
		if book.ID == id {
			selectedBook = book
			break
		}
	}

	tmpl := template.Must(template.ParseFiles("book_details.html"))
	if err := tmpl.Execute(w, selectedBook); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func main() {
	// إعداد المسارات
	http.HandleFunc("/", booksHandler)
	http.HandleFunc("/book", bookDetailsHandler)

	// تشغيل الخادم على المنفذ 8080
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
