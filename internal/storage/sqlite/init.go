package sqlite

import (
	"fmt"
	"log"
	"strings"
)

type colData struct {
	Type string
	PK   bool
}

var dbStructure = map[string]map[string]colData{

	"user": map[string]colData{
		"email":    colData{"text", true},
		"passhash": colData{"text", false},
		"passsalt": colData{"text", false},
	},

	"token": map[string]colData{
		"token":      colData{"text", true},
		"user_email": colData{"text", false},
	},

	"workspace": map[string]colData{
		"id": colData{"text", true},
	},

	"workspace_user": map[string]colData{
		"workspace_id": colData{"text", true},
		"user_email":   colData{"text", true},
	},
}

func (s *Service) checkInit() error {
	for name, structure := range dbStructure {
		if err := s.hasTable(name, structure); err != nil {
			return err
		}
	}
	return nil
}

// hasTable makes sure the database has this table, with the given columns
//
// name is the name of the table
// cols is a mapping from the column name and its type ("integer", "real", "text" or "blob")
func (s *Service) hasTable(name string, cols map[string]colData) error {
	row := s.db.QueryRow("SELECT COUNT(name) FROM sqlite_master WHERE type='table' AND name='" + name + "';")
	var tableCount int
	err := row.Scan(&tableCount)
	if err != nil {
		return err
	}
	if tableCount == 0 {
		var (
			colParts    []string
			primaryKeys []string
		)
		for colName, colData := range cols {
			if colData.PK {
				primaryKeys = append(primaryKeys, colName)
			}
			colParts = append(colParts, colName+" "+colData.Type)
		}
		if len(primaryKeys) > 0 {
			colParts = append(colParts, "PRIMARY KEY ("+strings.Join(primaryKeys, ", ")+")")
		}
		_, err := s.db.Exec("CREATE TABLE " + name + " (" + strings.Join(colParts, ", ") + ");")
		return err
	}
	rows, err := s.db.Query("PRAGMA table_info(" + name + ");")

	if err != nil {
		return err
	}
	for rows.Next() {
		var (
			colCid     int
			colName    string
			colType    string
			colNotnull int
			colDefault interface{}
			colPk      int
		)
		if err := rows.Scan(&colCid, &colName, &colType, &colNotnull, &colDefault, &colPk); err != nil {
			return err
		}
		needed, ok := cols[colName]
		if !ok {
			log.Printf("SQLite: column %s in table %s unneeded", colName, name)
			continue
		}
		colType = strings.ToLower(colType)
		if colType != needed.Type {
			return fmt.Errorf("SQLite: column %s in table %s is of type %s, but %s is needed", colName, name, colType, needed.Type)
		}
		if colPk == 0 && needed.PK {
			return fmt.Errorf("SQLite: column %s in table %s is not a primary key but should be", colName, name)
		}
		if colPk != 0 && !needed.PK {
			return fmt.Errorf("SQLite: column %s in table %s is a primary key but should not be", colName, name)
		}
		delete(cols, colName)
	}
	if err := rows.Close(); err != nil {
		return err
	}
	for colName, colData := range cols {
		var primaryKey = ""
		if colData.PK {
			primaryKey = " PRIMARY KEY"
		}
		if _, err := s.db.Exec("ALTER TABLE " + name + " ADD COLUMN " + colName + " " + colData.Type + primaryKey + ";"); err != nil {
			return err
		}
	}
	return rows.Close()
}
