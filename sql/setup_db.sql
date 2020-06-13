ALTER DATABASE db SET timezone TO 'UTC';

/* 
Revoke access and create for everyone on schema `public`.
Grant `USAGE` and `CREATE` for admin user `db_admin`.
Grant `USAGE` for `db_user` and `db_readonly`.
*/
REVOKE USAGE ON SCHEMA public FROM PUBLIC;
REVOKE CREATE ON SCHEMA public FROM PUBLIC;

GRANT USAGE ON SCHEMA public to db_admin;
GRANT CREATE ON SCHEMA public to db_admin;

GRANT USAGE ON SCHEMA public to db_user;
GRANT USAGE ON SCHEMA public to db_readonly;
