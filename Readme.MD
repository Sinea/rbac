Actions:
- create user
- delete user
- create role
- delete role
- grant permision to role
- revoke permision from role
- add user to role
- remove user from role
- ask if user has permission(s)


permission-userID

user -> role -> permission
user -> permission?

user_permissions table:
permission, userID

user_roles table:
roleName, userID

role_permissions table:
roleName, permission

delete role 
    => remove permissions from all users who have this role

delete user
    => remove user from all roles
    => remove permissions from user

delete permission
    => remove permission from roles (and maybe users?)


create table users
(
    user_id text not null
        constraint users_pk
            primary key
);

create unique index users_user_id_uindex
    on users (user_id);

