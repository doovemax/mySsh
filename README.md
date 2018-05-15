# mySsh

基于 islenbo的 autossh，这个是源地址： https://github.com/islenbo/autossh.git
- 去除了软件编辑配置文件的功能，个人认为直接编辑配置文件更加方便和习惯
- 支持目录
- 支持Name选择方式登录
- 支持指定host方式登录
- 配置文件保存在$HOME/.mySsh目录中(第一次没有配置文件会自动创建)
- 支持输入密码不可见(需安装openssl开发库)
- 支持指定配置文件



## 安装 openssl 库:
### Mac:
```
git clone https://github.com/openssl/openssl.git && cd openssl
./config
make
cp -a include/* /usr/include/openssl/
```

### Ubuntu:
```
sudo apt install libssl-dev
```

### Centos && Fedora
```
sudo yum install openssl-devel
```
