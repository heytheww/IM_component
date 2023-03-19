# 业务-技术-逻辑
本模块基本逻辑：
1.用户发送消息->微信转发到系统->系统转发到上线客服->客户在48小时内回复->系统转发到微信->微信转发到用户。
简化模式：
用户->微信->xml->系统->JSON->客服（xml格式由微信定义）
客服->JSON->系统->xml->微信->用户
本模块，现在只实现：JSON/XML -> 系统 -> JSON，并对微信定义的XML格式进行简化。

2.客服和系统通过socket通信。
3.对于用户发送的多媒体文件，系统先进行转存转格式，再把url给客服。
4.在系统中用一个map维护 用户-客服 关系，形成 联系人列表，用户-客服 是多对一。

参考文档：
【1】https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html
【2】https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_standard_messages.html
【3】https://github.com/nhooyr/websocket/blob/master/examples/chat/main.go

【注意】这是只是一个微信公众号系统对接模块，展示一种技术对接方式，不涉及具体业务实现。

# 接口列表
## 1. 连接socket & 从系统接收消息
wss://abc.com/im_component/wx_official_account/msg?userid=xxx

【收到消息格式】
1.1 系统消息
{
    "data":{
        "id": "", // 消息id
        "name": "系统消息",
        "avatar":"", // 头像url
        "msg_type": "system",
        "contact_id": "system01", // 对方id，统一为system01
        "company_id": "", // 业务id，客户所在公司id
        "send_time": 1678688556, // 消息产生的时间戳
        "msg":{}
    },
    "result":{
        "code":"102", // 系统消息102
        "message":"common"
    }
}


1.2 普通消息
1.2.1 text消息
{
    "data":{
        "id": "", // 消息id
        "name": "", // 用户名
        "avatar":"", // 头像url
         "msg_type": "text", 
        "contact_id": "", 
        "company_id": "", // 业务id，客户所在公司id
        "send_time": 1678688556, // 消息产生的时间戳
        "msg":{
            "content":"",// 文本内容
        }
    },
    "result":{
        "code":1001,  
        "message":"common"
    }
}

1.2.2 image消息
{
    "data":{
        "id": "", // 消息id
        "name": "", // 用户名
        "avatar":"", // 头像url
        "msg_type": "image", 
        "contact_id": "", 
        "company_id": "", // 业务id，客户所在公司id
        "send_time": 1678688556, // 消息产生的时间戳
        "msg":{
            "media_path":""
        }
    },
    "result":{
        "code":1002,    
        "message":"common"
    }
}

1.2.2 voice消息
{
    "data":{
        "id": "", // 消息id
        "name": "普通消息",
        "avatar":"", // 头像url
        "msg_type": "voice", 
        "contact_id": "", 
        "company_id": "", // 业务id，客户所在公司id
        "send_time": 1678688556, // 消息产生的时间戳
        "msg":{
            "media_path":""
        }
    },
    "result":{
        "code":1003,    
        "message":"common"
    }
}

1.2.2 video消息
{
    "data":{
        "id": "", // 消息id
        "name": "普通消息",
        "avatar":"", // 头像url
        "msg_type": "video", 
        "contact_id": "", 
        "company_id": "", // 业务id，客户所在公司id
        "send_time": 1678688556, // 消息产生的时间戳
        "msg":{
            "media_path":""
        }
    },
    "result":{
        "code":1004,    
        "message":"common"
    }
}

1.3 心跳消息
{
    "data":{},
    "result":{
        "code":101, // 心跳101
        "message":"heartbeat"
    }
}

1.4 错误消息（预留）
{
    "data":{},
    "result":{
        "code":-1, 
        "message":"error"
    }
}

【发送消息格式】


## 2. 向系统推送消息（JSON）
2.1 JSON
http://abc.com/im_component/wx_official_account/send
[POST]
{
    "data":{
        "msg_type": "text", 
        "contact_id": "", 
        "msg":{
            "content":"",// 文本内容
        }
    },
    "result":{
        "code":1001,  
        "message":"common"
    }
}

2.2 XML
http://abc.com/im_component/wx_official_account/send_xml
[POST]
text消息
```
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <CreateTime>1348831860</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>
  <Content><![CDATA[this is a test]]></Content>
  <MsgId>1234567890123456</MsgId>
</xml>
```
其他类型普通消息，仅msg_type和code不同，不赘述。

## 3. 关于websocket连接的管理
每一个连接请求通过HandleFunc注册的Handler处理，http包会开启一个协程以保证并发需要，然后建立连接后，该Handler不结束，一直监听该连接，直到连接close，该Handler函数结束执行。
【注意】不要尝试着把连接保存起来，然后另开一个函数轮询监听该连接，这样是行不通的。

## 4. 测试方式
向 localhost:1234/send 发送 post
body为：
```
<xml>
  <ToUserName>1</ToUserName>
  <FromUserName>魏秀兰</FromUserName>
  <CreateTime>1976-04-10 12:21:47</CreateTime>
  <MsgType>enim</MsgType>
  <Content>occaecat ut in fugiat dolor</Content>
  <MsgId>19</MsgId>
  <MsgDataId>27</MsgDataId>
  <Idx>74</Idx>
</xml>
```

chrome使用如下代码连接websocket
```
<html>
<head></head>
<body>
	<script type="text/javascript">
		var sock = null;
		var wsuri = "ws://localhost:1234?userId=1";

		window.onload = function() {

			console.log("onload");

			sock = new WebSocket(wsuri,"echo");

			sock.onopen = function() {
				console.log("connected to " + wsuri);
			}

			sock.onclose = function(e) {
				console.log("connection closed (" + e.code + ")");
			}

			sock.onmessage = function(e) {
				console.log("message received: " + e.data);
			}
		};

		function send() {
			var msg = document.getElementById('message').value;
			sock.send(msg);
		};
	</script>
	<h1>WebSocket Echo Test</h1>
	<form>
		<p>
			Message: <input id="message" type="text" value="pong">
		</p>
	</form>
	<button onclick="send();">Send Message</button>
</body>
</html>

```