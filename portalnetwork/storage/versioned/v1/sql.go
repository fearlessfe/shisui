package v1

import "fmt"

type StoreType string

const (
	History StoreType = "history"
	State StoreType = "state"
)

func tableName(typ StoreType) string {
	return "ii1_" + string(typ)
}

func CreateTableSql(typ StoreType) string {
	tableName := tableName(typ);
	createTableSQL := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            content_id BLOB PRIMARY KEY,
            content_key BLOB NOT NULL,
            content_value BLOB NOT NULL,
            distance_short INTEGER NOT NULL,
            content_size INTEGER NOT NULL
        );
        CREATE INDEX IF NOT EXISTS %s_distance_short_idx ON %s (distance_short);
        CREATE INDEX IF NOT EXISTS %s_content_size_idx ON %s (content_size);
    `, tableName, tableName, tableName, tableName, tableName)
		return createTableSQL;
}

func InsertSql(typ StoreType) string {
	tableName := tableName(typ);
	insertSql := fmt.Sprintf(`
	INSERT OR IGNORE INTO %s (
            content_id,
            content_key,
            content_value,
            distance_short,
            content_size
        )
				VALUES (?1, ?2, ?3, ?4, ?5)		
	`, tableName)
	return insertSql
}

func DeleteSql(typ StoreType) string {
	tableName := tableName(typ);
	deleteSql := fmt.Sprintf(`
	DELETE FROM %s
	WHERE content_id = (?1)
        RETURNING content_size		
	`, tableName)
	return deleteSql
}

func LookupKeySql(typ StoreType) string {
	tableName := tableName(typ);
	lookupKeySql := fmt.Sprintf("SELECT content_key FROM %s WHERE content_id = (?1) LIMIT 1", tableName)
	return lookupKeySql
}

func LookupValueSql(typ StoreType) string {
	tableName := tableName(typ);
	lookupValueSql := fmt.Sprintf("SELECT content_value FROM %s WHERE content_id = (?1) LIMIT 1", tableName)
	return lookupValueSql
}

func DeleteFarthestSql(typ StoreType) string {
	tableName := tableName(typ);
	// rowid is sqlite's default id
	deleteFarthestSql := fmt.Sprintf(`
	DELETE FROM %s
	WHERE rowid IN(
		SELECT rowid
    FROM %s
    ORDER BY distance_short DESC
    LIMIT (?1)
	)
	 RETURNING content_size
	`, tableName, tableName)
	return deleteFarthestSql
}

func LookupFarthestSql(typ StoreType) string {
	tableName := tableName(typ);
	// rowid is sqlite's default id
	deleteFarthestSql := fmt.Sprintf(`
	SELECT content_id, distance_short FROM %s
        ORDER BY distance_short DESC
        LIMIT (?1)
	`, tableName)
	return deleteFarthestSql
}

func EntryCountAndSizeSql(typ StoreType) string {
	tableName := tableName(typ)
	sql := fmt.Sprintf("SELECT COUNT(*) as count, TOTAL(content_size) as used_capacity FROM %s", tableName)
	return sql
}

