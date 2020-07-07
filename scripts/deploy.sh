echo "Building discord client..."
go build ./src/peach_discord_client
echo "."
echo "."
echo "Restarting service..."
systemctl --user restart peach
echo "."
echo "."
echo "Done :)"