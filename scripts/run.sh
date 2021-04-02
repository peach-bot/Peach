./scripts/build.sh $1

function on_exit()
{
    echo "Exiting..."
}

trap "on_exit" 2
echo "Running launcher"

cd build
./launcher.exe