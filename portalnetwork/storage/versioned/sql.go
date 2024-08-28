package versioned

const StoreInfoCreateTable = `
	CREATE TABLE IF NOT EXISTS store_info (
        content_type TEXT PRIMARY KEY,
        version TEXT NOT NULL
    )
`

const StoreInfoUpdate = `
		INSERT OR REPLACE INTO store_info (content_type, version)
    VALUES (?1, ?2)
`

const StoreInfoLookup = `
	  SELECT version
    FROM store_info
    WHERE content_type = (?1)
    LIMIT 1";
`

const TableExist = `
		SELECT name
    FROM sqlite_master
    WHERE type='table' AND name= (?1)"
`