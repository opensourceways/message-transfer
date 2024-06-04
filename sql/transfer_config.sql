create table message_center.transfer_config
(
    id           serial
        constraint transfer_config_pk
            primary key,
    source_topic text,
    field        text,
    template     text not null,
    created_at   timestamp,
    updated_at   timestamp,
    is_deleted   boolean
);

comment on column message_center.transfer_config.source_topic is '消息源';
