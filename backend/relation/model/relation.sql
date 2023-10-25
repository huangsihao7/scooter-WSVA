DROP TABLE IF EXISTS `relations`;
CREATE TABLE `relations`  (
                              `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
                              `follower_id` int UNSIGNED NOT NULL COMMENT '关注者id',
                              `following_id` int UNSIGNED NOT NULL COMMENT '被关注者id',
                              `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` datetime NULL DEFAULT NULL,
                              `deleted_at` datetime NULL DEFAULT NULL,
                              PRIMARY KEY (`id`) USING BTREE,
                              INDEX `follower_id`(`follower_id` ASC) USING BTREE,
                              INDEX `following_id`(`following_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;