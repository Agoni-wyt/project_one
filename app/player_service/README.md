1. 用户表和背包表
2. 用户注册-->登录-->读取用户信息
3. 支持功能
   1. 注册                       http
   2. 登录                       http
   3. 查看个人信息                http/rpc      自己查看/交易时显示
   4. 添加物品道具                http/rpc      自己获取/其它服务发放
   5. ~~删除物品道具                http/rpc      自己使用/其它服务使用~~物品仅可转移
   6. 查看个人背包                http/rpc      自己查看/交易时显示
   7. 修改背包物品                http/rpc      交易

如果需要HTTP和rpc同时存在，应该通过http服务调用rpc，中间转变为pb类型，并将rpc响应转变为http响应

微服务应该各个服务之间区分开，pb文件需要位于公共目录，以便不同客户端的获取

