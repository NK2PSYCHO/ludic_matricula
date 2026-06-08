package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/uptrace/bun/driver/pgdriver"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
)

type UserSeed struct {
	Name     string `yaml:"name"`
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
	Role     string `yaml:"role"`
}

type UserFile struct {
	Users []UserSeed `yaml:"users"`
}

type GameSeed struct {
	Name      string `yaml:"name"`
	Publisher string `yaml:"publisher"`
	Year      string `yaml:"year"`
	Platform  string `yaml:"platform"`
}

type GameFile struct {
	Games []GameSeed `yaml:"games"`
}

func main() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("pg", dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("ping db: %v", err)
	}

	seedUsers(db)
	seedGames(db)

	log.Println("seeding complete")
}

func seedUsers(db *sql.DB) {
	data, err := os.ReadFile("seed/users.yaml")
	if err != nil {
		log.Fatalf("read users.yaml: %v", err)
	}

	var f UserFile
	if err := yaml.Unmarshal(data, &f); err != nil {
		log.Fatalf("parse users.yaml: %v", err)
	}

	for _, u := range f.Users {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("hash password: %v", err)
		}
		_, err = db.Exec(
			`INSERT INTO users (name, email, password_hash, role) VALUES ($1, $2, $3, $4) ON CONFLICT (email) DO NOTHING`,
			u.Name, u.Email, string(hash), u.Role,
		)
		if err != nil {
			log.Fatalf("insert user %s: %v", u.Email, err)
		}
	}
}

func seedGames(db *sql.DB) {
	data, err := os.ReadFile("seed/games.yaml")
	if err != nil {
		log.Fatalf("read games.yaml: %v", err)
	}

	var f GameFile
	if err := yaml.Unmarshal(data, &f); err != nil {
		log.Fatalf("parse games.yaml: %v", err)
	}

	for _, g := range f.Games {
		_, err = db.Exec(
			`INSERT INTO games (name, publisher, year, platform) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING`,
			g.Name, g.Publisher, g.Year, g.Platform,
		)
		if err != nil {
			log.Fatalf("insert game %s: %v", g.Name, err)
		}
	}
}
