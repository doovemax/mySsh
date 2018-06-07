# mySsh

基于 islenbo的 autossh，这个是源地址： https://github.com/islenbo/autossh.git
- 去除了软件编辑配置文件的功能，个人认为直接编辑配置文件更加方便和习惯
- 配置文件支持目录
- 支持Name选择方式登录
- 支持指定host方式登录
- 配置文件保存在$HOME/.mySsh目录中(第一次没有配置文件会自动创建)以.json结尾的文件
- 支持输入密码不可见(需安装openssl开发库)
- 支持指定配置文件
- 支持将 pem 文件放在配置文件目录下,配置文件写 pem 名字即可



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

## Example:

```
Usage of ./mySsh:
  -f string
    	specify config file
  -h	Usage
  -help
    	Usage
  -host string
    	specity remote host
  -list
    	list remote hosts
  -port int
    	specity remote port (default 22)
  -version
    	Print version

./mySsh -h |-help
./mySsh  -f server.json
./mySsh  -host IP [-port 22](Default)
./mySsh  -version
./mySsh [-port 2222] user@host #port 参数必须在前面
```