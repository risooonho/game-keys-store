package mysqlStore

import (
	"database/sql"
	"github.com/rdsalakhov/game-keys-store/internal/model"
	"github.com/rdsalakhov/game-keys-store/internal/store"
)

type GameRepository struct {
	store *Store
}

func (repo *GameRepository) Find(ID int) (*model.Game, error) {
	selectQuery := "SELECT id, title, description, price, on_sale FROM games WHERE id = ?"
	game := &model.Game{}
	if err := repo.store.db.QueryRow(selectQuery, ID).Scan(
		&game.ID,
		&game.Title,
		&game.Description,
		&game.Price,
		&game.OnSale); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return game, nil
}

func (repo *GameRepository) Create(game *model.Game) error {
	insertQuery := "INSERT INTO games (title, description, price, on_sale) VALUES (?, ?, ?, ?);"
	getIdQuery := "select LAST_INSERT_ID();"

	if _, err := repo.store.db.Exec(insertQuery,
		game.Title,
		game.Description,
		game.Price,
		game.OnSale,
	); err != nil {
		return err
	}

	if err := repo.store.db.QueryRow(getIdQuery).Scan(&game.ID); err != nil {
		return err
	}
	return nil
}