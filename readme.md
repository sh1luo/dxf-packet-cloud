封包的云包模块，提供包的上传下载功能，目前通过文件方式读写。

搭配cronjob使用，周四自动清空云包举例：

```
5 6 * * 4 echo [] > /home/USER/packet_cloud/packets
```