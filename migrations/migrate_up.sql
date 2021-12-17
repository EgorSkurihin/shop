CREATE TABLE `users` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255) NOT NULL UNIQUE,
	`password` VARCHAR(255) NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `albums` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(255) NOT NULL, 
    `performer` VARCHAR(255) NOT NULL,
    `cost` INT NOT NULL,
    `image` VARCHAR(255) NOT NULL,
	PRIMARY KEY (`id`)
);


INSERT INTO `users` (`name`, `password`) VALUES ('Egor', '$2a$10$0UWpDXKCrmtrUAEXWkczk.hJdHusoMlaZAu8wvbNenU/mR3kF9fsy');
/* password: Admin */

INSERT INTO `albums` (`title`, `performer`, `cost`, `image`) VALUES 
("Abbey Road", "The Beatles", 30, "store/albums_img/abbey-road.jpg"),
("Real Gone", "Tom Waits", 25, "store/albums_img/real-gone.jpg"),
("Led Zeppelin 4", "Led Zeppelin", 27, "store/albums_img/led-zeppelin-4.jpg"),
("Nevermind", "Nirvana", 30, "store/albums_img/nevermind.jpg"),
("Tattoo You", "The Rolling Stones", 36, "store/albums_img/tatooyou.jpeg"),
("The Dark Side Of The Moon", "Pink Floyd", 30, "store/albums_img/the-dark-side-of-the-moon.jfif"),
("The Velvet Underground and Nico","The Velvet Underground", 35, "store/albums_img/velvet-underground-and-nico.jpg"),
("XTC", "Drums And Wires", 31, "store/albums_img/drums-and-wires.jpg");
