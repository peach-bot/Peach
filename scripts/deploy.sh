echo "Building discord client..."
go build ./src/peach_discord_client
echo "."
echo "."
echo "Restarting service..."
systemctl --user restart peach
echo "."
echo "."
echo "Done :)"
systemctl --user status peach
journalctl --user -u peach --follow --since=now