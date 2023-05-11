# pipe
pip enhance

## Usage:

```
$ ./pipe list
Website list:

https://pypi.douban.com/simple/
https://mirrors.aliyun.com/pypi/simple/
https://pypi.tuna.tsinghua.edu.cn/simple/
https://pypi.mirrors.ustc.edu.cn/simple/
http://mirrors.cloud.tencent.com/pypi/simple/
https://repo.huaweicloud.com/repository/pypi/simple/
https://mirrors.163.com/pypi/simple/

$ ./pipe probe
Ranking by response time:

1. Huawei - Response time: 137.045156ms
2. Aliyun - Response time: 229.805095ms
3. Douban - Response time: 310.949881ms
4. Tsinghua - Response time: 391.314535ms
5. Tencent - Response time: 456.283064ms
6. USTC - Response time: 808.627582ms
7. 163 - Response time: 844.355082ms

$ ./pipe config
Run command:

pip config set global.index-url https://repo.huaweicloud.com/repository/pypi/simple/
pip config set install.trusted-host repo.huaweicloud.com

$ ./pipe doc
Documentation:

https://pip.pypa.io/en/stable/topics/configuration/

$ ./pipe version
Pipe(pip enhance):

version 1.0.0

$ ./pipe demo
Run command:

pip install requests -i https://repo.huaweicloud.com/repository/pypi/simple/ --trusted-host repo.huaweicloud.com
```
