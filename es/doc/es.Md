下载es：

    https://www.elastic.co/cn/downloads/elasticsearch
    https://www.elastic.co/cn/downloads/kibana
    kibana和elasticsearch版本要一致

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
![img_1.png](img_1.png)

重启后得到下图

![img.png](img.png)

默认域名并访问测试

    es:http://127.0.0.1:9200/
    插件：http://127.0.0.1:9100/
    kibana：http://127.0.0.1:5601/app/home#/

汉化kibana：
    
    将文件D:\tools\kibana-7.16.2-windows-x86_64\config的yml文件的配置修改并重启
    此配置对应的是
    D:\tools\kibana-7.16.2-windows-x86_64\x-pack\plugins\translations\translations下的文件

![img_2.png](img_2.png)

![img_3.png](img_3.png)


整个启动流程
    
    ES：
![img_32.png](img_32.png)
    插件：
![img_30.png](img_30.png)

    kibana：
![img_31.png](img_31.png)


默认集群名称是

![img_5.png](img_5.png)

名称解释与DB对比

    索引（index)   相当于数据库
    文档（document)相当于一条条数据
    类型（types)   相当于数据类型

    倒排索引：统计了数据在文档中出现频率进行比对，过滤掉不相关数据 提高过滤效率

    索引
    字段类型
    文档
    分片（documents，Lucene的倒排索引）

IK分词器
    
    https://github.com/medcl/elasticsearch-analysis-ik/releases 下载对应版本
    下载完毕 放入elasticsearch中，然后重启即可
![img_6.png](img_6.png)

![img_7.png](img_7.png)

    ik_smart：最少切分（最细粒度，不可能有重复数据）
![img_9.png](img_9.png)

    ik_max_word:穷尽词库的可能（从字典中去获取）
![img_8.png](img_8.png)

ik分词器增加自己的配置（解决我们需要的名词）
    ![img_10.png](img_10.png)
    

基本操作
    
    创建索引
    PUT /索引名/~类型名~/文档ID
    {请求体}
![img_12.png](img_12.png)
    
    基本数据类型（看官网）
    
    创建具体索引规则
![img_11.png](img_11.png)
    
    如果文档自动没有指定，es会给我们配置默认的字段类型
![img_13.png](img_13.png)

    通过GET _cat命令可以获取es的当前很多信息

    修改方法
![img_14.png](img_14.png)
![img_15.png](img_15.png)

    删除
![img_16.png](img_16.png)
![img_17.png](img_17.png)


    关于文档的操作
    
    基本操作
![img_18.png](img_18.png)
    
![img_19.png](img_19.png)

    复杂操作
    分布式全文搜索引擎
    hits：索引和文档的信息，查询的结果总数，然后查询出来的数据都能遍历出来；
          我们可以通过score来判断谁更加符合结果
![img_20.png](img_20.png)

![img_21.png](img_21.png)
    must  must_not  should
![img_22.png](img_22.png)
    filter 添加过滤条件
![img_23.png](img_23.png)
    多条件根据空格分开（可以通过分值进行基本判断
![img_24.png](img_24.png)

    精确查询term
    关于分词：
    term：是直接查询精确的
    match：会使用分词器解析！（先分析文档，然后通过分析的文档进行查询）

两个类型

    text：会被分词器解析
    keyword：不会被分词器解析

![img_26.png](img_26.png)
![img_27.png](img_27.png)

高亮查询（与自定义标签）

![img_28.png](img_28.png)
![img_29.png](img_29.png)

golang的 api文档

    https://www.elastic.co/guide/en/elasticsearch/client/index.html

ES查看机器情况（windows)

    http://127.0.0.1:9201/_cat/nodes?v