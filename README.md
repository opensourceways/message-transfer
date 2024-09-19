# 消息中心清洗服务

## 功能：

1.接受message-collect服务采集到的消息，读取清洗映射配置，将上游消息转换为标准cloudevents格式的数据

[spec/cloudevents/languages/zh-CN/primer.md at main · cloudevents/spec (github.com)](https://github.com/cloudevents/spec/blob/main/cloudevents/languages/zh-CN/primer.md)



2.将清洗后的数据入库，并发送到下游kafka队列