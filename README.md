# urlcheck



### 一个简单的golang版本的URL检测工具



### 使用方式

1. 编译
2. 编写一个包含需要检测的url文件，以换行区分
  > 以 # 开头的内容表示注释
  > 空行将被忽略
  > 异常的行将被提示
3. ./urlcheck FILENAME [all|fail]
  > FILENAME url文件路径
  > all:输出所有url的检测结果
  > fail: 只输入检测异常的url结果
