start:
        test -z "`pidof bingwall`"
        nohup ./bingwall -

stop:
        killall bingwall

restart: stop start