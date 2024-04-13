# Telegrammer 

The bot features a streamlined interface designed for ease of operation. When sending a request to the bot, users have the option to include a parameter called type, which specifies the desired response format as either `ini` or `json`. The default response format is set to `json`.

It is important to note that the uid parameter is mandatory for all bot methods, ensuring proper functionality. The bot is capable of processing both `GET` and `POST` requests.

In the event of an error, the bot will return `{"error":"error description"}`, along with the corresponding `HTTP status code`. For successful requests, the bot will return `{"key_name":"value for key"}` as part of the response.

### API → Send

`/send?uid=123&text=Hello&type=ini`  
The method sends text to the user.

`/send?uid=123&html=<b>Hello</b>&type=json`  
The method sends text to the user in HTML format. Only simple tags are supported (see Telegram description).

`/send?uid=123&photo=Base64&type=ini`  
The method parses Base64 and sends it as an image to the user.

`/send?uid=123&photo=./sample/filename.png&type=json`  
The method sends the file specified in the local path to the user.

`/send?uid=123&photo=http://myserver.com/filename.png&type=ini`  
The method sends the file specified in the URL to the user.

### API → Check

`/check?uid=123&chat=12345&type=json`  
The method checks if the user is in the specified chat. To perform the check, the bot must be in the specified chat.

### API → Data

`/set?uid=123&key=mykey&value=myvalue&type=json`  
The method saves the key and value in the storage, and this data will be located in the same file of the user created at the start.

`/get?uid=123&key=mykey&type=json`  
The method retrieves the saved value from the storage.

`/set?uid=123&keys=key1,key2&values=value1,value2&type=ini`  
The method saves keys and values in the storage, and this data will be located in the same file of the user created at the start.

`/get?uid=123&keys=key1,key2&type=ini`  
The method retrieves the saved values from the storage.
