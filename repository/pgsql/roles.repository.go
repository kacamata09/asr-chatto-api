package repositoryPgSQL

import (
	"database/sql"
	"fmt"
	"go-clean-architecture-by-ahr/domain"
	"time"
)

type repoRoles struct{
    DB *sql.DB
}

func CreateRepoRoles(db *sql.DB) domain.RolesRepository {
    return &repoRoles{DB: db}
}

func (repo *repoRoles) GetAll() ([]domain.Roles, error) {

	rows, err := repo.DB.Query("SELECT * FROM roles")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var data []domain.Roles
    
    for rows.Next() {
        var roles domain.Roles
        err := rows.Scan(&roles.ID, &roles.RoleName, &roles.RoleDesc, &roles.AccessPermission, &roles.CreatedAt, &roles.UpdatedAt)
        if err != nil {
            panic("Failed to scan row: " + err.Error())
        }
        // roles.CreatedAt = time.Now().Add(24 * time.Hour)
        // roles.UpdatedAt = time.Now().Add(24 * time.Hour)
        data = append(data, roles)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }
    return data, err
}

func (repo *repoRoles) GetByID(id string) (domain.Roles, error) {
    row := repo.DB.QueryRow("SELECT * FROM roles where id=$1", id)
    fmt.Println(id)
    
    var data domain.Roles
    
    err := row.Scan(&data.ID, &data.RoleName, &data.RoleDesc, &data.AccessPermission, &data.CreatedAt, &data.UpdatedAt)
    if err != nil {
        panic("Failed to scan row: " + err.Error())
    }
    data.CreatedAt = time.Now().Add(24 * time.Hour)
    data.UpdatedAt = time.Now().Add(24 * time.Hour)
    
    if err := row.Err(); err != nil {
        return domain.Roles{}, err
    }
    // fmt.Println(data)
    return data, err
}