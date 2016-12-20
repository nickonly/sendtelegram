# Send Telegram messages from command-line
**telegramsend** is a simple crossplatform command-line tool to send messages over Telegram.

### 1. Build
```bash
git clone https://github.com/shpnick/sendtelegram
cd sendtelegram
go build
```

### 2. Create bot with the BotFather
On Telegram, open the @BotFather profile, and start to talk with “him”. You can open the conversation by clicking here. If the screen remains empty, update your Telegram client, or write “/start”
Type the command “/newbot”.
BotFather will ask you for a bot name, which is a free text, and a bot username: it always end by “bot”. If your wanted username is not available, try again with “yourname_bot”, “yournamerobot”, or “yourname_robot” for example.

### 3. Give your id
Open the conversation with your bot and send any message. Then use:
```bash
curl -X POST 'https://api.telegram.org/bot<token>/getUpdates'
```
<token> is we got in the previous step.
Find "from":{"id":<user_id>, ...

### 4. Usage
```bash
./sendtelegram -t <token> -m <message> -i <user_id>
```
