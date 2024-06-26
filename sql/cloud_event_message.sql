create table message_center.cloud_event_message
(
    id                bigserial
        constraint message_cloud_event_pk
            primary key,
    source            text,
    type              text,
    data_content_type text,
    data_schema       text,
    spec_version      text,
    time              timestamp,
    created_at        timestamp,
    updated_at        timestamp,
    event_id          text,
    data_json         json,
    "user"            text,
    source_url        text,
    title             text,
    summary           text,
    constraint message_source_eventid_pk
        unique (source, event_id)
);

comment on table message_center.cloud_event_message is '清洗后的标准消息';

comment on column message_center.cloud_event_message.source is '消息源';

comment on column message_center.cloud_event_message.type is '消息类型';

comment on column message_center.cloud_event_message.data_content_type is '数据内容类型';

comment on column message_center.cloud_event_message.data_schema is '数据分类';

comment on column message_center.cloud_event_message.spec_version is '版本';

comment on column message_center.cloud_event_message.time is '事件时间';

comment on column message_center.cloud_event_message.event_id is '事件id';

comment on column message_center.cloud_event_message.title is '标题';

comment on column message_center.cloud_event_message.summary is '摘要';
