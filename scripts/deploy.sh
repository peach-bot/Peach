echo "Building discord client..."
go build ./src/peach_discord_client
echo ".\n.\nRestarting service..."
systemctl --user restart peach
echo ".\n.\nDone :)"