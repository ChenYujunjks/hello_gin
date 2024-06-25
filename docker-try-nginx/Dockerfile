# 使用官方Nginx镜像作为基础镜像
FROM nginx:latest

# 将本地的配置文件复制到容器中的Nginx配置目录
COPY nginx.conf /etc/nginx/nginx.conf

# 暴露容器中的80端口
EXPOSE 80

# 运行Nginx
CMD ["nginx", "-g", "daemon off;"]
