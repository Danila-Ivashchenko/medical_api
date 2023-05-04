CREATE TABLE IF NOT EXISTS relatives (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` BIGINT NOT NULL,
    `phone` VARCHAR(11) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `surname` VARCHAR(255) NOT NULL,
    `patronymic` VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `users`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `surname` VARCHAR(255) NOT NULL,
    `patronymic` VARCHAR(255) NOT NULL,
    `polis` VARCHAR(16) NOT NULL,
    `phone` VARCHAR(11) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `city` VARCHAR(255) NOT NULL,
    `address` VARCHAR(255) NOT NULL,
    `birthday` DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS `diseases`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` BIGINT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `ill_date` DATE NOT NULL,
    `record_date` DATE NOT NULL,
    `status` TINYINT(1) NOT NULL DEFAULT '1' COMMENT 'Статус заболевания
true - болеет
false - не болеет'
);

CREATE TABLE IF NOT EXISTS `hospital`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `city` VARCHAR(255) NOT NULL,
    `address` VARCHAR(255) NOT NULL,
    `contact_phone` VARCHAR(11) NOT NULL,
    `cantact_email` BIGINT NOT NULL,
    `lon` DECIMAL(10, 10) NOT NULL,
    `lat` DECIMAL(10, 10) NOT NULL
);

CREATE TABLE IF NOT EXISTS `attachment`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` BIGINT NOT NULL,
    `hospital_id` BIGINT NOT NULL,
    `date` BIGINT NOT NULL
)
