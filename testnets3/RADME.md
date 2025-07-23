Erigon Service

sudo systemctl stop erigon.service

sudo systemctl daemon-reload

sudo systemctl start erigon.service

sudo journalctl -u erigon.service -f
