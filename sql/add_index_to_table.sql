-- Add index to items table on (code, created_at)
ALTER TABLE `items` ADD INDEX `index_code_created_at` (`code`, `created_at` desc);
