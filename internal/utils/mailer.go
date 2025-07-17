package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
)

func SendMail(to []string, subject, templatePath string, data any) error {
	// Load SMTP configuration from environment variables
	SMTP_HOST := os.Getenv("SMTP_HOST")
	SMTP_PORT := os.Getenv("SMTP_PORT")
	SMTP_USER := os.Getenv("SMTP_USER")
	// SMTP_PASSWORD := os.Getenv("SMTP_PASSWORD")
	if SMTP_HOST == "" || SMTP_PORT == "" {
		return fmt.Errorf("SMTP_HOST and SMTP_PORT must be set in environment variables")
	}
	path := filepath.Join(GetCurrentDir(), "..", "internal", "templates", templatePath)
	// Parsing file template
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Printf("Gagal parsing template: %v", err)
		return fmt.Errorf("gagal parsing template: %w", err)
	}

	var body bytes.Buffer

	// Header email
	body.WriteString(fmt.Sprintf("To: %s\r\n", to[0]))
	body.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	body.WriteString("MIME-Version: 1.0\r\n")
	body.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n")

	// Tulis isi template ke body
	err = tmpl.Execute(&body, data)
	if err != nil {
		return fmt.Errorf("gagal eksekusi template: %w", err)
	}

	addr := fmt.Sprintf("%s:%s", SMTP_HOST, SMTP_PORT)

	err = smtp.SendMail(addr, nil, SMTP_USER, to, body.Bytes())
	if err != nil {
		log.Printf("Gagal mengirim email: %v", err)
		return err
	}

	log.Println("Email berhasil dikirim")
	return nil
}

func GetCurrentDir() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Gagal dapatkan path: %v", err)
	}

	return filepath.Dir(exePath)
}
