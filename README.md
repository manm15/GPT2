# chatgpt
golang sdk Wrapper for ChatGPTAPIFree

repo:https://github.com/ayaka14732/ChatGPTAPIFree

To use ChatGPT API Free, simply send a POST request to the following endpoint:

https://chatgpt-api.shn.hk/v1/
For instance, to generate a response to the prompt "Hello, how are you?" using the gpt-3.5-turbo model, send the following curl command:
```bash
curl https://chatgpt-api.shn.hk/v1/ \
-H 'Content-Type: application/json' \
-d '{
"model": "gpt-3.5-turbo",
"messages": [{"role": "user", "content": "Hello, how are you?"}]
}'
```


