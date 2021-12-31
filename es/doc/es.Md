下载es：

    https://www.elastic.co/cn/downloads/elasticsearch
    https://www.elastic.co/cn/downloads/kibana

下载es插件：
    数据展示工具
    https://github.com/mobz/elasticsearch-head
    
启动插件

    npm i
    npm start

配置跨域

    在D:\tools\elasticsearch-7.16.2\config中修改elasticsearch.yml
    中添加
    http.cors.enabled: true
    http.cors.allow-origin: "*"
重启后得到下图
![img.png](img.png)



