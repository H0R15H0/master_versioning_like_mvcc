CREATE TABLE IF NOT EXISTS `items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_id` int(11) NOT NULL,
  `code` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `index_group_id` (`group_id`),
  UNIQUE KEY `group_id_code_created_at` (`group_id`, `code`, `created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
