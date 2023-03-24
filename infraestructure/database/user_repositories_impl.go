package database

import (
	"database/sql"
	"hexagonal02/domain/entities"
	"log"
)	

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) Save(user *entities.User) error {
 	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"

    res , err := r.db.Exec(query, user.Name, user.Email, user.Password)
    if err != nil {
        return err
    }
	id, _ := res.LastInsertId()
	log.Printf("ID guardado: %v", id )
    return nil
}

func (r *UserRepositoryImpl) Update(user *entities.User) error {
   
   query := "UPDATE users SET name = ?, email = ?, password= ?  WHERE id = ? "

   _ , err := r.db.Exec(query, user.Name, user.Email, user.Password, user.ID )
   if err != nil {
	   return err
   }   
   
   return nil
}