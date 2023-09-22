// bootstrap 用于管理项目所有的依赖, 并通过依赖注入的方式保证全局只初始化一次依赖实例.
// 在主进程结束前调用 bootstrap.Teardown 以合理的关闭资源.
package bootstrap
