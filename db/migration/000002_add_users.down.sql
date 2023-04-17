ALTER TABLE accounts DROP FOREIGN KEY accounts_ibfk_1;
DROP INDEX `user_index` ON accounts;
DROP TABLE IF EXISTS users;