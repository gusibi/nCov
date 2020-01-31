## 2020 新型肺炎疫情实时动态播报rss

### 数据来源

1. https://3g.dxy.cn/newh5/view/pneumonia

### server

提供rss 和 RestFul API

数据存储于 Mysql，服务使用 腾讯云 云函数 + 网关部署。

#### URL

*restful API*

```shell script
http://service-5lpzkj7x-1254035985.bj.apigw.tencentcs.com/release/api/ncov/news
```

*RSS Feed*

```shell script
http://service-5lpzkj7x-1254035985.bj.apigw.tencentcs.com/release/ncov/news.xml
```

#### TODO

[ ] 数据更新RSS
[ ] 数据更新API
[ ] 数据存储到飞书文档 需要ip 白名单
[ ] 数据存储到腾讯文档 需要申请 