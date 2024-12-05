CREATE TABLE `users` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT null,
  `username` varchar(255) DEFAULT null,
  `email` varchar(255) DEFAULT null,
  `password` varchar(255) DEFAULT null,
  `created_at` datetime DEFAULT null,
  `deleted_at` datetime DEFAULT null
);

CREATE TABLE `products` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT null,
  `description` varchar(255) DEFAULT null,
  `price` int DEFAULT null,
  `category_id` int DEFAULT null,
  `created_at` datetime DEFAULT null,
  `deleted_at` datetime DEFAULT null
);

CREATE TABLE `categories` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT null,
  `description` text DEFAULT null,
  `created_at` datetime DEFAULT null,
  `deleted_at` datetime DEFAULT null
);

CREATE TABLE `carts` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT null,
  `created_at` datetime DEFAULT null,
  `deleted_at` datetime DEFAULT null
);

CREATE TABLE `cart_item` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `cart_id` int DEFAULT null,
  `product_id` int DEFAULT null,
  `quantity` int DEFAULT null,
  `created_at` datetime DEFAULT null,
  `deleted_at` datetime DEFAULT null
);

CREATE TABLE `orders` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT null,
  `amount` int DEFAULT null,
  `created_at` datetime DEFAULT null,
  `deleted_at` datetime DEFAULT null
);

CREATE TABLE `order_item` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `order_id` int DEFAULT null,
  `product_id` int DEFAULT null,
  `quantity` int DEFAULT null,
  `created_at` datetime DEFAULT null,
  `deleted_at` datetime DEFAULT null
);

CREATE TABLE `payments` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `order_id` int DEFAULT null,
  `amount` int DEFAULT null,
  `provider` varchar(255) DEFAULT null,
  `status` varchar(255) DEFAULT null,
  `created_at` datetime DEFAULT null,
  `deleted_at` datetime DEFAULT null
);

CREATE INDEX `idx_category_id` ON `products` (`category_id`);

CREATE INDEX `idx_deleted_at` ON `products` (`deleted_at`);

CREATE INDEX `idx_user_id` ON `carts` (`user_id`);

CREATE INDEX `idx_deleted_at` ON `carts` (`deleted_at`);

CREATE INDEX `idx_cart_id` ON `cart_item` (`cart_id`);

CREATE INDEX `idx_product_id` ON `cart_item` (`product_id`);

CREATE INDEX `deleted_at` ON `cart_item` (`deleted_at`);

CREATE INDEX `idx_user_id` ON `orders` (`user_id`);

CREATE INDEX `idx_order_id` ON `order_item` (`order_id`);

CREATE INDEX `idx_product_id` ON `order_item` (`product_id`);

CREATE INDEX `deleted_at` ON `order_item` (`deleted_at`);

ALTER TABLE `products` ADD FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);

ALTER TABLE `cart_item` ADD FOREIGN KEY (`cart_id`) REFERENCES `carts` (`id`);

ALTER TABLE `cart_item` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE `order_item` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

ALTER TABLE `order_item` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE `orders` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `payments` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);
