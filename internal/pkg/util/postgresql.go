package util

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgreSQL struct {
	pool *pgxpool.Pool
}

func init() {
	log.Println("pgsl init().")
}

type Name struct {
	NConst    string `json:"nconst"`
	Name      string `json:"name"`
	BirthYear string `json:"birthYear"`
	DeathYear string `json:"deathYear"`
}

// NewPostgreSQL returns the psgl active connection.
// Ref for me: return struct instead of just connection, so we can make functions. (*pgxpool.Pool, error) to (struct, err)(*pgxpool.Pool, error).
func NewPostgreSQL() (*PostgreSQL, error) {
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &PostgreSQL{
		pool: conn,
	}, err
}

// Close closes the psgl connection.
// Ref for me: Used pointer here as the p value should be consistent across called place.
func (p *PostgreSQL) Close() {
	p.pool.Close()
}

func (p *PostgreSQL) FindByValue(nconst string) (Name, error) {
	query := `SELECT nconst, primary_name, birth_year, death_year FROM "names" WHERE nconst = $1`

	var res Name

	if err := p.pool.QueryRow(context.Background(), query, nconst).
		Scan(&res.NConst, &res.Name, &res.BirthYear, &res.DeathYear); err != nil {
		return Name{}, err
	}

	return res, nil
}

/*
 nconst   	 | primary_name | birth_year | death_year | primary_professions              | known_for_titles
------------------------------------------------------------------------------------------------------------------------------------
 nm0000001  | Fred Astaire | 1899       | 1987       | {soundtrack,actor,miscellaneous} | {tt0050419,tt0031983,tt0072308,tt0053137}
*/
