wt -w 0 nt -p "Command Prompt" cmd /k "cd C:\Users\abahbob\source\eft-private-server\client-server && make && make run && pause"
wt -w 0 nt -p "Command Prompt" cmd /k "cd C:\Users\abahbob\source\eft-private-server\notifier-server && make && make run && pause"
wt -w 0 nt -p "Command Prompt" cmd /k "cd C:\Users\abahbob\source\eft-private-server\game-server && make && make run && pause"
wt -w 0 nt -p "Command Prompt" cmd /k "npx local-ssl-proxy --source 8081 --target 8080 && pause"