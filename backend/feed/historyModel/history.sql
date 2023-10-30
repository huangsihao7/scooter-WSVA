CREATE TABLE `history`  (
                              `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                              `uid` int UNSIGNED NOT NULL,
                              `vid` int NOT NULL,
                              `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;