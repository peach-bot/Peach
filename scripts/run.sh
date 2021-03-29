./scripts/build.sh

function on_exit()
{
    echo "Exiting..."
}

trap "on_exit" 2
echo "Running launcher"

cd build
./launcher