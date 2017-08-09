# 根据git提交导出仓库中变化过的文件

![](http://i.imgur.com/sNqT9kd.gif)



gitexport 将 `commits` 之间发生过变化或者新增的文件导出到指定的目录中（默认为系统的临时目录），导出后的文件会保持在原先仓库中的目录结构。  

如果涉及到被删除的文件，会将以删除的文件路径输出到终端。   

导出后会用系统文件浏览器打开导出后的目录。



## 使用方法

* `gitexport -r <commit>` 
* `gitexport -r <commit>..<commit>` 

`-r` 参数后面的内容实际上会传给  [`git diff`](https://git-scm.com/docs/git-diff)   命令，所以参数的使用可以参考  [`git diff`](https://git-scm.com/docs/git-diff) 和 [gitrevisions](https://git-scm.com/docs/gitrevisions) 。



默认情况下，导出的文件会放到系统的临时目录，如果要指定导出的路径可以加上  `-o` 参数，指定导出路径。

## 配合SourceTree使用

在SourceTree中添加一个自定义操作，如下图:

![](http://i.imgur.com/IMQdjxd.png)



然后就可以在查看提交记录时导出了



![](http://i.imgur.com/WTT03kK.gif)