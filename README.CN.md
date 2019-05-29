# go-blog

地址: [https://www.iphpt.com](https://www.iphpt.com)

> 自己花了一些时间单独重构的!

## 主要功能有:

1. 文章发布和修改
2. 文章回收站和撤回
3. 文章分类
4. 文章标签
5. 网站信息自由设置
6. 文章支持markdown
7. ★ 网站静态文件可自由配置`本地`或 `CDN`
8. ★ 可上传图片,可选择上传至`服务器`或 `七牛` 或 `两者同时`
9. 自由添加友链和管理友链顺序
10. 采用`github`的`issue`功能作为评论系统,界面优美且方便其他用户留言和通知
11. ★ 定时备份数据和静态资源
12. ★ 备份数据邮件发送至邮箱

## 技术点:

###  主要代码是 `golang`+`vue`+`HTML`+`CSS`+`MySQL`
    - 后端代码是基于`golang`的`gin`框架封装成的一个自用的包  [https://github.com/izghua/zgh](https://github.com/izghua/zgh)
    - 后台代码是基于`vue`的`iview`UI组件开发的后台, [https://github.com/izghua/go-blog-backend](https://github.com/izghua/go-blog-backend)
    - 前台是基于`HTML+CSS`,展示页面 [https://www.iphpt.com](https://www.iphpt.com)
    - 缓存用的`redis`
    - 数据库用的是 `MySQL`
    - 配置文件用的 `yaml`
    
### 安装方法

[go-blog安装教程](https://www.iphpt.com/detail/130)    
[关于go-blog的配置文件解析](https://www.iphpt.com/detail/131)
    
### 历史

本站起始于2015年,最开始由`zblog`开源框架搭建

在2016年,我利用`PHP`的`Laravel`框架 重构了前后端所有代码!
    - 界面: [http://php.iphpt.com](http://php.iphpt.com)
    - 代码: [https://github.com/xzghua/laravel-5-myblog](https://github.com/xzghua/laravel-5-myblog)

在2019年,我利用`Golang`再次重构了它
    - 界面: [https://www.iphpt.com](https://www.iphpt.com)
    - 代码: [https://github.com/izghua/go-blog](https://github.com/izghua/go-blog)

### 历史版本:
    Z-blog: [https://github.com/xzghua/iphpt](https://github.com/xzghua/iphpt)
    Laravel: [https://github.com/xzghua/laravel-5-myblog](https://github.com/xzghua/laravel-5-myblog)
    Beego(未上线): [https://github.com/xzghua/beego-postgres-blog](https://github.com/xzghua/beego-postgres-blog)
    Golang: [https://github.com/izghua/go-blog](https://github.com/izghua/go-blog)
    
### 展示图        
 
新:
   - 后台登录: ![后台登录](./static/uploads/images/bc-login.png)
   - 后台首页: ![后台首页](./static/uploads/images/bc-home.png)
   - 后台列表: ![后台列表](./static/uploads/images/bc-post.png)

旧:
   - 后台登录: ![后台登录](./static/uploads/images/old-bc-login.png)
   - 后台首页: ![后台首页](./static/uploads/images/old-bc-home.png)
   - 后台列表: ![后台列表](./static/uploads/images/old-bc-post.png)



 
