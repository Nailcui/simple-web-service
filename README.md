此服务用来练习k8s，提供如下几个接口

> /hostname

```shell script
$ curl 10.106.207.157/hostname
{"hostname":"sws-deployment-6788869d58-8mk9j"}

$ curl 10.106.207.157/hostname
{"hostname":"sws-deployment-6788869d58-25nzg"}
```


```shell script
kubectl apply -f sws.yml
kubectl apply -f sws-service.yml
```
