封包的云包模块，提供包的上传、下载和删除功能，目前通过文件方式读写。

内置搭配了cronjob使用，周四自动清空云包，效果等同于：

```
5 6 * * 4 echo [] > /home/USER/packet_cloud/packets
```