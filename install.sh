

echo "安装Ngrokd服务"
wget https://github.com/wuweiit/ngrok/archive/1.tar.gz

echo "解压文件包..."
tar -xvf 1.tar.gz


echo "centos7 安装依赖组件..."
yum -y install zlib-devel openssl-devel perl hg cpio expat-devel gettext-devel curl \
curl-devel perl-ExtUtils-MakeMaker hg wget gcc gcc-c++ git

echo "下载gosdk ..."
wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz
#vim /etc/profile
#echo "添加以下内容"
#export PATH=$PATH:/usr/local/go/bin
#source /etc/profile

echo "检测是否安装成功go"
go version

GOOS=linux GOARCH=amd64 make release-server