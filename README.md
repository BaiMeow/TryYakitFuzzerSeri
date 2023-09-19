# YakitFuzzerSeriPlayground

适用于体验 Yakit Fuzzer 序列的模拟 AWD 靶场

## 使用方法

`git clone` 后 `cd` 进来 `go run .` 即可

对靶机数量不满意可以改 `const TargetCount = 3` 这个常量

默认 flag 提交开在 8080 端口，靶机端口从 8081 开始按照顺序开

flag 提交

```curl
curl -X POST -d "flag={123123123}" http://localhost:8080/pushflag
```
