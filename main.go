package main

import (
    "flag"
    "fmt"
    "net/http"
    "net/url"
    "sort"
    "time"
)

type site struct {
    Name         string
    URL          string
    ResponseTime time.Duration
}

func main() {
    // 定义测试的网站列表
    sites := []site{
        {"Douban", "https://pypi.douban.com/simple/", 0},
        {"Aliyun", "https://mirrors.aliyun.com/pypi/simple/", 0},
        {"Tsinghua", "https://pypi.tuna.tsinghua.edu.cn/simple/", 0},
        {"USTC", "https://pypi.mirrors.ustc.edu.cn/simple/", 0},
        {"Tencent", "http://mirrors.cloud.tencent.com/pypi/simple/", 0},
        {"Huawei", "https://repo.huaweicloud.com/repository/pypi/simple/", 0},
        {"163", "https://mirrors.163.com/pypi/simple/", 0},
    }

    // 根据子命令执行不同的操作
    flag.Parse()
    switch flag.Arg(0) {
    case "probe":
        probe(sites)
    case "list":
        list(sites)
    case "config":
        config(sites)
    case "demo":
        demo(sites)
    case "doc":
        doc()
    case "version":
        version()
    default:
        fmt.Println("Usage: pipe [probe|list|config|demo]")
    }
}

// 探测网站访问速度并按照访问时间排名
func _probe(sites []site) {
    for i, s := range sites {
        start := time.Now()
        _, err := http.Get(s.URL)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            continue
        }
        elapsed := time.Since(start)
        sites[i].ResponseTime = elapsed
        // fmt.Printf("%v - Response time: %v\n", s.Name, elapsed)
    }

    sort.Slice(sites, func(i, j int) bool {
        return sites[i].ResponseTime < sites[j].ResponseTime
    })
}


// 探测网站访问速度并按照访问时间排名
func probe(sites []site) {
    _probe(sites)

    fmt.Println("Ranking by response time:\n")
    for i, s := range sites {
        fmt.Printf("%d. %s - Response time: %v\n", i+1, s.Name, s.ResponseTime)
    }
}

// 显示网站列表
func list(sites []site) {
    fmt.Println("Website list:\n")
    for _, s := range sites {
        fmt.Println(s.URL)
    }
}

// 访问最快的网站，并将结果输出到文件
func config(sites []site) {
    _probe(sites)

    // 访问时间最短的网站
    fastest := sites[0]

    u, err := url.Parse(fastest.URL)
    if err != nil {
        panic(err)
    }

    // 输出到文件
    output := fmt.Sprintf(`Run command:

pip config set global.index-url %s
pip config set install.trusted-host %s`, fastest.URL, u.Hostname())

    fmt.Println(output)
}

// 访问最快的网站，并将结果输出到文件
func demo(sites []site) {
    _probe(sites)

    // 访问时间最短的网站
    fastest := sites[0]

    u, err := url.Parse(fastest.URL)
    if err != nil {
        panic(err)
    }

    // 输出到文件
    output := fmt.Sprintf(`Run command:

pip install requests -i %s --trusted-host %s`, fastest.URL, u.Hostname())

    fmt.Println(output)
}

func doc() {
    fmt.Println("Documentation:\n")
    fmt.Println("https://pip.pypa.io/en/stable/topics/configuration/")
}

func version() {
    fmt.Println("Pipe(pip enhance):\n")
    fmt.Println("version 1.0.0")
}
