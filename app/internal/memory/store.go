package memory

import (
	"database/sql"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
}

func Open(sqlitePath string) (*Store, error) {
	// Ensure directory exists (./agent.db なら "." になるのでOK)
	if dir := filepath.Dir(sqlitePath); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("sqlite", sqlitePath)
	if err != nil {
		return nil, err
	}

	s := &Store{db: db}
	if err := s.migrate(); err != nil {
		_ = db.Close()
		return nil, err
	}
	return s, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) migrate() error {
	_, err := s.db.Exec(`
CREATE TABLE IF NOT EXISTS messages (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  session_id TEXT NOT NULL,
  role TEXT NOT NULL,
  content TEXT NOT NULL,
  created_at TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_messages_session ON messages(session_id);
`)
	return err
}

func (s *Store) AddMessage(sessionID, role, content string) error {
	_, err := s.db.Exec(
		`INSERT INTO messages(session_id, role, content, created_at) VALUES (?, ?, ?, ?)`,
		sessionID, role, content, time.Now().UTC().Format(time.RFC3339),
	)
	return err
}
