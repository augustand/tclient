# encoding=utf-8

import socket

import time

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM, 0)
s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
# 也可以使用s=socket.socket()来通过默认参数来生成TCP的流套接字
# host=''   host可以为空，表示bind()函数可以绑定在所有有效地地址上
host = "localhost"
port = 1235
s.bind((host, port))  # 注意，bind函数的参数只有一个，是（host,port）的元组
s.listen(3)

client, ipaddr = s.accept()
for i in range(10000):
    try:
        print i
        print ipaddr

    # data = client.recv(1024)

    # print "%d receive data:  %s" % (i, data)

    # client.send("echo:" + data + "\n")
        client.send(str(i)+'{"token":"2345"}\n')
        time.sleep(0.5)
    except:
        client, ipaddr = s.accept()

    # time.sleep(1)
#for i in range(10000):
 #   time.sleep(3)
  #  print s.accept()

time.sleep(10000)
client.close()
