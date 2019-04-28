

### 使用方法

* 从远程仓库搜索镜像

        docker search 镜像名称
        
* 从远程下载镜像       
        
        docker pull 镜像名称 

* 查看镜像
 
        docker images  或者  docker images | grep '镜像名称'
      

* 构建镜像，名称 docker_go_demo . 是Dockerfile所在的目录
    
        docker build -t docker_go_demo .
        
* 删除镜像

        docker rmi 镜像ID|镜像名称     
        
* 删除所有镜像

        docker rmi $(docker images -q)        
 
    
* 运行这个这个镜像
 
        docker run docker_go_demo
        
* 运行镜像暴露80端口，第一个80为host（物理机）端口，第二个80为contain(容器)端口
    
        docker run -p 8080:8080 docker_go_demo
        
* 为一个镜像打标签
        
        docker tag docker_go_demo sunpengwei/docker_go_demo:1.0   此时再运行docker images 就看会看到 sunpengwei/docker_go_demo:1.0 这个镜像了
        
* 查看有那些容器
          
        docker ps
        参数 -a  显示所有的容器  docker ps -a
        参数 -f 根据条件过滤显示的内容,-f后面跟的都是键值对 docker ps -f name=docker_go_demo  ； docker ps -f id=a1b2c3 -f name=bingohuang
        参数 -q 只显示容器的ID docker ps -q
        
* 如何启动，重启，停止容器

        docker start|restart|stop 容器ID|容器名称

* 停用全部运行中的容器

        docker stop $(docker ps -q)
        
* 删除全部容器

        docker rm $(docker ps -aq) 

* 既停用又删除容器

        docker stop $(docker ps -q) & docker rm $(docker ps -aq)                    
        

* 如何删除容器
    
        docker rm   容器ID|容器名称
        
* 进入容器的终端

        docker exec -it 容器ID /bin/bash    
        
        
* docker run 常用选项说明
 
        -d, --detach=false， 指定容器运行于前台还是后台，默认为false 
        -i, --interactive=false， 打开STDIN，用于控制台交互
        -t, --tty=false， 分配tty设备，该可以支持终端登录，默认为false
        -u, --user=""， 指定容器的用户
        -a, --attach=[]， 登录容器（必须是以docker run -d启动的容器）
        -w, --workdir=""， 指定容器的工作目录
        -c, --cpu-shares=0， 设置容器CPU权重，在CPU共享场景使用
        -e, --env=[]， 指定环境变量，容器中可以使用该环境变量
        -m, --memory=""， 指定容器的内存上限
        -P, --publish-all=false， 指定容器暴露的端口
        -p, --publish=[]， 指定容器暴露的端口
        -h, --hostname=""， 指定容器的主机名
        -v, --volume=[]， 给容器挂载存储卷，挂载到容器的某个目录
        --volumes-from=[]， 给容器挂载其他容器上的卷，挂载到容器的某个目录
        --cap-add=[]， 添加权限，权限清单详见：http://linux.die.net/man/7/capabilities
        --cap-drop=[]， 删除权限，权限清单详见：http://linux.die.net/man/7/capabilities
        --cidfile=""， 运行容器后，在指定文件中写入容器PID值，一种典型的监控系统用法
        --cpuset=""， 设置容器可以使用哪些CPU，此参数可以用来容器独占CPU
        --device=[]， 添加主机设备给容器，相当于设备直通
        --dns=[]， 指定容器的dns服务器
        --dns-search=[]， 指定容器的dns搜索域名，写入到容器的/etc/resolv.conf文件
        --entrypoint=""， 覆盖image的入口点
        --env-file=[]， 指定环境变量文件，文件格式为每行一个环境变量
        --expose=[]， 指定容器暴露的端口，即修改镜像的暴露端口
        --link=[]， 指定容器间的关联，使用其他容器的IP、env等信息
        --lxc-conf=[]， 指定容器的配置文件，只有在指定--exec-driver=lxc时使用
        --name=""， 指定容器名字，后续可以通过名字进行容器管理，links特性需要使用名字
        --net="bridge"， 容器网络设置:
        bridge 使用docker daemon指定的网桥
        host //容器使用主机的网络
        container:NAME_or_ID >//使用其他容器的网路，共享IP和PORT等网络资源
        none 容器使用自己的网络（类似--net=bridge），但是不进行配置
        --privileged=false， 指定容器是否为特权容器，特权容器拥有所有的capabilities
        --restart="no"， 指定容器停止后的重启策略:
        no：容器退出时不重启
        on-failure：容器故障退出（返回值非零）时重启
        always：容器退出时总是重启
        --rm=false， 指定容器停止后自动删除容器(不支持以docker run -d启动的容器)
        --sig-proxy=true， 设置由代理接受并处理信号，但是SIGCHLD、SIGSTOP和SIGKILL不能被代理
        
* docker run 示例

        运行一个在后台执行的容器，同时，还能用控制台管理：docker run -i -t -d ubuntu:latest
        运行一个带命令在后台不断执行的容器，不直接展示容器内部信息：docker run -d ubuntu:latest ping www.docker.com
        运行一个在后台不断执行的容器，同时带有命令，程序被终止后还能重启继续跑，还能用控制台管理，docker run -d --restart=always ubuntu:latest ping www.docker.com
        为容器指定一个名字，docker run -d --name=ubuntu_server ubuntu:latest
        容器暴露80端口，并指定宿主机80端口与其通信(: 之前是宿主机端口，之后是容器需暴露的端口)，docker run -d --name=ubuntu_server -p 80:80 ubuntu:latest
        指定容器内目录与宿主机目录共享(: 之前是宿主机文件夹，之后是容器需共享的文件夹)，docker run -d --name=ubuntu_server -v /etc/www:/var/www ubuntu:latest
                    

* 查看容器运行过程中CPU，内存使用情况

        docker stats 容器ID        
        
        
              
        
        
        