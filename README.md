# 业务-技术-逻辑
本模块基本逻辑：  
1.用户发送消息->微信转发到系统->系统转发到上线客服->客户在48小时内回复->系统转发到微信->微信转发到用户。  
简化模式：  
用户->微信->xml->系统->JSON->客服（xml格式由微信定义）  
客服->JSON->系统->xml->微信->用户  
本模块，现在只实现：XML -> 系统 -> JSON，并对微信定义的XML格式进行简化。  
2.客服和系统通过socket通信。  
3.对于用户发送的多媒体文件，系统先进行转存转格式，再把url给客服。  
4.在系统中用一个map维护 接收者-消息 关系，形成 接收者形成联系人列表，接收者-消息 是一对多。  
5.若需要持久化消息，可以使用一个 消息队列 以接收者id为主键，异步保存到mysql。

参考文档：  
【1】https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html  
【2】https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_standard_messages.html  
【3】https://github.com/gorilla/websocket/tree/master/examples

【注意】这是只是一个微信公众号系统对接模块，展示一种技术对接方式，不涉及具体业务实现。

# 关于websocket连接的管理
每一个连接请求通过HandleFunc注册的Handler处理，http包会开启一个协程以保证并发需要，然后建立连接后，该Handler不结束，一直监听该连接，直到连接close，该Handler函数结束执行。
【注意】不要尝试着把连接保存起来，然后另开一个函数轮询监听该连接，这样是行不通的。  
<br />