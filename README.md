# resource_consume
用于消耗指定的cpu使用量和内存使用量

### 编译

执行 如下命令完成编译：

```shell
go build -o consumer github.com/Chuang1993/resource_consume/consumer
go build -o gdcpu github.com/Chuang1993/resource_consume/cpu
```

### 使用

将二进制文件consumer和gdcpu赋予运行权限即可执行。其中：

* consumer为整体启动命令，可以消耗CPU和内存
* gdcpu仅可以消耗cpu

consumer命令行参数：
-cpu-gradient	字符串	cpu使用梯度（单位为m,1c=1000m），循环执行，70m以下不敏感
-cpu-interval	整数类型	cpu梯度中每个梯度执行时长（单位为s）默认为60
-cpuexec	字符串	指定消耗cpu的二进制文件绝对路径，默认为gdcpu，默认情况下，需要将gdcpu放置PATH环境变量下目录才能使用
-memory	整数类型	内存消耗量（单位为M）

cs_cpu命令行参数：
-cpu-gradient	字符串	cpu使用梯度（单位为m）
-cpu-interval	整数类型	cpu梯度中每个梯度执行时长（单位为s）默认为60

### 示例
```
./consumer -cpu-gradient 50,100,30,500,10 -cpu-interval 10 -memory 100 -cpuexec gdcpu
```


cpu使用量梯度为50m,100m,30m,500m,10m，每个梯度cpu执行时间为10s，内存消耗100M

### 说明

指定cpu消耗量部分，部分代码参考了k8s内的实现：

https://github.com/kubernetes/kubernetes/blob/master/test/images/resource-consumer/consume-cpu/consume_cpu.go

在这个基础上进行了一次升级改造，可实现梯度指定，尽可能模拟容器内服务周期性的cpu消耗