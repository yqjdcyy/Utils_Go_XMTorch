
# 公寓申请时长统计
## 目的
- 确认申请时长，进行时间估算，方便做选择

## 框架
- 配置
- 日志服务
- 参数化服务调用
- 工具
    - 文件
    - 命令行


## 数据下载
- 网站请求
    - 获取 cookie
- 服务请求
    - 请求链接拼接
    - 延时请求
    - 数据处理
        - 回包解析
        - 数据封装
        - 存储
            - 3117 records

## 统计报表
- 数据读取、解析
- 业务过滤、处理
- 报表模板
- 报表数据
    - 时长
        - [申请耗时分布图.html](http://doc.yqjdcyy.com/a04f2148-9a27-4994-9591-81193a910831.html)
        - ![申请耗时分布图.png](http://doc.yqjdcyy.com/6b2aefdf-00f1-4657-83b3-460b04ea4709.png)
        - ![申请耗时分布数据.png](http://otzm88f21.bkt.clouddn.com/7bc17b28-d53f-466c-84a2-45c2dd1a1363.png)
    - 企业
        - 补充过滤条件
            - 通过3人以上
            - 总申请人数5人以上
        - [企业通过与不通过.html](http://doc.yqjdcyy.com/bba2244f-fada-4a0d-ab3f-0ab8bf8a4a6c.html)
        - ![企业通过与不通过.png](http://doc.yqjdcyy.com/48e0b22a-0bdb-4d6c-858c-05fa6cfbbc72.png)
        - ![企业角度通过情况.png](http://doc.yqjdcyy.com/b1b97b6a-1ad2-4c0a-965d-5ef85e0d51cf.png)
        - ![企业通过与不通过数据.png](http://doc.yqjdcyy.com/f2412135-9ca5-4abd-967e-ccfa7d6666a5.png)
- 自动打开网页
