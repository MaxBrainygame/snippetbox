package mysql

import (
	"database/sql"
	"errors"

	"golangify.com/snippetbox/pkg/models"
)

// SnippetModel - Определяем тип который обертывает пул подключения sql.DB
type SnippetModel struct {
	DB *sql.DB
}

// Insert - Метод для создания новой заметки в базе дынных.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {

	stmt := `INSERT INTO snippets (title, content, created, expires)
   		VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Используем метод LastInsertId(), чтобы получить последний ID
	// созданной записи из таблицу snippets.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {

	smtp := `SELECT * FROM snippets
		WHERE expires > UTC_TIMESTAMP() AND id = ?`

	sqlRow := m.DB.QueryRow(smtp, id)

	tableSnippet := &models.Snippet{}

	err := sqlRow.Scan(&tableSnippet.ID, &tableSnippet.Title, &tableSnippet.Content, &tableSnippet.Created,
		&tableSnippet.Expires)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}

	}

	return tableSnippet, nil
}

// Latest - Метод возвращает 10 наиболее часто используемые заметки.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {

	smtp := `SELECT * FROM snippets
		WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	sqlRows, err := m.DB.Query(smtp)
	if err != nil {
		return nil, err
	}
	defer sqlRows.Close()

	var tablesSnippet []*models.Snippet

	for sqlRows.Next() {

		tableSnippet := &models.Snippet{}
		err = sqlRows.Scan(&tableSnippet.ID, &tableSnippet.Title, &tableSnippet.Content, &tableSnippet.Created,
			&tableSnippet.Expires)
		if err != nil {
			return nil, err
		}

		tablesSnippet = append(tablesSnippet, tableSnippet)

	}

	return tablesSnippet, nil

}
