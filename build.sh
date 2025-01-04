# 容器名
name=`cat version | awk '{print $1}'`
# 容器标签
tag=`cat version | awk '{print $2}'`


# 构建Docker镜像
docker build -t $name:$tag .

# 停止该镜像正在运行的Docker容器
lines=`docker ps | grep $name`
if [ -n "$lines" ]; then
  echo "存在正在运行的$name容器, 正在使其停止运行..."
  container_id=$(echo $lines | awk '{print $1}')
  docker stop $container_id
  echo "$name容器, 已停止运行"
fi

# 删除该镜像的Docker容器
lines=`docker ps -a | grep $name`
if [ -n "$lines" ]; then
    echo "存在$name容器, 对其进行删除..."
    container_id=$(echo $lines | awk '{print $1}')
    docker rm $container_id
    echo "$name容器, 已被删除"
fi

# 启动容器
docker run --rm --name $name -p 9002:8081 $name:$tag