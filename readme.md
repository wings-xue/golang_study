通过学习本教程，提高工程能力
1. 如何从成熟项目学习工程能力

day 1
1. http标准库基础是构造一个地址和处理函数的词典

day 2 
1. 请求和路由分离

day 3
1. 请求支持两种匹配
   1. :name
   2. *fildpath

错误
没有设计出来应该有的数据结构
设计混乱。实现了树但是不知道如何关联在一起


文章
对于路由来说，最重要的当然是注册与匹配了。开发服务时，注册路由规则，映射handler；访问时，匹配路由规则，查找到对应的handler。因此，Trie 树需要支持节点的插入与查询。插入功能很简单，递归查找每一层的节点，如果没有匹配到当前part的节点，则新建一个，有一点需要注意，/p/:lang/doc只有在第三层节点，即doc节点，pattern才会设置为/p/:lang/doc。p和:lang节点的pattern属性皆为空。因此，当匹配结束时，我们可以使用n.pattern == ""来判断路由规则是否匹配成功。例如，/p/python虽能成功匹配到:lang，但:lang的pattern值为空，因此匹配失败。查询功能，同样也是递归查询每一层的节点，退出规则是，匹配到了*，匹配失败，或者匹配到了第len(parts)层节点。

1. 分析功能
路由注册与匹配
注册
映射handle，path和handle关联
树需要有插入功能
插入功能很简单，递归查找每一层的节点，如果没有匹配到当前part的节点，则新建一个

节点调用search函数，如果当前节点part是传入的part则用此节点递归调用下一层节点，否则新建节点，用新建节点调用下一层节点
基于这句话设计api
pattern: 待匹配路由
path: 待匹配路由组成的词典
height: 插入路由的路径位置
insert(pattern, paths []string, height int)

错误
search 基于树递归应该先考虑逻辑设计api而不是数据结构，但是利用数据结构完成api的实现



匹配
insert