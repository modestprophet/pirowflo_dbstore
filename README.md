# pirowflo_mqtt
A simple MQTT subscriber service that processes [pirowflo_mqtt](https://github.com/modestprophet/pirowflo_mqtt) events and stores them to DB.  This projet is based on RabbitMQ and Postgres.  It should work with other MQ server types but will require customization to work with other database types.  

#### Installation
Install golang
```bash
# Apt (Ubuntu/Debian)
sudo apt install golang
```

Install pirowflo_dbstore with the user that will own the systemctl service
```bash
go install github.com/modestprophet/pirowflo_dbstore@latest
```


#### Config
pirowflow_dbstore expects a config (`.pirowfloconfig.json`) file in the user's home directory.  You can manually create the file using [.pirowfloconfig.sample.json](.pirowfloconfig.sample.json) as a template
```bash
vim ~/.pirowfloconfig.json
```

Alternatively, running pirowflo_dbstore with no config file present will automatically create a config template in the user's home directory
```bash
pirowflo_dbstore
```

```bash
user@computer:~/$ ./pirowflo_dbstore                    
Failed to read config: failed to create default config: config not found. please edit the newly created default config created at /home/user/.pirowfloconfig.json
```



#### Service setup
Create a .service file
```bash
sudo vim /etc/systemd/system/pirowflo-dbstore.service
```

Modify the `User` parameter to whatever user will own the service
```bash
[Unit]
Description=PiRowFlo Database Store Service
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=1
User=user
ExecStart=/bin/sh -c 'exec $(go env GOPATH)/bin/pirowflo_dbstore start'

[Install]
WantedBy=multi-user.target
```

Start the service
```bash
sudo systemctl enable pirowflo-dbstore 
sudo systemctl start pirowflo-dbstore
```




