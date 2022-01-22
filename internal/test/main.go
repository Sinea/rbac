package main

import (
	"context"
	"log"

	"sinea.xyz/rbac/v1/domain"
	"sinea.xyz/rbac/v1/gateways/postgresql"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	url := "postgresql://sin:pass@0.0.0.0/rbac"
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		panic(err)
	}
	cfg.MinConns = 10
	cfg.MaxConns = 100
	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		panic(err)
	}

	userID := "cain"

	usersRepo := postgresql.NewUsersRepository(pool, "users", "user_roles")
	rolesRepo := postgresql.NewRolesRepository(pool, "roles", "role_permissions")
	_, err = usersRepo.Create(context.Background(), userID)
	if err != nil {
		log.Println(err)
	}
	_, err = rolesRepo.CreateRole(context.Background(), "sudoers")
	if err != nil {
		log.Println(err)
	}
	err = rolesRepo.GrantPermissions(context.Background(), "sudoers",
		domain.NewPermission("bucket", "write"),
		domain.NewPermission("bucket", "read"),
		domain.NewPermission("bucket", "peek"))
	if err != nil {
		log.Println(err)
	}
	err = usersRepo.AssignRole(context.Background(), userID, "sudoers")

	if err != nil {
		log.Println(err)
	}
}
