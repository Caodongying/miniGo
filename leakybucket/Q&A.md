# v1.0
1. go文件的命名规则是什么？带下划线还是大小写？变量和函数的命名规则是什么
   大写开头的是暴露给外界的，相当于public。文件命名就是含有下划线。包名要和目录一样，全部小写字母。变量驼峰。bool类型有动词前缀。
   https://www.cnblogs.com/rickiyang/p/11074174.html
2. package main，容易混淆，深入了解一下
3. 学一下mermaid画图
4. struct的声明
5. 有几种int，怎么区别使用
6. 为什么else一定要紧跟在if的"}"后面
7. 小数和整数的直接运算，不可以
8. log管理方法，原生的log或者Uber开发的zap
9. package和module的区别？为什么leakybucket folder下面只有一个文件写了```package leakybucket```就没问题，但是两个文件用了这个package就报错