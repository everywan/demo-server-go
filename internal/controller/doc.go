// controller 对外交互层, 经典的 controller-service-dao 三层结构
// controller 负责输入数据的转换, 输出数据的处理, 为核心业务逻辑 service 层服务.
// 一般而言, 所有的错误日志都放到出口(controller)打印, 内部不打印日志, 从而避免重复打印.

package controller
