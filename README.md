# cobra-gen-drone
通过cobra生成drone的模板


## 1、执行生成 .drone.yaml 文件

```bash
# 长指令
drone --droneName="base" --go_private="gitee.com" --service_name="baserpc.go" --service_type="rpc" --gitBranch="master" --registry="registry.cn-beijing.aliyuncs.com" --repo="registry.cn-beijing.aliyuncs.com/ctra_test/xxx-rpc" --tag="latest"

# 短指令
drone -d="base" -g="gitee.com" -s="baserpc.go" -x="rpc" -b="master" -r="registry.cn-beijing.aliyuncs.com" -o="registry.cn-beijing.aliyuncs.com/ctra_test/xxx-rpc" -t="latest"
```
