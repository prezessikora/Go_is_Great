package notes

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) Note {
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
}

func (note Note) Display() {
	fmt.Printf("Title: %v , Content: %v\n", note.Title, note.Content)
}

func (n Note) Save() error {
	fileName := strings.ToLower(n.Title) + ".json"
	fileName = strings.ReplaceAll(fileName, " ", "_")
	jsonFile, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonFile, 0644)
}
