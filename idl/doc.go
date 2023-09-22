// idl 分布式对象接口定义, 即我们的 rpc 接口定义.
// 一般建议将生成的文件放到单独仓库进行管理, 从而避免直接依赖业务仓库产生循环引用.
// 为减少流程复杂性, 一般会引入 proto.ci, 当需要发包时手动触发流程, makefile
// 实现了一个简单的 demo.
// ps: 建议 protocl 大家使用一致的格式, 且区分 dev/master, dev 只发小版本.

package idl