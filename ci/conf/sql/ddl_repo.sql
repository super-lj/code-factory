CREATE TABLE `code_factory_repo` (
        `id`                bigint  unsigned    NOT NULL   AUTO_INCREMENT COMMENT 'primary key',
        `name`              varchar(256)        NOT NULL   COMMENT 'name',
        `url`               varchar(256)        NOT NULL   COMMENT 'url',
        `status`            int                 NOT NULL   COMMENT 'status',
        `history`           varchar(512)        NOT NULL   COMMENT 'history',
        `update_time`       timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
        `create_time`       timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT 'create time',
        PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;