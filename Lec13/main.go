package main

import "fmt"

func main() {

	// константы - это неизменяемые переменные, могут быть только примитивные типы
	// значения констант должны быть известны на момент компиляции
	const a = 10
	fmt.Println(a)

	// блок констант
	const (
		ipAddress = "127.127.0.3"
		port      = "8000"
		dbname    = "postgres"
	)

	// типизированные и нетипизированные константы
	const ADMIN_EMAIL string = "admin@admin.com"
}
